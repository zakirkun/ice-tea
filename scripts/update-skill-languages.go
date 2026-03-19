//go:build ignore
// +build ignore

// Script to add kotlin, dart, zig, elixir to all skills' languages.
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

var newLanguages = []string{"kotlin", "dart", "zig", "elixir"}

func addLanguagesIfMissing(langs []string) ([]string, bool) {
	seen := make(map[string]bool)
	for _, l := range langs {
		seen[strings.ToLower(l)] = true
	}
	modified := false
	for _, nl := range newLanguages {
		if !seen[nl] {
			langs = append(langs, nl)
			seen[nl] = true
			modified = true
		}
	}
	return langs, modified
}

func hasGeneric(langs []string) bool {
	for _, l := range langs {
		if strings.EqualFold(l, "generic") {
			return true
		}
	}
	return false
}

func processSkillMD(path string) (bool, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return false, err
	}

	// Parse frontmatter
	reader := bufio.NewReader(bytes.NewReader(content))
	firstLine, err := reader.ReadString('\n')
	if err != nil || strings.TrimSpace(firstLine) != "---" {
		return false, nil // No frontmatter, skip
	}

	var frontmatterBuf bytes.Buffer
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return false, err
		}
		if strings.TrimSpace(line) == "---" {
			break
		}
		frontmatterBuf.WriteString(line)
	}

	var bodyBuf bytes.Buffer
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		bodyBuf.WriteString(line)
	}

	var meta struct {
		Languages []string `yaml:"languages"`
	}
	if err := yaml.Unmarshal(frontmatterBuf.Bytes(), &meta); err != nil {
		return false, err
	}

	// Ensure we have a languages list to add to (may be empty)
	if meta.Languages == nil {
		meta.Languages = []string{}
	}
	updated, modified := addLanguagesIfMissing(meta.Languages)
	if !modified {
		return false, nil
	}
	meta.Languages = updated

	// Re-read full frontmatter to preserve other fields, then replace languages
	var fullMeta map[string]interface{}
	if err := yaml.Unmarshal(frontmatterBuf.Bytes(), &fullMeta); err != nil {
		return false, err
	}
	// Convert []string to []interface{} for YAML
	langsIf := make([]interface{}, len(updated))
	for i, l := range updated {
		langsIf[i] = l
	}
	fullMeta["languages"] = langsIf

	newFront, err := yaml.Marshal(fullMeta)
	if err != nil {
		return false, err
	}

	var out bytes.Buffer
	out.WriteString("---\n")
	out.Write(newFront)
	out.WriteString("---\n")
	out.Write(bodyBuf.Bytes())

	return true, os.WriteFile(path, out.Bytes(), 0644)
}

func processPatternsYAML(path string) (bool, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return false, err
	}

	var doc yaml.Node
	if err := yaml.Unmarshal(data, &doc); err != nil {
		return false, err
	}

	modified := false
	var walk func(*yaml.Node)
	walk = func(n *yaml.Node) {
		if n == nil {
			return
		}
		if n.Kind == yaml.MappingNode {
			for i := 0; i < len(n.Content); i += 2 {
				if i+1 >= len(n.Content) {
					break
				}
				key, val := n.Content[i], n.Content[i+1]
				if key.Value == "rules" && val.Kind == yaml.SequenceNode {
					for _, rule := range val.Content {
						if rule.Kind == yaml.MappingNode {
							for j := 0; j < len(rule.Content); j += 2 {
								if j+1 >= len(rule.Content) {
									break
								}
								k, v := rule.Content[j], rule.Content[j+1]
								if k.Value == "languages" && v.Kind == yaml.SequenceNode {
									var langs []string
									for _, item := range v.Content {
										if item.Value != "" {
											langs = append(langs, item.Value)
										}
									}
									if hasGeneric(langs) || len(langs) < 4 {
										continue
									}
									updated, m := addLanguagesIfMissing(langs)
									if m {
										modified = true
										v.Content = nil
										for _, l := range updated {
											v.Content = append(v.Content, &yaml.Node{
												Kind:  yaml.ScalarNode,
												Value: l,
											})
										}
									}
								}
							}
						}
					}
					return
				}
			}
		}
		for _, c := range n.Content {
			walk(c)
		}
	}
	walk(&doc)

	if !modified {
		return false, nil
	}

	out, err := yaml.Marshal(&doc)
	if err != nil {
		return false, err
	}
	return true, os.WriteFile(path, out, 0644)
}

func main() {
	skillsDir := "skills"
	if len(os.Args) > 1 {
		skillsDir = os.Args[1]
	}

	skillCount := 0
	patternCount := 0

	filepath.Walk(skillsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		base := strings.ToLower(filepath.Base(path))
		if base == "skill.md" {
			mod, err := processSkillMD(path)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error %s: %v\n", path, err)
				return nil
			}
			if mod {
				skillCount++
				fmt.Println("SKILL:", path)
			}
		} else if base == "patterns.yaml" {
			mod, err := processPatternsYAML(path)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error %s: %v\n", path, err)
				return nil
			}
			if mod {
				patternCount++
				fmt.Println("patterns:", path)
			}
		}
		return nil
	})

	fmt.Printf("Updated %d SKILL.md, %d patterns.yaml\n", skillCount, patternCount)
}
