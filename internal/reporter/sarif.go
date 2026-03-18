package reporter

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/zakirkun/ice-tea/internal/finding"
)

// SarifReporter generates SARIF v2.1.0 JSON output
type SarifReporter struct{}

// NewSarifReporter creates a new SARIF reporter
func NewSarifReporter() *SarifReporter {
	return &SarifReporter{}
}

// Format returns the format identifier
func (r *SarifReporter) Format() string {
	return "sarif"
}

// Generate writes the SARIF v2.1.0 format report
func (r *SarifReporter) Generate(ctx context.Context, findings []*finding.Finding, w io.Writer) error {
	// Group findings by rule definition
	ruleMap := make(map[string]*finding.Finding)
	for _, f := range findings {
		if _, exists := ruleMap[f.RuleID]; !exists {
			ruleMap[f.RuleID] = f
		}
	}

	// Create SARIF rules array
	var rules []map[string]interface{}
	ruleIndexMap := make(map[string]int) // Maps RuleID to its index in the rules array

	for _, f := range ruleMap {
		ruleIndexMap[f.RuleID] = len(rules)

		var tags []string
		if len(f.CWE) > 0 {
			tags = append(tags, f.CWE...)
		}
		if len(f.OWASP) > 0 {
			tags = append(tags, f.OWASP...)
		}

		rule := map[string]interface{}{
			"id":   f.RuleID,
			"name": f.Type,
			"shortDescription": map[string]interface{}{
				"text": f.Message,
			},
			"fullDescription": map[string]interface{}{
				"text": f.Message,
			},
			"defaultConfiguration": map[string]interface{}{
				"level": convertSarifLevel(f.Severity),
			},
			"properties": map[string]interface{}{
				"tags": tags,
			},
		}

		rules = append(rules, rule)
	}

	// Create SARIF results array
	var results []map[string]interface{}

	for _, f := range findings {
		result := map[string]interface{}{
			"ruleId":    f.RuleID,
			"ruleIndex": ruleIndexMap[f.RuleID],
			"level":     convertSarifLevel(f.Severity),
			"message": map[string]interface{}{
				"text": f.Message,
			},
			"locations": []map[string]interface{}{
				{
					"physicalLocation": map[string]interface{}{
						"artifactLocation": map[string]interface{}{
							"uri": f.File,
						},
						"region": map[string]interface{}{
							"startLine": f.StartLine,
						},
					},
				},
			},
		}

		// Add column info if available
		if f.StartColumn > 0 {
			loc := result["locations"].([]map[string]interface{})[0]["physicalLocation"].(map[string]interface{})["region"].(map[string]interface{})
			loc["startColumn"] = f.StartColumn
		}

		// Add code snippet
		if f.CodeSnippet != "" {
			loc := result["locations"].([]map[string]interface{})[0]["physicalLocation"].(map[string]interface{})["region"].(map[string]interface{})
			loc["snippet"] = map[string]interface{}{
				"text": f.CodeSnippet,
			}
		}

		// Add data flow (codeFlows)
		if len(f.DataFlow) > 0 {
			var threadFlowLocs []map[string]interface{}
			for _, step := range f.DataFlow {
				threadFlowLocs = append(threadFlowLocs, map[string]interface{}{
					"location": map[string]interface{}{
						"message": map[string]interface{}{
							"text": step,
						},
					},
				})
			}
			result["codeFlows"] = []map[string]interface{}{
				{
					"threadFlows": []map[string]interface{}{
						{
							"locations": threadFlowLocs,
						},
					},
				},
			}
		}

		// Add LLM fix (remediation)
		if f.Fix != "" {
			result["fixes"] = []map[string]interface{}{
				{
					"description": map[string]interface{}{
						"text": f.Fix,
					},
					"artifactChanges": []map[string]interface{}{
						{
							"artifactLocation": map[string]interface{}{
								"uri": f.File,
							},
							"replacements": []map[string]interface{}{
                                // Placeholder for actual replacement struct
                                // A real implementation would parse f.FixCode
							},
						},
					},
				},
			}
		}

		results = append(results, result)
	}

	// Build final SARIF structure
	sarif := map[string]interface{}{
		"$schema": "https://raw.githubusercontent.com/oasis-tcs/sarif-spec/master/Schemata/sarif-schema-2.1.0.json",
		"version": "2.1.0",
		"runs": []map[string]interface{}{
			{
				"tool": map[string]interface{}{
					"driver": map[string]interface{}{
						"name":            "Ice Tea",
						"informationUri":  "https://github.com/zakirkun/ice-tea",
						"semanticVersion": "1.0.0",
						"rules":           rules,
					},
				},
				"results": results,
			},
		},
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(sarif); err != nil {
		return fmt.Errorf("failed to encode SARIF report: %w", err)
	}

	return nil
}

func convertSarifLevel(severity string) string {
	switch severity {
	case finding.SeverityCritical, finding.SeverityHigh:
		return "error"
	case finding.SeverityMedium:
		return "warning"
	case finding.SeverityLow, finding.SeverityInfo:
		return "note"
	default:
		return "none"
	}
}
