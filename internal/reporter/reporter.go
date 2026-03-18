package reporter

import (
	"context"
	"io"

	"github.com/zakirkun/ice-tea/internal/finding"
)

// Reporter is the common interface for all output generators
type Reporter interface {
	// Generate writes the scan report to the provided writer
	Generate(ctx context.Context, findings []*finding.Finding, w io.Writer) error

	// Format returns the identifier for this reporter (e.g., "sarif")
	Format() string
}

// Registry holds all registered reporters
type Registry struct {
	reporters map[string]Reporter
}

// NewRegistry creates a new reporter registry
func NewRegistry() *Registry {
	return &Registry{
		reporters: make(map[string]Reporter),
	}
}

// Register adds a reporter to the registry
func (r *Registry) Register(reporter Reporter) {
	r.reporters[reporter.Format()] = reporter
}

// Get finds a reporter by format name
func (r *Registry) Get(format string) Reporter {
	return r.reporters[format]
}
