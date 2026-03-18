package scanner

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/zakirkun/ice-tea/internal/analyzer/llm"
	"github.com/zakirkun/ice-tea/internal/analyzer/pattern"
	"github.com/zakirkun/ice-tea/internal/analyzer/taint"
	"github.com/zakirkun/ice-tea/internal/config"
	"github.com/zakirkun/ice-tea/internal/finding"
	"github.com/zakirkun/ice-tea/internal/parser"
	"github.com/zakirkun/ice-tea/internal/reporter"
	"github.com/zakirkun/ice-tea/internal/skill"
)

// Engine is the main scan orchestrator
type Engine struct {
	config    *config.Config
	logger    *zap.SugaredLogger
	walker    *Walker
	parsers   *parser.Registry
	skills    *skill.Loader
	ruleIndex   *skill.RuleIndex
	reporters   *reporter.Registry
	llmProvider llm.Provider
}

// NewEngine creates a new scanning engine
func NewEngine(cfg *config.Config, logger *zap.SugaredLogger) *Engine {
	return &Engine{
		config:    cfg,
		logger:    logger,
		parsers:   parser.NewRegistry(),
		reporters: reporter.NewRegistry(),
	}
}

// Init initializes the engine components (loaders, registries)
func (e *Engine) Init() error {
	// 1. Initialize file walker
	e.walker = NewWalker(WalkerConfig{
		ExcludeDirs:       e.config.Exclude.Dirs,
		ExcludeFiles:      e.config.Exclude.Files,
		ExcludeExtensions: e.config.Exclude.Extensions,
		Languages:         e.config.Scan.Languages,
	}, e.logger)

	// 2. Load skills
	// In a real app, this would also load from external dirs if configured
	e.skills = skill.NewLoader([]string{e.config.Skills.Dir}, e.logger)
	loadedSkills, err := e.skills.LoadAll()
	if err != nil {
		e.logger.Warnw("Failed to load skills", "error", err)
		// We can continue even if some skills failed to load
	}
	e.ruleIndex = skill.BuildIndex(loadedSkills)
	e.logger.Infow("Skills loaded and indexed", "total_rules", len(e.ruleIndex.All))

	return nil
}

// RegisterParser adds a parser to the engine
func (e *Engine) RegisterParser(p parser.Parser) {
	e.parsers.Register(p)
}

// SetLLMProvider sets the LLM processing provider
func (e *Engine) SetLLMProvider(p llm.Provider) {
	e.llmProvider = p
}

// RegisterReporter adds a reporter to the engine
func (e *Engine) RegisterReporter(r reporter.Reporter) {
	e.reporters.Register(r)
}

// Run executes the full scan process
func (e *Engine) Run(ctx context.Context, target string) ([]*finding.Finding, error) {
	// 1. Discover files
	files, err := e.walker.Walk(target)
	if err != nil {
		return nil, fmt.Errorf("failed to discover files: %w", err)
	}

	if len(files) == 0 {
		e.logger.Info("No files to scan")
		return nil, nil
	}
	e.logger.Infow("Starting scan", "files", len(files), "concurrency", e.config.Scan.Concurrency)

	// 2. Setup concurrency channels
	fileChan := make(chan DiscoveredFile, len(files))
	resultChan := make(chan []*finding.Finding, len(files))

	for _, f := range files {
		fileChan <- f
	}
	close(fileChan)

	// 3. Start worker pool
	var wg sync.WaitGroup
	concurrency := e.config.Scan.Concurrency
	if concurrency > len(files) {
		concurrency = len(files)
	}

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go e.worker(ctx, &wg, fileChan, resultChan)
	}

	// Close result channel when all workers finish
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// 4. Aggregate findings
	aggregator := finding.NewAggregator(e.config.Scan.Severity)
	var processed int

	for workerFindings := range resultChan {
		processed++
		for _, f := range workerFindings {
			// Ensure file path is relative to target for reporting
			if f.File != "" && target != "." {
				// We keep the original path but can relative it if needed
			}
			aggregator.Add(f)
		}
	}

	e.logger.Debugw("Scan complete", "processed", processed)

	// 5. Generate Report
	findings := aggregator.Results()
	if err := e.generateReport(ctx, findings); err != nil {
		e.logger.Errorw("Failed to generate report", "error", err)
		return findings, err // return findings even if report fails
	}

	return findings, nil
}

// worker processes files from the queue
func (e *Engine) worker(ctx context.Context, wg *sync.WaitGroup, files <-chan DiscoveredFile, results chan<- []*finding.Finding) {
	defer wg.Done()

	// Initialize thread-local matchers
	patternMatcher := pattern.NewMatcher(e.ruleIndex.All, e.logger)

	for f := range files {
		select {
		case <-ctx.Done():
			return // Context cancelled
		default:
		}

		start := time.Now()
		
		// 1. Parse File
		src, err := os.ReadFile(f.Path)
		if err != nil {
			e.logger.Debugw("Failed to read file", "path", f.Path, "error", err)
			continue
		}

		lang := parser.Language(f.Language)
		parseResult, err := e.parsers.Parse(f.Path, src, lang)
		if err != nil {
			e.logger.Debugw("Failed to parse file", "path", f.Path, "lang", lang, "error", err)
			continue // Skip unsupported languages for now
		}

		var fileFindings []*finding.Finding

		// 2. Engine 1: Pattern Matching
		patternFindings := patternMatcher.Analyze(parseResult)
		fileFindings = append(fileFindings, patternFindings...)

		// 3. Engine 2: Taint Analysis
		taintTracker := taint.NewTracker(e.logger)
		taintFindings := taintTracker.Analyze(parseResult)
		fileFindings = append(fileFindings, taintFindings...)

		// 4. Engine 3: LLM Reasoning
		if e.config.LLM.Enabled && e.llmProvider != nil {
			llmEngine := llm.NewEngine(&e.config.LLM, e.llmProvider, e.logger)
			var enhancedFindings []*finding.Finding
			
			for _, f := range fileFindings {
				enhanced, err := llmEngine.Analyze(ctx, f, e.ruleIndex)
				if err != nil {
					e.logger.Debugw("LLM analysis failed for finding", "error", err)
					enhancedFindings = append(enhancedFindings, f)
				} else if enhanced != nil {
					enhancedFindings = append(enhancedFindings, enhanced)
				}
			}
			fileFindings = enhancedFindings
		}

		if len(fileFindings) > 0 {
			results <- fileFindings
		}

		e.logger.Debugw("Processed file", "path", f.Path, "findings", len(fileFindings), "duration", time.Since(start))
	}
}

// generateReport outputs the formatted results
func (e *Engine) generateReport(ctx context.Context, findings []*finding.Finding) error {
	format := e.config.Output.Format
	reporter := e.reporters.Get(format)
	
	if reporter == nil {
		return fmt.Errorf("unsupported output format: %s", format)
	}

	var w io.Writer = os.Stdout
	if e.config.Output.File != "" {
		f, err := os.Create(e.config.Output.File)
		if err != nil {
			return fmt.Errorf("failed to create output file: %w", err)
		}
		defer f.Close()
		w = f
	}

	return reporter.Generate(ctx, findings, w)
}
