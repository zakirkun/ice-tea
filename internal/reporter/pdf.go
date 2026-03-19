package reporter

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/go-pdf/fpdf"

	"github.com/zakirkun/ice-tea/internal/finding"
)

// PDFReporter generates a structured PDF security report
type PDFReporter struct{}

// NewPDFReporter creates a new PDF reporter
func NewPDFReporter() *PDFReporter {
	return &PDFReporter{}
}

// Format implements Reporter
func (p *PDFReporter) Format() string { return "pdf" }

// severityColor maps severity to RGB color components
func severityColor(severity string) (r, g, b int) {
	switch strings.ToLower(severity) {
	case "critical":
		return 220, 38, 38 // red-600
	case "high":
		return 234, 88, 12 // orange-600
	case "medium":
		return 202, 138, 4 // yellow-600
	case "low":
		return 22, 163, 74 // green-600
	default:
		return 100, 116, 139 // slate-500 (info)
	}
}

// Generate implements Reporter — produces a PDF security report
func (p *PDFReporter) Generate(ctx context.Context, findings []*finding.Finding, w io.Writer) error {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.SetAuthor("Ice Tea Security Scanner", false)
	pdf.SetTitle("Security Scan Report", false)
	pdf.SetCreator("Ice Tea Security Scanner", false)
	pdf.SetCreationDate(time.Now())

	// Compute summary
	summary := finding.Summarize(findings)

	// ── Cover Page ──────────────────────────────────────────────────────────────
	pdf.AddPage()
	pdf.SetMargins(20, 20, 20)

	// Title
	pdf.SetFont("Helvetica", "B", 28)
	pdf.SetTextColor(30, 64, 175) // blue-700
	pdf.CellFormat(0, 20, "Ice Tea Security Scanner", "", 1, "C", false, 0, "")

	pdf.SetFont("Helvetica", "", 14)
	pdf.SetTextColor(71, 85, 105) // slate-600
	pdf.CellFormat(0, 10, "Security Scan Report", "", 1, "C", false, 0, "")

	pdf.Ln(10)
	pdf.SetDrawColor(203, 213, 225) // slate-300
	pdf.SetLineWidth(0.3)
	pdf.Line(20, pdf.GetY(), 190, pdf.GetY())
	pdf.Ln(10)

	// Scan metadata
	pdf.SetFont("Helvetica", "B", 11)
	pdf.SetTextColor(30, 41, 59)
	pdf.CellFormat(40, 8, "Scan Date:", "", 0, "L", false, 0, "")
	pdf.SetFont("Helvetica", "", 11)
	pdf.CellFormat(0, 8, time.Now().Format("2006-01-02 15:04:05 UTC"), "", 1, "L", false, 0, "")

	pdf.SetFont("Helvetica", "B", 11)
	pdf.CellFormat(40, 8, "Tool Version:", "", 0, "L", false, 0, "")
	pdf.SetFont("Helvetica", "", 11)
	pdf.CellFormat(0, 8, "Ice Tea v1.0.0", "", 1, "L", false, 0, "")

	pdf.SetFont("Helvetica", "B", 11)
	pdf.CellFormat(40, 8, "Total Findings:", "", 0, "L", false, 0, "")
	pdf.SetFont("Helvetica", "", 11)
	pdf.CellFormat(0, 8, fmt.Sprintf("%d", summary.Total), "", 1, "L", false, 0, "")

	pdf.Ln(15)

	// ── Executive Summary ────────────────────────────────────────────────────────
	pdf.SetFont("Helvetica", "B", 16)
	pdf.SetTextColor(30, 41, 59)
	pdf.CellFormat(0, 12, "Executive Summary", "", 1, "L", false, 0, "")

	pdf.SetFont("Helvetica", "", 11)
	pdf.SetTextColor(71, 85, 105)
	pdf.MultiCell(0, 7, "The following table summarizes the security findings identified during this scan, grouped by severity level.", "", "L", false)
	pdf.Ln(5)

	// Summary table header
	pdf.SetFillColor(30, 64, 175)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetFont("Helvetica", "B", 11)
	pdf.CellFormat(80, 10, "Severity", "1", 0, "C", true, 0, "")
	pdf.CellFormat(50, 10, "Count", "1", 0, "C", true, 0, "")
	pdf.CellFormat(40, 10, "Risk Level", "1", 1, "C", true, 0, "")

	// Summary rows
	type summaryRow struct {
		label    string
		count    int
		risk     string
		r, g, b  int
	}
	rows := []summaryRow{
		{"Critical", summary.Critical, "Immediate Action", 220, 38, 38},
		{"High", summary.High, "Urgent", 234, 88, 12},
		{"Medium", summary.Medium, "Important", 202, 138, 4},
		{"Low", summary.Low, "Moderate", 22, 163, 74},
		{"Info", summary.Info, "Informational", 100, 116, 139},
	}

	for _, row := range rows {
		pdf.SetFillColor(row.r, row.g, row.b)
		pdf.SetTextColor(255, 255, 255)
		pdf.SetFont("Helvetica", "B", 10)
		pdf.CellFormat(80, 9, row.label, "1", 0, "L", true, 0, "")

		pdf.SetFillColor(248, 250, 252)
		pdf.SetTextColor(30, 41, 59)
		pdf.SetFont("Helvetica", "", 10)
		pdf.CellFormat(50, 9, fmt.Sprintf("%d", row.count), "1", 0, "C", true, 0, "")
		pdf.CellFormat(40, 9, row.risk, "1", 1, "C", true, 0, "")
	}

	// Total row
	pdf.SetFillColor(30, 41, 59)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetFont("Helvetica", "B", 10)
	pdf.CellFormat(80, 9, "TOTAL", "1", 0, "L", true, 0, "")
	pdf.CellFormat(50, 9, fmt.Sprintf("%d", summary.Total), "1", 0, "C", true, 0, "")
	pdf.CellFormat(40, 9, "", "1", 1, "C", true, 0, "")

	// ── Findings Detail Pages ────────────────────────────────────────────────────
	if len(findings) == 0 {
		pdf.Ln(20)
		pdf.SetFont("Helvetica", "B", 14)
		pdf.SetTextColor(22, 163, 74)
		pdf.CellFormat(0, 12, "✓ No security findings detected", "", 1, "C", false, 0, "")
		pdf.SetFont("Helvetica", "", 11)
		pdf.SetTextColor(71, 85, 105)
		pdf.MultiCell(0, 7, "The scan completed successfully with no vulnerabilities found above the configured severity threshold.", "", "C", false)
	} else {
		// Group findings by severity
		order := []string{"critical", "high", "medium", "low", "info"}
		groups := make(map[string][]*finding.Finding)
		for _, f := range findings {
			groups[strings.ToLower(f.Severity)] = append(groups[strings.ToLower(f.Severity)], f)
		}

		for _, sev := range order {
			sevFindings := groups[sev]
			if len(sevFindings) == 0 {
				continue
			}

			pdf.AddPage()
			r, g, b := severityColor(sev)

			// Section header
			pdf.SetFillColor(r, g, b)
			pdf.SetTextColor(255, 255, 255)
			pdf.SetFont("Helvetica", "B", 14)
			pdf.CellFormat(0, 12, fmt.Sprintf("%s Findings (%d)", strings.Title(sev), len(sevFindings)), "", 1, "L", true, 0, "")
			pdf.Ln(5)

			for i, f := range sevFindings {
				// Check if we need a new page
				if pdf.GetY() > 250 {
					pdf.AddPage()
				}

				// Finding card header
				pdf.SetFillColor(r, g, b)
				pdf.SetTextColor(255, 255, 255)
				pdf.SetFont("Helvetica", "B", 10)
				pdf.CellFormat(0, 8, fmt.Sprintf("[%d] %s — %s", i+1, f.RuleID, f.Type), "", 1, "L", true, 0, "")

				// File + line
				pdf.SetFillColor(241, 245, 249)
				pdf.SetTextColor(30, 41, 59)
				pdf.SetFont("Helvetica", "", 9)
				pdf.CellFormat(25, 7, "File:", "LTB", 0, "L", true, 0, "")
				pdf.SetFont("Courier", "", 9)
				pdf.CellFormat(0, 7, fmt.Sprintf("%s:%d", f.File, f.StartLine), "RTB", 1, "L", true, 0, "")

				// CWE + OWASP
				if len(f.CWE) > 0 || len(f.OWASP) > 0 {
					pdf.SetFont("Helvetica", "", 9)
					pdf.SetFillColor(241, 245, 249)
					tags := strings.Join(f.CWE, ", ")
					if len(f.OWASP) > 0 {
						if tags != "" {
							tags += " | "
						}
						tags += strings.Join(f.OWASP, ", ")
					}
					pdf.CellFormat(25, 7, "Tags:", "LTB", 0, "L", true, 0, "")
					pdf.CellFormat(0, 7, tags, "RTB", 1, "L", true, 0, "")
				}

				// Description
				pdf.SetFont("Helvetica", "", 9)
				pdf.SetTextColor(30, 41, 59)
				pdf.SetFillColor(255, 255, 255)
				pdf.CellFormat(25, 7, "Description:", "LTB", 0, "L", false, 0, "")

				descText := f.Message
				if descText == "" {
					descText = "No description available."
				}
				pdf.MultiCell(0, 7, descText, "RTB", "L", false)

				// Code snippet
				if f.CodeSnippet != "" {
					pdf.SetFont("Courier", "", 8)
					pdf.SetFillColor(30, 41, 59)
					pdf.SetTextColor(248, 250, 252)
					snippet := f.CodeSnippet
					if len(snippet) > 200 {
						snippet = snippet[:200] + "..."
					}
					pdf.MultiCell(0, 5, snippet, "1", "L", true)
				}

				// Fix suggestion
				if f.Fix != "" {
					pdf.SetFont("Helvetica", "I", 8)
					pdf.SetFillColor(240, 253, 244)
					pdf.SetTextColor(22, 101, 52)
					pdf.MultiCell(0, 5, "Remediation: "+f.Fix, "1", "L", true)
				}

				pdf.Ln(4)
			}
		}
	}

	// ── Footer on all pages ──────────────────────────────────────────────────────
	pdf.SetFooterFunc(func() {
		pdf.SetY(-15)
		pdf.SetFont("Helvetica", "I", 8)
		pdf.SetTextColor(148, 163, 184)
		pdf.CellFormat(0, 10,
			fmt.Sprintf("Ice Tea Security Scanner | Generated %s | Page %d",
				time.Now().Format("2006-01-02"),
				pdf.PageNo()),
			"", 0, "C", false, 0, "")
	})

	// Write to io.Writer
	var buf bytes.Buffer
	if err := pdf.Output(&buf); err != nil {
		return fmt.Errorf("PDF generation failed: %w", err)
	}
	_, err := w.Write(buf.Bytes())
	return err
}
