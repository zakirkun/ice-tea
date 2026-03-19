//go:build ignore
// +build ignore

// Script to remove unsupported (?!...) negative lookahead from regex patterns.
// Go's regexp uses RE2 which does not support lookahead.
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// removeNegativeLookahead removes (?!...) from pattern. Handles nested parens
// and character classes (so ) inside [^)] is not treated as closing paren).
func removeNegativeLookahead(s string) string {
	var result strings.Builder
	i := 0
	for i < len(s) {
		if i+3 <= len(s) && s[i:i+3] == `(?!` {
			depth := 1
			j := i + 3
			for j < len(s) && depth > 0 {
				if s[j] == '\\' && j+1 < len(s) {
					j += 2
					continue
				}
				if s[j] == '[' {
					// Skip character class [...]
					j++
					for j < len(s) {
						if s[j] == '\\' && j+1 < len(s) {
							j += 2
							continue
						}
						if s[j] == ']' {
							j++
							break
						}
						j++
					}
					continue
				}
				if s[j] == '(' {
					depth++
				} else if s[j] == ')' {
					depth--
				}
				j++
			}
			i = j
			continue
		}
		result.WriteByte(s[i])
		i++
	}
	return result.String()
}

func processRules(node *yaml.Node) bool {
	modified := false
	if node.Kind == yaml.MappingNode {
		for i := 0; i < len(node.Content); i += 2 {
			if i+1 >= len(node.Content) {
				break
			}
			key, val := node.Content[i], node.Content[i+1]
			if key.Value == "rules" && val.Kind == yaml.SequenceNode {
				for _, rule := range val.Content {
					modified = processRule(rule) || modified
				}
			}
		}
	}
	return modified
}

func processRule(node *yaml.Node) bool {
	modified := false
	if node.Kind != yaml.MappingNode {
		return false
	}
	for i := 0; i < len(node.Content); i += 2 {
		if i+1 >= len(node.Content) {
			break
		}
		key, val := node.Content[i], node.Content[i+1]
		if key.Value == "patterns" && val.Kind == yaml.SequenceNode {
			for _, p := range val.Content {
				if p.Kind == yaml.MappingNode {
					for j := 0; j < len(p.Content); j += 2 {
						if j+1 >= len(p.Content) {
							break
						}
						k, v := p.Content[j], p.Content[j+1]
						if k.Value == "regex" && strings.Contains(v.Value, `(?!`) {
							v.Value = removeNegativeLookahead(v.Value)
							modified = true
						}
					}
				}
			}
		}
	}
	return modified
}

func processFile(path string) (bool, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return false, err
	}
	var doc yaml.Node
	if err := yaml.Unmarshal(data, &doc); err != nil {
		return false, err
	}
	modified := false
	for _, node := range doc.Content {
		if processRules(node) {
			modified = true
		}
	}
	if modified {
		out, err := yaml.Marshal(&doc)
		if err != nil {
			return false, err
		}
		if err := os.WriteFile(path, out, 0644); err != nil {
			return false, err
		}
	}
	return modified, nil
}

func main() {
	skillsDir := "skills"
	if len(os.Args) > 1 {
		skillsDir = os.Args[1]
	}
	var count int
	filepath.Walk(skillsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".yaml" {
			return nil
		}
		mod, err := processFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error %s: %v\n", path, err)
			return nil
		}
		if mod {
			count++
			fmt.Println(path)
		}
		return nil
	})
	fmt.Printf("Fixed %d files\n", count)
}
