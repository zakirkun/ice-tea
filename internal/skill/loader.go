package skill

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

// Loader discovers and loads SKILLs from the filesystem
type Loader struct {
	dirs   []string
	logger *zap.SugaredLogger
}

// NewLoader creates a new SKILL loader
func NewLoader(dirs []string, logger *zap.SugaredLogger) *Loader {
	return &Loader{dirs: dirs, logger: logger}
}

// LoadAll discovers and loads all SKILLs from configured directories
func (l *Loader) LoadAll() ([]*Skill, error) {
	var allSkills []*Skill

	for _, dir := range l.dirs {
		skills, err := l.loadFromDir(dir)
		if err != nil {
			l.logger.Warnw("Failed to load skills from directory", "dir", dir, "error", err)
			continue
		}
		allSkills = append(allSkills, skills...)
	}

	l.logger.Infow("Skills loaded", "count", len(allSkills))
	return allSkills, nil
}

// loadFromDir discovers SKILLs in a directory (recursively)
func (l *Loader) loadFromDir(dir string) ([]*Skill, error) {
	absDir, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}

	info, err := os.Stat(absDir)
	if err != nil {
		return nil, fmt.Errorf("skills directory not found: %s", absDir)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("not a directory: %s", absDir)
	}

	var skills []*Skill

	// Walk directory looking for SKILL.md files
	err = filepath.WalkDir(absDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if d.IsDir() {
			return nil
		}

		if strings.ToUpper(d.Name()) == "SKILL.MD" {
			skill, err := l.loadSkill(path)
			if err != nil {
				l.logger.Warnw("Failed to load skill", "path", path, "error", err)
				return nil
			}
			skills = append(skills, skill)
		}

		return nil
	})

	return skills, err
}

// loadSkill loads a single SKILL from its SKILL.md file
func (l *Loader) loadSkill(path string) (*Skill, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read SKILL.md: %w", err)
	}

	metadata, body, err := parseFrontmatter(content)
	if err != nil {
		return nil, fmt.Errorf("failed to parse SKILL.md frontmatter: %w", err)
	}

	skill := &Skill{
		Metadata: *metadata,
		Content:  body,
		Dir:      filepath.Dir(path),
		Loaded:   true,
	}

	// Try to load rules from patterns.yaml in the same directory
	rulesPath := filepath.Join(skill.Dir, "patterns.yaml")
	if rules, err := loadRules(rulesPath); err == nil {
		skill.Rules = rules
	}

	l.logger.Debugw("Loaded skill", "name", metadata.Name, "tags", metadata.Tags)
	return skill, nil
}

// parseFrontmatter parses YAML frontmatter from SKILL.md content
func parseFrontmatter(content []byte) (*SkillMetadata, string, error) {
	reader := bufio.NewReader(bytes.NewReader(content))

	// Check for frontmatter delimiter
	firstLine, err := reader.ReadString('\n')
	if err != nil || strings.TrimSpace(firstLine) != "---" {
		// No frontmatter — treat entire content as body
		return &SkillMetadata{}, string(content), nil
	}

	// Read until closing delimiter
	var frontmatterBuf bytes.Buffer
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, "", fmt.Errorf("unterminated frontmatter")
		}
		if strings.TrimSpace(line) == "---" {
			break
		}
		frontmatterBuf.WriteString(line)
	}

	// Parse YAML frontmatter
	var metadata SkillMetadata
	if err := yaml.Unmarshal(frontmatterBuf.Bytes(), &metadata); err != nil {
		return nil, "", fmt.Errorf("invalid YAML frontmatter: %w", err)
	}

	// Remaining content is the body
	var bodyBuf bytes.Buffer
	for {
		line, err := reader.ReadString('\n')
		bodyBuf.WriteString(line)
		if err != nil {
			break
		}
	}

	return &metadata, bodyBuf.String(), nil
}

// loadRules loads detection rules from a patterns.yaml file
func loadRules(path string) ([]Rule, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var wrapper struct {
		Rules []Rule `yaml:"rules"`
	}
	if err := yaml.Unmarshal(data, &wrapper); err != nil {
		return nil, fmt.Errorf("invalid rules YAML: %w", err)
	}

	return wrapper.Rules, nil
}

// BuildIndex creates a RuleIndex from a set of loaded skills
func BuildIndex(skills []*Skill) *RuleIndex {
	index := NewRuleIndex()

	for _, skill := range skills {
		for i := range skill.Rules {
			rule := &skill.Rules[i]

			// Inherit metadata from skill if rule doesn't specify
			if rule.Severity == "" {
				rule.Severity = skill.Metadata.Severity
			}
			if rule.Confidence == "" {
				rule.Confidence = skill.Metadata.Confidence
			}
			if len(rule.Languages) == 0 {
				rule.Languages = skill.Metadata.Languages
			}
			if len(rule.CWE) == 0 {
				rule.CWE = skill.Metadata.CWE
			}
			if len(rule.OWASP) == 0 {
				rule.OWASP = skill.Metadata.OWASP
			}

			index.AddRule(rule)
		}
	}

	return index
}
