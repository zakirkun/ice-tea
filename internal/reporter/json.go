package reporter

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/zakirkun/ice-tea/internal/finding"
)

// JSONReporter generates raw JSON output
type JSONReporter struct{}

// NewJSONReporter creates a new JSON reporter
func NewJSONReporter() *JSONReporter {
	return &JSONReporter{}
}

// Format returns the format identifier
func (r *JSONReporter) Format() string {
	return "json"
}

// jsonReport represents the top-level JSON structure
type jsonReport struct {
	Summary  finding.Summary    `json:"summary"`
	Findings []*finding.Finding `json:"findings"`
}

// Generate writes the JSON report
func (r *JSONReporter) Generate(ctx context.Context, findings []*finding.Finding, w io.Writer) error {
	report := jsonReport{
		Summary:  finding.Summarize(findings),
		Findings: findings,
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(report); err != nil {
		return fmt.Errorf("failed to encode JSON report: %w", err)
	}

	return nil
}

// GitLabReporter generates GitLab SAST JSON output
type GitLabReporter struct{}

// NewGitLabReporter creates a new GitLab reporter
func NewGitLabReporter() *GitLabReporter {
	return &GitLabReporter{}
}

// Format returns the format identifier
func (r *GitLabReporter) Format() string {
	return "gitlab"
}

// Generate writes the GitLab SAST format JSON report
// Reference: https://docs.gitlab.com/ee/user/application_security/sast/analyzers.html#sast-report-json-format
func (r *GitLabReporter) Generate(ctx context.Context, findings []*finding.Finding, w io.Writer) error {
	vulnerabilities := make([]map[string]interface{}, 0, len(findings))

	for _, f := range findings {
		// Convert severity to GitLab SAST format
		severity := convertGitLabSeverity(f.Severity)

		vuln := map[string]interface{}{
			"id":          f.ID,
			"category":    "sast",
			"name":        f.Type,
			"message":     f.Message,
			"description": f.Message,
			"severity":    severity,
			"confidence":  convertGitLabConfidence(f.Confidence),
			"scanner": map[string]interface{}{
				"id":   "ice-tea",
				"name": "Ice Tea Security Scanner",
			},
			"location": map[string]interface{}{
				"file":       f.File,
				"start_line": f.StartLine,
				"end_line":   f.EndLine,
			},
			"identifiers": []map[string]interface{}{
				{
					"type":  "ice-tea-rule",
					"name":  fmt.Sprintf("Ice Tea Rule %s", f.RuleID),
					"value": f.RuleID,
				},
			},
		}

		// Add CWE if available
		for _, cwe := range f.CWE {
			if strings.HasPrefix(cwe, "CWE-") {
				vuln["identifiers"] = append(vuln["identifiers"].([]map[string]interface{}), map[string]interface{}{
					"type":  "cwe",
					"name":  cwe,
					"value": strings.TrimPrefix(cwe, "CWE-"),
				})
			}
		}

		// Add solution if LLM fix is available
		if f.Fix != "" {
			vuln["solution"] = f.Fix
		}

		vulnerabilities = append(vulnerabilities, vuln)
	}

	report := map[string]interface{}{
		"version": "14.1.2",
		"vulnerabilities": vulnerabilities,
		"remediations":    []interface{}{},
		"scan": map[string]interface{}{
			"analyzer": map[string]interface{}{
				"id":      "ice-tea",
				"name":    "Ice Tea Security Scanner",
				"version": "latest",
				"vendor": map[string]interface{}{
					"name": "Zakirkun",
				},
			},
			"scanner": map[string]interface{}{
				"id":      "ice-tea",
				"name":    "Ice Tea Security Scanner",
				"version": "latest",
				"vendor": map[string]interface{}{
					"name": "Zakirkun",
				},
			},
			"type":       "sast",
			"start_time": "2024-01-01T00:00:00Z", // Timestamp could be dynamic
			"end_time":   "2024-01-01T00:00:00Z", // Timestamp could be dynamic
			"status":     "success",
		},
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(report); err != nil {
		return fmt.Errorf("failed to encode GitLab SAST report: %w", err)
	}

	return nil
}

func convertGitLabSeverity(severity string) string {
	switch severity {
	case finding.SeverityCritical:
		return "Critical"
	case finding.SeverityHigh:
		return "High"
	case finding.SeverityMedium:
		return "Medium"
	case finding.SeverityLow:
		return "Low"
	case finding.SeverityInfo:
		return "Info"
	default:
		return "Unknown"
	}
}

func convertGitLabConfidence(confidence string) string {
	switch confidence {
	case finding.ConfidenceHigh:
		return "High"
	case finding.ConfidenceMedium:
		return "Medium"
	case finding.ConfidenceLow:
		return "Low"
	default:
		return "Unknown"
	}
}
