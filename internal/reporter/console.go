package reporter

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/fatih/color"

	"github.com/zakirkun/ice-tea/internal/finding"
)

// ConsoleReporter generates human-readable terminal output
type ConsoleReporter struct {
	useColor bool
}

// NewConsoleReporter creates a new Console reporter
func NewConsoleReporter(useColor bool) *ConsoleReporter {
	return &ConsoleReporter{
		useColor: useColor,
	}
}

// Format returns the format identifier
func (r *ConsoleReporter) Format() string {
	return "console"
}

// Generate writes the human-readable console report
func (r *ConsoleReporter) Generate(ctx context.Context, findings []*finding.Finding, w io.Writer) error {
	color.NoColor = !r.useColor

	// Color definitions
	red := color.New(color.FgRed, color.Bold)
	yellow := color.New(color.FgYellow, color.Bold)
	cyan := color.New(color.FgCyan)
	whiteBold := color.New(color.FgWhite, color.Bold)
	gray := color.New(color.FgHiBlack)
	green := color.New(color.FgGreen)
	magenta := color.New(color.FgHiMagenta)

	if len(findings) == 0 {
		green.Fprintln(w, "\n✓ No vulnerabilities found. Your code looks secure!")
		return nil
	}

	fmt.Fprintln(w, "\nFindings")
	fmt.Fprintln(w, strings.Repeat("═", 80))

	// Group findings by file
	byFile := make(map[string][]*finding.Finding)
	for _, f := range findings {
		byFile[f.File] = append(byFile[f.File], f)
	}

	for file, fileFindings := range byFile {
		whiteBold.Fprintf(w, "FILE: %s\n", file)
		fmt.Fprintln(w, strings.Repeat("─", 80))

		for _, f := range fileFindings {
			// Format severity
			sevOut := f.Severity
			switch f.Severity {
			case finding.SeverityCritical, finding.SeverityHigh:
				sevOut = red.Sprintf("[%s]", strings.ToUpper(f.Severity))
			case finding.SeverityMedium:
				sevOut = yellow.Sprintf("[%s]", strings.ToUpper(f.Severity))
			case finding.SeverityLow, finding.SeverityInfo:
				sevOut = cyan.Sprintf("[%s]", strings.ToUpper(f.Severity))
			}

			// Format CWE/OWASP tags
			var tags []string
			tags = append(tags, f.CWE...)
			tags = append(tags, f.OWASP...)
			tagStr := ""
			if len(tags) > 0 {
				tagStr = gray.Sprintf("(%s)", strings.Join(tags, ", "))
			}

			// Main finding header
			fmt.Fprintf(w, "  %s %s %s\n", sevOut, whiteBold.Sprint(f.Type), tagStr)
			cyan.Fprintf(w, "  Line: %d | Rule: %s | Engines: %s\n", 
				f.StartLine, f.RuleID, strings.Join(f.Engines, ", "))

			// Description
			fmt.Fprintf(w, "  %s\n", f.Message)

			// Code snippet
			if f.CodeSnippet != "" {
				fmt.Fprintln(w, "")
				gray.Fprintf(w, "  % 4d | %s\n", f.StartLine, strings.TrimSpace(f.CodeSnippet))
			}

			// Data flow (if any)
			if len(f.DataFlow) > 0 {
				fmt.Fprintln(w, "")
				magenta.Fprintln(w, "  Data Flow:")
				for i, step := range f.DataFlow {
					gray.Fprintf(w, "    %d. %s\n", i+1, step)
				}
			}

			// LLM Fix (if any)
			if f.Fix != "" {
				fmt.Fprintln(w, "")
				green.Fprintln(w, "  Suggested Fix (LLM):")
				sentences := strings.Split(f.Fix, ". ")
				for _, s := range sentences {
					if s != "" {
						fmt.Fprintf(w, "    • %s\n", strings.TrimSpace(s))
					}
				}
			}

			fmt.Fprintln(w, strings.Repeat("-", 40))
		}
		fmt.Fprintln(w, "")
	}

	// Print Summary
	summary := finding.Summarize(findings)
	fmt.Fprintln(w, strings.Repeat("═", 80))
	whiteBold.Fprintln(w, "Scan Summary")
	fmt.Fprintf(w, "  Total Findings: %d\n", summary.Total)
	
	if summary.Critical > 0 {
		red.Fprintf(w, "  Critical: %d\n", summary.Critical)
	}
	if summary.High > 0 {
		red.Fprintf(w, "  High:     %d\n", summary.High)
	}
	if summary.Medium > 0 {
		yellow.Fprintf(w, "  Medium:   %d\n", summary.Medium)
	}
	if summary.Low > 0 {
		cyan.Fprintf(w, "  Low:      %d\n", summary.Low)
	}
	if summary.Info > 0 {
		gray.Fprintf(w, "  Info:     %d\n", summary.Info)
	}
	
	return nil
}
