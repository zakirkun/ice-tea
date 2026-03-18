package skill

import "time"

// Skill represents a loaded vulnerability detection skill
type Skill struct {
	Metadata    SkillMetadata `yaml:"-"`
	Content     string        `yaml:"-"` // full SKILL.md content (loaded on demand)
	Rules       []Rule        `yaml:"-"` // parsed detection rules
	Dir         string        `yaml:"-"` // directory containing this skill
	Loaded      bool          `yaml:"-"` // whether full content has been loaded
}

// SkillMetadata represents SKILL.md frontmatter
type SkillMetadata struct {
	Name        string   `yaml:"name"`
	Version     string   `yaml:"version"`
	Description string   `yaml:"description"`
	Tags        []string `yaml:"tags"`
	Languages   []string `yaml:"languages"`
	Severity    string   `yaml:"severity"`
	Confidence  string   `yaml:"confidence"`
	CWE         []string `yaml:"cwe"`
	OWASP       []string `yaml:"owasp"`
}

// Rule defines a vulnerability detection rule
type Rule struct {
	ID          string    `yaml:"id"`
	Name        string    `yaml:"name"`
	Description string    `yaml:"description"`
	Severity    string    `yaml:"severity"`
	Confidence  string    `yaml:"confidence"`
	CWE         []string  `yaml:"cwe"`
	OWASP       []string  `yaml:"owasp"`
	Languages   []string  `yaml:"languages"`
	Patterns    []Pattern `yaml:"patterns"`
}

// Pattern defines a detection pattern within a rule
type Pattern struct {
	// AST-based patterns
	ASTNodeType string `yaml:"ast_node_type"` // e.g., "call_expression"
	Function    string `yaml:"function"`      // e.g., "db.Query", "exec.Command"
	Object      string `yaml:"object"`        // e.g., "db", "os"
	Method      string `yaml:"method"`        // e.g., "Query", "Command"

	// Import-based patterns
	ImportPath string `yaml:"import_path"` // e.g., "crypto/md5"

	// Text/regex-based patterns
	Regex string `yaml:"regex"` // regex pattern to match in source text

	// Context
	Context string `yaml:"context"` // additional context requirement
}

// RuleIndex provides fast lookup of rules by various criteria
type RuleIndex struct {
	ByID       map[string]*Rule
	ByLanguage map[string][]*Rule // language → rules
	ByTag      map[string][]*Rule // tag → rules
	ByCWE      map[string][]*Rule // CWE ID → rules
	All        []*Rule
	UpdatedAt  time.Time
}

// NewRuleIndex creates a new empty rule index
func NewRuleIndex() *RuleIndex {
	return &RuleIndex{
		ByID:       make(map[string]*Rule),
		ByLanguage: make(map[string][]*Rule),
		ByTag:      make(map[string][]*Rule),
		ByCWE:      make(map[string][]*Rule),
		UpdatedAt:  time.Now(),
	}
}

// AddRule indexes a rule for fast lookup
func (idx *RuleIndex) AddRule(rule *Rule) {
	idx.ByID[rule.ID] = rule
	idx.All = append(idx.All, rule)

	for _, lang := range rule.Languages {
		idx.ByLanguage[lang] = append(idx.ByLanguage[lang], rule)
	}

	for _, cwe := range rule.CWE {
		idx.ByCWE[cwe] = append(idx.ByCWE[cwe], rule)
	}
}

// GetRulesForLanguage returns all rules applicable to a given language
func (idx *RuleIndex) GetRulesForLanguage(lang string) []*Rule {
	return idx.ByLanguage[lang]
}

// MatchesTags checks if a skill matches any of the given tags
func (m *SkillMetadata) MatchesTags(tags []string) bool {
	for _, tag := range tags {
		for _, skillTag := range m.Tags {
			if tag == skillTag {
				return true
			}
		}
	}
	return false
}

// MatchesLanguage checks if a skill supports the given language
func (m *SkillMetadata) MatchesLanguage(lang string) bool {
	if len(m.Languages) == 0 {
		return true // no language restriction
	}
	for _, l := range m.Languages {
		if l == lang {
			return true
		}
	}
	return false
}
