package finding

// Severity levels for findings
const (
	SeverityCritical = "critical"
	SeverityHigh     = "high"
	SeverityMedium   = "medium"
	SeverityLow      = "low"
	SeverityInfo     = "info"
)

// Confidence levels for findings
const (
	ConfidenceHigh   = "high"
	ConfidenceMedium = "medium"
	ConfidenceLow    = "low"
)

// Engine identifiers
const (
	EnginePattern = "pattern"
	EngineTaint   = "taint"
	EngineLLM     = "llm"
)

// Finding represents a detected security vulnerability
type Finding struct {
	ID          string   `json:"id"`
	RuleID      string   `json:"ruleId"`
	Type        string   `json:"type"`
	Severity    string   `json:"severity"`
	Confidence  string   `json:"confidence"`
	CWE         []string `json:"cwe,omitempty"`
	OWASP       []string `json:"owasp,omitempty"`
	File        string   `json:"file"`
	StartLine   int      `json:"startLine"`
	EndLine     int      `json:"endLine"`
	StartColumn int      `json:"startColumn,omitempty"`
	EndColumn   int      `json:"endColumn,omitempty"`
	CodeSnippet string   `json:"codeSnippet,omitempty"`
	Message     string   `json:"message"`
	DataFlow    []string `json:"dataFlow,omitempty"`
	Fix         string   `json:"fix,omitempty"`
	FixCode     string   `json:"fixCode,omitempty"`
	Engines     []string `json:"engines"`
}

// SeverityOrder returns numeric priority (higher = more severe)
func SeverityOrder(sev string) int {
	switch sev {
	case SeverityCritical:
		return 5
	case SeverityHigh:
		return 4
	case SeverityMedium:
		return 3
	case SeverityLow:
		return 2
	case SeverityInfo:
		return 1
	default:
		return 0
	}
}

// MeetsThreshold checks if a severity meets the minimum threshold
func MeetsThreshold(severity, threshold string) bool {
	return SeverityOrder(severity) >= SeverityOrder(threshold)
}

// Aggregator collects and deduplicates findings from multiple engines
type Aggregator struct {
	findings  []*Finding
	threshold string // minimum severity
}

// NewAggregator creates a new finding aggregator
func NewAggregator(severityThreshold string) *Aggregator {
	return &Aggregator{
		threshold: severityThreshold,
	}
}

// Add adds a finding to the aggregator
func (a *Aggregator) Add(f *Finding) {
	if !MeetsThreshold(f.Severity, a.threshold) {
		return // below threshold
	}
	a.findings = append(a.findings, f)
}

// Results returns all aggregated findings, deduplicated and sorted
func (a *Aggregator) Results() []*Finding {
	deduped := a.deduplicate()
	sortBySeverity(deduped)
	return deduped
}

// deduplicate merges findings with matching file + line range + rule
func (a *Aggregator) deduplicate() []*Finding {
	type key struct {
		RuleID string
		File   string
		Line   int
	}

	merged := make(map[key]*Finding)
	var order []key

	for _, f := range a.findings {
		k := key{RuleID: f.RuleID, File: f.File, Line: f.StartLine}

		if existing, ok := merged[k]; ok {
			// Merge: take highest severity
			if SeverityOrder(f.Severity) > SeverityOrder(existing.Severity) {
				existing.Severity = f.Severity
			}
			// Add engine
			existing.Engines = appendUnique(existing.Engines, f.Engines...)
			// Take LLM explanation if available
			if f.Fix != "" && existing.Fix == "" {
				existing.Fix = f.Fix
			}
			if f.FixCode != "" && existing.FixCode == "" {
				existing.FixCode = f.FixCode
			}
			// Merge data flow
			if len(f.DataFlow) > 0 && len(existing.DataFlow) == 0 {
				existing.DataFlow = f.DataFlow
			}
		} else {
			merged[k] = f
			order = append(order, k)
		}
	}

	var result []*Finding
	for _, k := range order {
		result = append(result, merged[k])
	}
	return result
}

// sortBySeverity sorts findings by severity (critical first)
func sortBySeverity(findings []*Finding) {
	for i := 0; i < len(findings); i++ {
		for j := i + 1; j < len(findings); j++ {
			if SeverityOrder(findings[j].Severity) > SeverityOrder(findings[i].Severity) {
				findings[i], findings[j] = findings[j], findings[i]
			}
		}
	}
}

func appendUnique(slice []string, items ...string) []string {
	seen := make(map[string]bool)
	for _, s := range slice {
		seen[s] = true
	}
	for _, item := range items {
		if !seen[item] {
			slice = append(slice, item)
			seen[item] = true
		}
	}
	return slice
}

// Summary provides aggregate statistics
type Summary struct {
	Total    int `json:"total"`
	Critical int `json:"critical"`
	High     int `json:"high"`
	Medium   int `json:"medium"`
	Low      int `json:"low"`
	Info     int `json:"info"`
}

// Summarize creates a summary from findings
func Summarize(findings []*Finding) Summary {
	s := Summary{Total: len(findings)}
	for _, f := range findings {
		switch f.Severity {
		case SeverityCritical:
			s.Critical++
		case SeverityHigh:
			s.High++
		case SeverityMedium:
			s.Medium++
		case SeverityLow:
			s.Low++
		case SeverityInfo:
			s.Info++
		}
	}
	return s
}
