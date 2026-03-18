package llm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/zakirkun/ice-tea/internal/skill"
)

func TestBuildPrompts(t *testing.T) {
	req := AnalysisRequest{
		Rule: &skill.Rule{
			ID:          "TEST-001",
			Name:        "sql-injection",
			Description: "SQL Injection vulnerability",
		},
		File:        "main.go",
		CodeSnippet: `db.Query("SELECT * FROM users WHERE id = " + id)`,
		StartLine:   42,
	}

	sys, user, err := buildPrompts(req)
	require.NoError(t, err)

	assert.Contains(t, sys, "You are Ice Tea")
	assert.Contains(t, sys, "is_vulnerable")

	assert.Contains(t, user, "File: main.go")
	assert.Contains(t, user, "Line: 42")
	assert.Contains(t, user, "TEST-001")
	assert.Contains(t, user, "sql-injection")
	assert.Contains(t, user, "db.Query(\"SELECT * FROM users WHERE id = \" + id)")
}
