package pattern

import (
	"fmt"
	"regexp"
	"strings"

	"go.uber.org/zap"

	"github.com/zakirkun/ice-tea/internal/finding"
	"github.com/zakirkun/ice-tea/internal/parser"
	"github.com/zakirkun/ice-tea/internal/skill"
)

// Matcher performs static pattern matching against AST nodes
type Matcher struct {
	rules  []*skill.Rule
	logger *zap.SugaredLogger
}

// NewMatcher creates a new pattern matcher with the given rules
func NewMatcher(rules []*skill.Rule, logger *zap.SugaredLogger) *Matcher {
	return &Matcher{
		rules:  rules,
		logger: logger,
	}
}

// Analyze scans a parsed file against all loaded rules
func (m *Matcher) Analyze(result *parser.ParseResult) []*finding.Finding {
	if result == nil || result.Root == nil {
		return nil
	}

	var findings []*finding.Finding
	lang := string(result.Language)

	for _, rule := range m.rules {
		// Skip rules not applicable to this language
		if !ruleMatchesLanguage(rule, lang) {
			continue
		}

		for _, pattern := range rule.Patterns {
			matches := m.matchPattern(result, &pattern)
			for _, match := range matches {
				f := &finding.Finding{
					ID:          fmt.Sprintf("%s-%s-%d", rule.ID, result.FilePath, match.Line),
					RuleID:      rule.ID,
					Type:        rule.Name,
					Severity:    rule.Severity,
					Confidence:  rule.Confidence,
					CWE:         rule.CWE,
					OWASP:       rule.OWASP,
					File:        result.FilePath,
					StartLine:   match.Line,
					EndLine:     match.EndLine,
					StartColumn: match.Column,
					CodeSnippet: match.Text,
					Message:     rule.Description,
					Engines:     []string{finding.EnginePattern},
				}
				findings = append(findings, f)
			}
		}
	}

	return findings
}

// matchResult represents a pattern match location
type matchResult struct {
	Line    int
	EndLine int
	Column  int
	Text    string
}

// matchPattern applies a single pattern against the AST
func (m *Matcher) matchPattern(result *parser.ParseResult, pattern *skill.Pattern) []matchResult {
	var matches []matchResult

	// AST node type matching (or shorthand function matching)
	if pattern.ASTNodeType != "" || pattern.Function != "" {
		nodeType := pattern.ASTNodeType
		if nodeType == "" && pattern.Function != "" {
			nodeType = "call_expression"
		}

		nodes := parser.FindAll(result.Root, nodeType)
		for _, node := range nodes {
			if m.nodeMatchesPattern(node, pattern) {
				matches = append(matches, matchResult{
					Line:    node.Start.Line,
					EndLine: node.End.Line,
					Column:  node.Start.Column,
					Text:    node.Text,
				})
			}
		}
	}

	// Import path matching
	if pattern.ImportPath != "" {
		imports := parser.FindAll(result.Root, "import")
		for _, imp := range imports {
			importPath := strings.Trim(imp.Fields["path"], "\"")
			if importPath == pattern.ImportPath {
				matches = append(matches, matchResult{
					Line:    imp.Start.Line,
					EndLine: imp.End.Line,
					Text:    imp.Text,
				})
			}
		}
	}

	// Regex pattern matching on source text
	if pattern.Regex != "" {
		regexMatches := m.matchRegex(result.Source, pattern.Regex)
		matches = append(matches, regexMatches...)
	}

	return matches
}

// nodeMatchesPattern checks if an AST node matches all pattern criteria
func (m *Matcher) nodeMatchesPattern(node *parser.Node, pattern *skill.Pattern) bool {
	// Check function name if specified
	if pattern.Function != "" {
		if node.Fields["function"] != pattern.Function {
			return false
		}
	}

	// Check object if specified
	if pattern.Object != "" {
		if node.Fields["object"] != pattern.Object {
			return false
		}
	}

	// Check method if specified
	if pattern.Method != "" {
		if node.Fields["method"] != pattern.Method {
			return false
		}
	}

	// Check context (e.g., look for string concatenation in arguments)
	if pattern.Context != "" {
		return m.checkContext(node, pattern.Context)
	}

	return true
}

// checkContext evaluates contextual conditions
func (m *Matcher) checkContext(node *parser.Node, context string) bool {
	switch context {
	case "string_concatenation_in_args":
		// Check if any child is a binary expression with "+" operator
		for _, child := range node.Children {
			if child.Type == "binary_expression" && child.Fields["operator"] == "+" {
				return true
			}
		}
	case "sql_query_argument":
		// Check if this call is within a SQL query context
		text := strings.ToUpper(node.Text)
		sqlKeywords := []string{"SELECT", "INSERT", "UPDATE", "DELETE", "WHERE", "FROM", "JOIN"}
		for _, kw := range sqlKeywords {
			if strings.Contains(text, kw) {
				return true
			}
		}
	}
	return false
}

// matchRegex applies a regex pattern against source code
func (m *Matcher) matchRegex(src []byte, pattern string) []matchResult {
	re, err := regexp.Compile(pattern)
	if err != nil {
		m.logger.Warnw("Invalid regex pattern", "pattern", pattern, "error", err)
		return nil
	}

	var matches []matchResult
	lines := strings.Split(string(src), "\n")

	for i, line := range lines {
		if re.MatchString(line) {
			matches = append(matches, matchResult{
				Line:    i + 1,
				EndLine: i + 1,
				Text:    strings.TrimSpace(line),
			})
		}
	}

	return matches
}

// ruleMatchesLanguage checks if a rule applies to the given language
func ruleMatchesLanguage(rule *skill.Rule, lang string) bool {
	if len(rule.Languages) == 0 {
		return true // no language restriction
	}
	for _, l := range rule.Languages {
		if strings.EqualFold(l, lang) {
			return true
		}
	}
	return false
}
