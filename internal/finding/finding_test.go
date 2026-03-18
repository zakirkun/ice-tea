package finding

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeverityOrder(t *testing.T) {
	assert.True(t, SeverityOrder(SeverityCritical) > SeverityOrder(SeverityHigh))
	assert.True(t, SeverityOrder(SeverityHigh) > SeverityOrder(SeverityMedium))
	assert.True(t, SeverityOrder(SeverityMedium) > SeverityOrder(SeverityLow))
	assert.True(t, SeverityOrder(SeverityLow) > SeverityOrder(SeverityInfo))
}

func TestMeetsThreshold(t *testing.T) {
	assert.True(t, MeetsThreshold(SeverityCritical, SeverityMedium))
	assert.True(t, MeetsThreshold(SeverityMedium, SeverityMedium))
	assert.False(t, MeetsThreshold(SeverityLow, SeverityMedium))
	assert.False(t, MeetsThreshold(SeverityInfo, SeverityMedium))
}

func TestAggregatorDeduplication(t *testing.T) {
	agg := NewAggregator(SeverityInfo)

	// Add finding from pattern engine
	agg.Add(&Finding{
		ID:        "1",
		RuleID:    "TEST-001",
		File:      "main.go",
		StartLine: 10,
		Severity:  SeverityMedium,
		Engines:   []string{EnginePattern},
		Message:   "Static match",
	})

	// Add same finding from taint engine (higher severity)
	agg.Add(&Finding{
		ID:        "2",
		RuleID:    "TEST-001",
		File:      "main.go",
		StartLine: 10,
		Severity:  SeverityHigh, // higher severity should win
		Engines:   []string{EngineTaint},
		Message:   "Data flow match",
		DataFlow:  []string{"source -> sink"},
	})

	// Add different finding
	agg.Add(&Finding{
		ID:        "3",
		RuleID:    "TEST-002",
		File:      "main.go",
		StartLine: 20,
		Severity:  SeverityCritical,
		Engines:   []string{EnginePattern},
	})

	results := agg.Results()

	// Should deduplicate 3 findings down to 2
	assert.Len(t, results, 2)

	// Array is sorted by severity (Critical first, then High)
	assert.Equal(t, "TEST-002", results[0].RuleID)
	assert.Equal(t, SeverityCritical, results[0].Severity)

	// Check the deduplicated finding properties
	merged := results[1]
	assert.Equal(t, "TEST-001", merged.RuleID)
	assert.Equal(t, SeverityHigh, merged.Severity) // Higher severity won
	assert.Contains(t, merged.Engines, EnginePattern)
	assert.Contains(t, merged.Engines, EngineTaint)
	assert.Equal(t, []string{"source -> sink"}, merged.DataFlow)
}

func TestAggregatorThreshold(t *testing.T) {
	agg := NewAggregator(SeverityHigh)

	agg.Add(&Finding{ID: "1", RuleID: "R1", Severity: SeverityCritical}) // kept
	agg.Add(&Finding{ID: "2", RuleID: "R2", Severity: SeverityHigh})     // kept
	agg.Add(&Finding{ID: "3", RuleID: "R3", Severity: SeverityMedium})   // filtered
	agg.Add(&Finding{ID: "4", RuleID: "R4", Severity: SeverityLow})      // filtered

	results := agg.Results()
	assert.Len(t, results, 2)
}

func TestSummarize(t *testing.T) {
	findings := []*Finding{
		{Severity: SeverityCritical},
		{Severity: SeverityCritical},
		{Severity: SeverityHigh},
		{Severity: SeverityMedium},
		{Severity: SeverityMedium},
		{Severity: SeverityMedium},
		{Severity: SeverityInfo},
	}

	summary := Summarize(findings)
	assert.Equal(t, 7, summary.Total)
	assert.Equal(t, 2, summary.Critical)
	assert.Equal(t, 1, summary.High)
	assert.Equal(t, 3, summary.Medium)
	assert.Equal(t, 0, summary.Low)
	assert.Equal(t, 1, summary.Info)
}
