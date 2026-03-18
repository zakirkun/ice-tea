package taint

import (
	"github.com/zakirkun/ice-tea/internal/finding"
	"github.com/zakirkun/ice-tea/internal/parser"
	"go.uber.org/zap"
)

// Tracker tracks data flow for taint analysis
// NOTE: Full inter-procedural taint analysis requires SSA form, which is complex.
// This implementation provides a structural placeholder that tracks direct assignments within a single function.
type Tracker struct {
	logger *zap.SugaredLogger
}

// NewTracker creates a new taint tracker
func NewTracker(logger *zap.SugaredLogger) *Tracker {
	return &Tracker{
		logger: logger,
	}
}

// Analyze performs simple intra-procedural taint tracking on a parsed file
func (t *Tracker) Analyze(result *parser.ParseResult) []*finding.Finding {
	if result == nil || result.Root == nil {
		return nil
	}

	var findings []*finding.Finding
	
	// Phase 5 skeletal implementation:
	// Find all function declarations to run intra-procedural analysis
	funcs := parser.FindAll(result.Root, "function_declaration")
	
	for _, fn := range funcs {
		fnFindings := t.analyzeFunction(fn, result.FilePath)
		findings = append(findings, fnFindings...)
	}

	// Also check methods
	methods := parser.FindAll(result.Root, "method_declaration")
	for _, fn := range methods {
		fnFindings := t.analyzeFunction(fn, result.FilePath)
		findings = append(findings, fnFindings...)
	}

	return findings
}

func (t *Tracker) analyzeFunction(fn *parser.Node, filepath string) []*finding.Finding {
	// 1. Identify taint sources (e.g., http.Request.FormValue)
	// 2. Track tainted variables down the AST
	// 3. Identify taint sinks (e.g., db.Query, exec.Command)
	// 4. Report if tainted variable reaches a sink

	// Placeholder returning nothing for now.
	// Will be fully implemented if requested by user in future iterations.
	return nil
}

// DefaultSources defines common taint sources like HTTP requests, ENV vars
var DefaultSources = []string{
	"r.FormValue",
	"r.URL.Query().Get",
	"os.Getenv",
	"bufio.NewReader(os.Stdin).ReadString",
}

// DefaultSinks defines common dangerous functions where taint matters
var DefaultSinks = []string{
	"db.Query",
	"exec.Command",
	"os.OpenFile",
}
