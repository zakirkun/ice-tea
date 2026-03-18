package pattern

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/zakirkun/ice-tea/internal/parser/goparser"
	"github.com/zakirkun/ice-tea/internal/skill"
)

func testLogger() *zap.SugaredLogger {
	logger, _ := zap.NewDevelopment()
	return logger.Sugar()
}

func TestMatcherRegex(t *testing.T) {
	rule := &skill.Rule{
		ID:        "TEST-001",
		Name:      "hardcoded-password",
		Severity:  "high",
		Languages: []string{"go"},
		Patterns: []skill.Pattern{
			{Regex: `(?i)password\s*(:=|=)\s*".+"`},
		},
	}

	matcher := NewMatcher([]*skill.Rule{rule}, testLogger())

	src := []byte(`package main

func main() {
	user := "admin"
	password := "supersecret123" // vulnerable
	
	// safe
	var dynPassword string
}`)

	// Using GoParser to get the ParseResult
	gp := goparser.New()
	result, err := gp.Parse("main.go", src)
	require.NoError(t, err)

	findings := matcher.Analyze(result)

	assert.Len(t, findings, 1)
	assert.Equal(t, "TEST-001", findings[0].RuleID)
	assert.Equal(t, 5, findings[0].StartLine)
	assert.Contains(t, findings[0].CodeSnippet, "password := \"supersecret123\"")
}

func TestMatcherASTCallExpression(t *testing.T) {
	rule := &skill.Rule{
		ID:        "TEST-002",
		Name:      "command-injection",
		Severity:  "critical",
		Languages: []string{"go"},
		Patterns: []skill.Pattern{
			{
				ASTNodeType: "call_expression",
				Function:    "exec.Command",
			},
		},
	}

	matcher := NewMatcher([]*skill.Rule{rule}, testLogger())

	src := []byte(`package main

import "os/exec"

func run(cmd string) {
	// Vulnerable to command injection
	exec.Command("sh", "-c", cmd).Run()
	
	// Safe
	exec.LookPath("ls")
}`)

	gp := goparser.New()
	result, err := gp.Parse("main.go", src)
	require.NoError(t, err)

	findings := matcher.Analyze(result)

	for _, f := range findings {
		t.Logf("Finding: %s (line: %d)", f.CodeSnippet, f.StartLine)
	}

	assert.Len(t, findings, 1)
	assert.Equal(t, "TEST-002", findings[0].RuleID)
	assert.Equal(t, 7, findings[0].StartLine)
}

func TestMatcherImports(t *testing.T) {
	rule := &skill.Rule{
		ID:        "TEST-003",
		Name:      "weak-crypto",
		Severity:  "medium",
		Languages: []string{"go"},
		Patterns: []skill.Pattern{
			{ImportPath: "crypto/md5"},
		},
	}

	matcher := NewMatcher([]*skill.Rule{rule}, testLogger())

	src := []byte(`package main

import (
	"crypto/md5"
	"fmt"
)

func hash() {
	h := md5.New()
	fmt.Println(h)
}`)

	gp := goparser.New()
	result, err := gp.Parse("main.go", src)
	require.NoError(t, err)

	findings := matcher.Analyze(result)

	assert.Len(t, findings, 1)
	assert.Equal(t, "TEST-003", findings[0].RuleID)
}

func TestMatcherContextStringConcat(t *testing.T) {
	rule := &skill.Rule{
		ID:        "TEST-004",
		Name:      "sql-injection",
		Severity:  "critical",
		Languages: []string{"go"},
		Patterns: []skill.Pattern{
			{
				ASTNodeType: "call_expression",
				Function:    "db.Query",
				Context:     "string_concatenation_in_args",
			},
		},
	}

	matcher := NewMatcher([]*skill.Rule{rule}, testLogger())

	src := []byte(`package main

func getUser(id string) {
	// Vulnerable (string concat)
	db.Query("SELECT * FROM users WHERE id = " + id)
	
	// Safe (parameterized)
	db.Query("SELECT * FROM users WHERE id = $1", id)
}`)

	gp := goparser.New()
	result, err := gp.Parse("main.go", src)
	require.NoError(t, err)

	findings := matcher.Analyze(result)

	for _, f := range findings {
		t.Logf("Finding: %s (line: %d)", f.CodeSnippet, f.StartLine)
	}

	assert.Len(t, findings, 1)
	assert.Equal(t, "TEST-004", findings[0].RuleID)
	assert.Equal(t, 5, findings[0].StartLine)
}

func TestMatcherMultipleLanguages(t *testing.T) {
	rule := &skill.Rule{
		ID:        "TEST-005",
		Name:      "python-only-rule",
		Languages: []string{"python"},
		Patterns: []skill.Pattern{
			{Regex: "eval\\(.*\\)"},
		},
	}

	matcher := NewMatcher([]*skill.Rule{rule}, testLogger())

	// This is Go code, the rule shouldn't trigger even though
	// the regex matches the comment.
	src := []byte(`package main
	// eval("bad")
	func main() {}`)

	gp := goparser.New()
	result, err := gp.Parse("main.go", src)
	require.NoError(t, err)

	findings := matcher.Analyze(result)

	// No findings because the file language (Go) doesn't match rule language (Python)
	assert.Len(t, findings, 0)
}
