package skill

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func testLogger() *zap.SugaredLogger {
	logger, _ := zap.NewDevelopment()
	return logger.Sugar()
}

func TestLoaderParseFrontmatter(t *testing.T) {
	content := []byte(`---
name: SQL Injection
version: 1.0.0
description: Detects SQLi
tags: [injection, sql]
languages: [go, python]
severity: critical
confidence: high
cwe: [CWE-89]
---

# Body Header
Some body text
`)

	metadata, body, err := parseFrontmatter(content)
	require.NoError(t, err)

	// Check metadata
	assert.Equal(t, "SQL Injection", metadata.Name)
	assert.Equal(t, "1.0.0", metadata.Version)
	assert.Equal(t, []string{"injection", "sql"}, metadata.Tags)
	assert.Equal(t, []string{"go", "python"}, metadata.Languages)
	assert.Equal(t, "critical", metadata.Severity)
	assert.Equal(t, "high", metadata.Confidence)
	assert.Equal(t, []string{"CWE-89"}, metadata.CWE)

	// Check body
	assert.Contains(t, body, "# Body Header")
	assert.Contains(t, body, "Some body text")
}

func TestLoaderParseNoFrontmatter(t *testing.T) {
	content := []byte(`# Just Body
No frontmatter at all.
`)
	metadata, body, err := parseFrontmatter(content)
	require.NoError(t, err)

	assert.Empty(t, metadata.Name) // empty metadata
	assert.Contains(t, body, "# Just Body")
}

func TestLoaderLoadFromDir(t *testing.T) {
	tmpDir := t.TempDir()

	// Create a fake SKILL directory
	skillDir := filepath.Join(tmpDir, "sql-injection")
	err := os.MkdirAll(skillDir, 0755)
	require.NoError(t, err)

	// Create SKILL.md
	skillMd := `---
name: Fake SQLi
severity: high
---
Body text`
	err = os.WriteFile(filepath.Join(skillDir, "SKILL.md"), []byte(skillMd), 0644)
	require.NoError(t, err)

	// Create patterns.yaml
	patternsYaml := `rules:
  - id: FAKE-001
    name: fake-rule
    patterns:
      - import_path: "database/sql"`
	err = os.WriteFile(filepath.Join(skillDir, "patterns.yaml"), []byte(patternsYaml), 0644)
	require.NoError(t, err)

	// Load skills
	loader := NewLoader([]string{tmpDir}, testLogger())
	skills, err := loader.LoadAll()
	require.NoError(t, err)

	assert.Len(t, skills, 1)
	assert.Equal(t, "Fake SQLi", skills[0].Metadata.Name)
	assert.Equal(t, "high", skills[0].Metadata.Severity)
	assert.Contains(t, skills[0].Content, "Body text")
	
	// Check rules loaded correctly
	assert.Len(t, skills[0].Rules, 1)
	assert.Equal(t, "FAKE-001", skills[0].Rules[0].ID)
	assert.Equal(t, "database/sql", skills[0].Rules[0].Patterns[0].ImportPath)
}

func TestBuildIndex(t *testing.T) {
	skills := []*Skill{
		{
			Metadata: SkillMetadata{
				Name:      "Parent",
				Severity:  "critical",
				Languages: []string{"go", "python"},
			},
			Rules: []Rule{
				{
					ID:   "R1",
					Name: "Rule1",
					// Should inherit languages and severity
				},
				{
					ID:        "R2",
					Name:      "Rule2",
					Severity:  "low",     // Override parent
					Languages: []string{"java"}, // Override parent
				},
			},
		},
	}

	index := BuildIndex(skills)
	
	// Check full list
	assert.Len(t, index.All, 2)
	assert.NotNil(t, index.ByID["R1"])
	assert.NotNil(t, index.ByID["R2"])

	// Check inheritance
	assert.Equal(t, "critical", index.ByID["R1"].Severity)
	assert.Equal(t, []string{"go", "python"}, index.ByID["R1"].Languages)

	// Check overrides
	assert.Equal(t, "low", index.ByID["R2"].Severity)
	assert.Equal(t, []string{"java"}, index.ByID["R2"].Languages)

	// Check lookup by language
	goRules := index.GetRulesForLanguage("go")
	assert.Len(t, goRules, 1)
	assert.Equal(t, "R1", goRules[0].ID)

	javaRules := index.GetRulesForLanguage("java")
	assert.Len(t, javaRules, 1)
	assert.Equal(t, "R2", javaRules[0].ID)
}

func TestMetadataMatches(t *testing.T) {
	m := SkillMetadata{
		Tags:      []string{"sql", "injection", "database"},
		Languages: []string{"go", "java"},
	}

	// Tags
	assert.True(t, m.MatchesTags([]string{"sql"}))
	assert.True(t, m.MatchesTags([]string{"web", "database"}))
	assert.False(t, m.MatchesTags([]string{"xss", "crypto"}))

	// Languages
	assert.True(t, m.MatchesLanguage("go"))
	assert.True(t, m.MatchesLanguage("java"))
	assert.False(t, m.MatchesLanguage("python"))

	// Empty languages should match all
	mEmpty := SkillMetadata{}
	assert.True(t, mEmpty.MatchesLanguage("ruby"))
}
