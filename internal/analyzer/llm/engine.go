package llm

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/zakirkun/ice-tea/internal/config"
	"github.com/zakirkun/ice-tea/internal/finding"
	"github.com/zakirkun/ice-tea/internal/skill"
)

// Provider defines the interface for LLM backends
type Provider interface {
	// Analyze sends the code snippet and context to the LLM and returns an analysis
	Analyze(ctx context.Context, req AnalysisRequest) (*AnalysisResponse, error)
}

// AnalysisRequest contains the data sent to the LLM
type AnalysisRequest struct {
	Rule        *skill.Rule
	File        string
	CodeSnippet string
	StartLine   int
}

// AnalysisResponse contains the LLM's findings
type AnalysisResponse struct {
	IsVulnerable bool   `json:"is_vulnerable"`
	Confidence   string `json:"confidence"` // high, medium, low
	Explanation  string `json:"explanation"`
	Fix          string `json:"fix"`
	FixCode      string `json:"fix_code"`
}

// Engine wraps LLM providers to perform deep reasoning on findings
type Engine struct {
	provider Provider
	logger   *zap.SugaredLogger
	config   *config.LLMConfig
}

// NewEngine creates a new LLM reasoning engine
func NewEngine(cfg *config.LLMConfig, p Provider, logger *zap.SugaredLogger) *Engine {
	return &Engine{
		provider: p,
		logger:   logger,
		config:   cfg,
	}
}

// Analyze performs deep reasoning on existing findings to filter false positives
// and generate fix suggestions.
func (e *Engine) Analyze(ctx context.Context, f *finding.Finding, index *skill.RuleIndex) (*finding.Finding, error) {
	if !e.config.Enabled {
		return f, nil
	}

	rule := index.ByID[f.RuleID]
	if rule == nil {
		e.logger.Debugw("Rule not found for LLM analysis", "rule_id", f.RuleID)
		return f, nil
	}

	req := AnalysisRequest{
		Rule:        rule,
		File:        f.File,
		CodeSnippet: f.CodeSnippet,
		StartLine:   f.StartLine,
	}

	e.logger.Debugw("Requesting LLM analysis", "rule", f.RuleID, "file", f.File, "line", f.StartLine)

	resp, err := e.provider.Analyze(ctx, req)
	if err != nil {
		e.logger.Warnw("LLM analysis failed", "error", err)
		return f, err // return original finding on error
	}

	// If LLM says it's not vulnerable (high confidence false positive), we can drop it
	if !resp.IsVulnerable && resp.Confidence == finding.ConfidenceHigh {
		e.logger.Debugw("LLM determined false positive", "rule", f.RuleID, "file", f.File)
		return nil, nil // drop finding
	}

	// Enhance finding with LLM insights
	enhanced := *f // copy
	enhanced.Engines = append(enhanced.Engines, finding.EngineLLM)

	if resp.Explanation != "" {
		// Append LLM explanation if it provides new context
		enhanced.Message = fmt.Sprintf("%s\n\nAI Analysis: %s", enhanced.Message, resp.Explanation)
	}

	if resp.Fix != "" {
		enhanced.Fix = resp.Fix
	}

	if resp.FixCode != "" {
		enhanced.FixCode = resp.FixCode
	}

	// We can upgrade/downgrade confidence based on LLM
	if resp.Confidence != "" {
		enhanced.Confidence = resp.Confidence
	}

	return &enhanced, nil
}
