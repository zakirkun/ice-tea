package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type SkillDef struct {
	Dir         string
	Name        string
	Description string
	Tags        []string
	CWE         []string
	OWASP       []string
	Severity    string
	Remediation string
	Rules       []RuleDef
}

type RuleDef struct {
	ID          string
	Name        string
	Description string
	Language    string
	Severity    string
	Patterns    []string // YAML lines for patterns
}

func main() {
	baseDir := "skills"
	os.MkdirAll(baseDir, 0755)

	skills := []SkillDef{
		{
			Dir:         "web/sql-injection",
			Name:        "SQL Injection",
			Description: "Detects untrusted input concatenated directly into SQL queries.",
			Tags:        []string{"sqli", "web", "injection", "database"},
			CWE:         []string{"CWE-89"},
			OWASP:       []string{"A03:2021", "A03:2025"},
			Severity:    "critical",
			Remediation: "Use parameterized queries or prepared statements instead of string concatenation.",
			Rules: []RuleDef{
				{ID: "SQLI-GO-01", Name: "go-sqli-concat", Language: "go", Severity: "critical", Description: "SQL query using string concatenation", Patterns: []string{
					"- ast_node_type: call_expression",
					"  function: db.Query",
					"  context: string_concatenation_in_args",
					"- regex: '(?i)(SELECT|UPDATE|INSERT|DELETE).*?\\+.*?'",
				}},
				{ID: "SQLI-PY-01", Name: "python-sqli-format", Language: "python", Severity: "critical", Description: "SQL query using f-strings or format", Patterns: []string{
					"- regex: '(?i)execute\\s*\\(\\s*f?[\"\\'].*(SELECT|UPDATE|INSERT|DELETE).*?\\{.*\\}.*[\"\\']\\s*\\)'",
					"- regex: '(?i)execute\\s*\\(\\s*[\"\\'].*(SELECT|UPDATE|INSERT|DELETE).*?[\"\\']\\s*%.*\\)'",
				}},
				{ID: "SQLI-JS-01", Name: "js-sqli-template", Language: "javascript", Severity: "critical", Description: "SQL query using template literals", Patterns: []string{
					"- regex: '(?i)query\\s*\\(\\s*`.*(SELECT|UPDATE|INSERT|DELETE).*?\\$\\{.*\\}.*`\\s*\\)'",
				}},
				{ID: "SQLI-TS-01", Name: "ts-sqli-template", Language: "typescript", Severity: "critical", Description: "SQL query using template literals in TS", Patterns: []string{
					"- regex: '(?i)query\\s*\\(\\s*`.*(SELECT|UPDATE|INSERT|DELETE).*?\\$\\{.*\\}.*`\\s*\\)'",
				}},
				{ID: "SQLI-JAVA-01", Name: "java-sqli-concat", Language: "java", Severity: "critical", Description: "SQL query string concatenation", Patterns: []string{
					"- regex: '(?i)executeQuery\\s*\\(.*(SELECT|UPDATE|INSERT|DELETE).*?\\+.*\\)'",
					"- regex: '(?i)createStatement\\s*\\(\\s*\\)\\.execute\\s*\\(.*(SELECT|UPDATE|INSERT|DELETE).*?\\+.*\\)'",
				}},
				{ID: "SQLI-PHP-01", Name: "php-sqli-concat", Language: "php", Severity: "critical", Description: "SQL query using string concatenation", Patterns: []string{
					"- regex: '(?i)mysql_query\\s*\\(.*(SELECT|UPDATE|INSERT|DELETE).*?\\..*\\)'",
					"- regex: '(?i)->query\\s*\\(.*(SELECT|UPDATE|INSERT|DELETE).*?\\..*\\)'",
				}},
				{ID: "SQLI-RUBY-01", Name: "ruby-sqli-interp", Language: "ruby", Severity: "critical", Description: "SQL query using string interpolation", Patterns: []string{
					"- regex: '(?i)(?:execute|query)\\s*[\"\\'].*(SELECT|UPDATE|INSERT|DELETE).*?\\#\\{.*\\}.*[\"\\']'",
				}},
			},
		},
		{
			Dir:         "web/command-injection",
			Name:        "Command Injection",
			Description: "Detects untrusted input passed directly to operating system shell commands.",
			Tags:        []string{"cmdi", "rce", "injection", "os"},
			CWE:         []string{"CWE-78"},
			OWASP:       []string{"A03:2025"},
			Severity:    "critical",
			Remediation: "Avoid calling OS commands directly. Use built-in language APIs. If necessary, use exec arrays instead of shell strings.",
			Rules: []RuleDef{
				{ID: "CMD-GO-01", Name: "go-cmdi", Language: "go", Severity: "critical", Description: "exec.Command with untrusted input", Patterns: []string{
					"- ast_node_type: call_expression",
					"  function: exec.Command",
				}},
				{ID: "CMD-PY-01", Name: "python-cmdi-os-system", Language: "python", Severity: "critical", Description: "os.system with untrusted input", Patterns: []string{
					"- ast_node_type: call",
					"  function: os.system",
					"- ast_node_type: call",
					"  function: subprocess.Popen",
					"- regex: 'subprocess\\.run\\(.*shell\\s*=\\s*True.*\\)'",
				}},
				{ID: "CMD-JS-01", Name: "js-cmdi-exec", Language: "javascript", Severity: "critical", Description: "child_process.exec with untrusted input", Patterns: []string{
					"- ast_node_type: call_expression",
					"  function: exec",
					"- regex: 'child_process\\.exec\\(.*\\+.*\\)'",
				}},
				{ID: "CMD-JAVA-01", Name: "java-cmdi-runtime", Language: "java", Severity: "critical", Description: "Runtime.exec with untrusted input", Patterns: []string{
					"- regex: 'Runtime\\.getRuntime\\(\\)\\.exec\\(.*\\+.*\\)'",
				}},
				{ID: "CMD-PHP-01", Name: "php-cmdi-system", Language: "php", Severity: "critical", Description: "system/exec/passthru with untrusted input", Patterns: []string{
					"- regex: '(?i)(system|exec|passthru|shell_exec)\\s*\\(.*\\$.*\\)'",
				}},
				{ID: "CMD-RUBY-01", Name: "ruby-cmdi-system", Language: "ruby", Severity: "critical", Description: "system or backticks with untrusted input", Patterns: []string{
					"- regex: 'system\\s*\\(.*\\#\\{.*\\}\\)'",
					"- regex: '`.*\\#\\{.*\\}.*`'",
				}},
				{ID: "CMD-C-01", Name: "c-cmdi-system", Language: "c", Severity: "critical", Description: "system() in C", Patterns: []string{
					"- ast_node_type: call_expression",
					"  function: system",
					"- regex: '\\bsystem\\s*\\(.*\\)'",
				}},
				{ID: "CMD-CPP-01", Name: "cpp-cmdi-system", Language: "cpp", Severity: "critical", Description: "system() in C++", Patterns: []string{
					"- ast_node_type: call_expression",
					"  function: system",
					"- regex: '\\bsystem\\s*\\(.*\\)'",
				}},
				{ID: "CMD-RUST-01", Name: "rust-cmdi-command", Language: "rust", Severity: "critical", Description: "Command::new in Rust", Patterns: []string{
					"- regex: 'Command::new\\s*\\(.*\\)'",
				}},
			},
		},
		{
			Dir:         "crypto/weak-hashing",
			Name:        "Weak Hashing Algorithms",
			Description: "Detects the use of cryptographically weak hashing algorithms like MD5 and SHA1.",
			Tags:        []string{"crypto", "hash", "md5", "sha1", "owasp-a02"},
			CWE:         []string{"CWE-327", "CWE-328"},
			OWASP:       []string{"A02:2025"},
			Severity:    "high",
			Remediation: "Use strong hashing algorithms like SHA-256, SHA-3, or argon2/bcrypt for passwords.",
			Rules: []RuleDef{
				{ID: "HASH-GO-01", Name: "go-md5-sha1", Language: "go", Severity: "high", Description: "MD5 or SHA1 in Go", Patterns: []string{"- regex: 'crypto/(md5|sha1)'"}},
				{ID: "HASH-PY-01", Name: "py-md5-sha1", Language: "python", Severity: "high", Description: "MD5 or SHA1 in Python", Patterns: []string{"- regex: 'hashlib\\.(md5|sha1)'"}},
				{ID: "HASH-JS-01", Name: "js-md5-sha1", Language: "javascript", Severity: "high", Description: "MD5 or SHA1 in JS", Patterns: []string{"- regex: 'createHash\\([\\\"\\'](md5|sha1)[\\\"\\']\\)'", "- regex: 'crypto-js/(md5|sha1)'"}},
				{ID: "HASH-TS-01", Name: "ts-md5-sha1", Language: "typescript", Severity: "high", Description: "MD5 or SHA1 in TS", Patterns: []string{"- regex: 'createHash\\([\\\"\\'](md5|sha1)[\\\"\\']\\)'"}},
				{ID: "HASH-JAVA-01", Name: "java-md5-sha1", Language: "java", Severity: "high", Description: "MD5 or SHA1 in Java", Patterns: []string{"- regex: 'MessageDigest\\.getInstance\\([\\\"\\'](MD5|SHA-1)[\\\"\\']\\)'"}},
				{ID: "HASH-PHP-01", Name: "php-md5-sha1", Language: "php", Severity: "high", Description: "MD5 or SHA1 in PHP", Patterns: []string{"- regex: '\\b(md5|sha1)\\s*\\('"}},
				{ID: "HASH-RUBY-01", Name: "ruby-md5-sha1", Language: "ruby", Severity: "high", Description: "MD5 or SHA1 in Ruby", Patterns: []string{"- regex: 'Digest::(MD5|SHA1)'"}},
				{ID: "HASH-C-01", Name: "c-md5-sha1", Language: "c", Severity: "high", Description: "MD5 or SHA1 in C (OpenSSL)", Patterns: []string{"- regex: '\\b(MD5|SHA1)\\s*\\('"}},
				{ID: "HASH-CPP-01", Name: "cpp-md5-sha1", Language: "cpp", Severity: "high", Description: "MD5 or SHA1 in C++", Patterns: []string{"- regex: '\\b(MD5|SHA1)\\s*\\('"}},
				{ID: "HASH-RUST-01", Name: "rust-md5-sha1", Language: "rust", Severity: "high", Description: "MD5 or SHA1 in Rust", Patterns: []string{"- regex: 'md5::compute'", "- regex: 'sha1::Sha1'"}},
			},
		},
		{
			Dir:         "auth/hardcoded-secrets",
			Name:        "Hardcoded Secrets",
			Description: "Detects API keys, passwords, and tokens embedded directly in the source code.",
			Tags:        []string{"secrets", "auth", "hardcoded", "owasp-a07"},
			CWE:         []string{"CWE-798"},
			OWASP:       []string{"A07:2025"},
			Severity:    "critical",
			Remediation: "Use a secure secrets manager or environment variables to inject sensitive credentials at runtime.",
			Rules: []RuleDef{
				{ID: "SEC-GEN-01", Name: "generic-aws-key", Language: "generic", Severity: "critical", Description: "AWS Access Key", Patterns: []string{"- regex: '(?i)(A3T[A-Z0-9]|AKIA|AGPA|AIDA|AROA|AIPA|ANPA|ANVA|ASIA)[A-Z0-9]{16}'"}},
				{ID: "SEC-GEN-02", Name: "generic-google-key", Language: "generic", Severity: "critical", Description: "Google API Key", Patterns: []string{"- regex: 'AIza[0-9A-Za-z\\-_]{35}'"}},
				{ID: "SEC-GEN-03", Name: "generic-stripe-key", Language: "generic", Severity: "critical", Description: "Stripe Key", Patterns: []string{"- regex: '(?i)sk_live_[0-9a-zA-Z]{24}'"}},
				{ID: "SEC-GEN-04", Name: "generic-slack-token", Language: "generic", Severity: "critical", Description: "Slack Token", Patterns: []string{"- regex: 'xox[baprs]-[0-9a-zA-Z]{10,48}'"}},
				{ID: "SEC-GEN-05", Name: "generic-github-token", Language: "generic", Severity: "critical", Description: "GitHub Token", Patterns: []string{"- regex: '(?i)gh[pousr]_[0-9a-zA-Z]{36}'"}},
				{ID: "SEC-GEN-06", Name: "generic-password", Language: "generic", Severity: "high", Description: "Variables named password with hardcoded string", Patterns: []string{"- regex: '(?i)(password|passwd|pwd)\\s*[:=]\\s*[\"\\'][^\"\\']+[\"\\']'"}},
				{ID: "SEC-GEN-07", Name: "generic-secret", Language: "generic", Severity: "high", Description: "Variables named secret with hardcoded string", Patterns: []string{"- regex: '(?i)(secret|client_secret|api_secret)\\s*[:=]\\s*[\"\\'][^\"\\']+[\"\\']'"}},
				{ID: "SEC-GEN-08", Name: "generic-bearer", Language: "generic", Severity: "critical", Description: "Hardcoded Bearer Token", Patterns: []string{"- regex: '(?i)bearer\\s+[a-zA-Z0-9_\\-\\.]{30,}'"}},
				{ID: "SEC-GEN-09", Name: "generic-rsa-private", Language: "generic", Severity: "critical", Description: "RSA Private Key", Patterns: []string{"- regex: '-----BEGIN RSA PRIVATE KEY-----'"}},
				{ID: "SEC-GEN-10", Name: "generic-pgp-private", Language: "generic", Severity: "critical", Description: "PGP Private Key", Patterns: []string{"- regex: '-----BEGIN PGP PRIVATE KEY BLOCK-----'"}},
			},
		},
		{
			Dir:         "web/ssrf",
			Name:        "Server-Side Request Forgery (SSRF)",
			Description: "Detects HTTP requests made with user-controlled URLs, potentially allowing internal network access.",
			Tags:        []string{"ssrf", "web", "network", "owasp-a10"},
			CWE:         []string{"CWE-918"},
			OWASP:       []string{"A10:2025"},
			Severity:    "high",
			Remediation: "Validate URLs against a strict allowlist. Do not permit access to loopback or private IP ranges.",
			Rules: []RuleDef{
				{ID: "SSRF-GO-01", Name: "go-http-get", Language: "go", Severity: "high", Description: "http.Get with variable", Patterns: []string{"- regex: 'http\\.(Get|Post|Do)\\([a-zA-Z_]'", "- ast_node_type: call_expression\n  function: http.Get"}},
				{ID: "SSRF-PY-01", Name: "py-requests-get", Language: "python", Severity: "high", Description: "requests.get with variable", Patterns: []string{"- ast_node_type: call\n  function: requests.get", "- regex: 'requests\\.(get|post|request)\\([a-zA-Z_]'"}},
				{ID: "SSRF-JS-01", Name: "js-fetch", Language: "javascript", Severity: "high", Description: "fetch with variable", Patterns: []string{"- ast_node_type: call_expression\n  function: fetch", "- regex: 'axios\\.(get|post)\\([a-zA-Z_]'"}},
				{ID: "SSRF-JAVA-01", Name: "java-url-conn", Language: "java", Severity: "high", Description: "URL openConnection with variable", Patterns: []string{"- regex: 'new\\s+URL\\([a-zA-Z_][a-zA-Z0-9_]*\\)\\.openConnection'"}},
				{ID: "SSRF-PHP-01", Name: "php-curl-file-get", Language: "php", Severity: "high", Description: "file_get_contents or curl with variable", Patterns: []string{"- regex: '(file_get_contents|curl_init)\\(\\s*\\$'"}},
				{ID: "SSRF-RUBY-01", Name: "ruby-net-http", Language: "ruby", Severity: "high", Description: "Net::HTTP.get with variable", Patterns: []string{"- regex: 'Net::HTTP\\.(get|post_form)\\([a-zA-Z_]'"}},
				{ID: "SSRF-C-01", Name: "c-libcurl", Language: "c", Severity: "high", Description: "curl_easy_setopt with URL", Patterns: []string{"- regex: 'curl_easy_setopt\\(.*CURLOPT_URL.*\\)'"}},
			},
		},
		{
			Dir:         "crypto/weak-random",
			Name:        "Weak Random Number Generation",
			Description: "Detects the use of PRNGs (Pseudo-Random Number Generators) that are not cryptographically secure.",
			Tags:        []string{"crypto", "random", "prng", "owasp-a02"},
			CWE:         []string{"CWE-338"},
			OWASP:       []string{"A02:2025"},
			Severity:    "medium",
			Remediation: "Use cryptographically secure PRNGs (CSPRNG), like crypto/rand in Go, secrets in Python, or Crypto.getRandomValues in JS.",
			Rules: []RuleDef{
				{ID: "RAND-GO-01", Name: "go-math-rand", Language: "go", Severity: "medium", Description: "Use of math/rand", Patterns: []string{"- regex: 'math/rand'"}},
				{ID: "RAND-PY-01", Name: "py-random", Language: "python", Severity: "medium", Description: "Use of random module", Patterns: []string{"- ast_node_type: call", "  function: random.random", "- regex: 'random\\.(randint|choice|random|randrange)'"}},
				{ID: "RAND-JS-01", Name: "js-math-random", Language: "javascript", Severity: "medium", Description: "Use of Math.random", Patterns: []string{"- regex: 'Math\\.random\\(\\)'"}},
				{ID: "RAND-JAVA-01", Name: "java-util-random", Language: "java", Severity: "medium", Description: "Use of java.util.Random", Patterns: []string{"- regex: 'new\\s+java\\.util\\.Random\\(\\)'", "- regex: 'new\\s+Random\\(\\)'"}},
				{ID: "RAND-PHP-01", Name: "php-rand-mt-rand", Language: "php", Severity: "medium", Description: "Use of rand or mt_rand", Patterns: []string{"- regex: '\\b(rand|mt_rand|uniqid)\\s*\\('"}},
				{ID: "RAND-RUBY-01", Name: "ruby-rand", Language: "ruby", Severity: "medium", Description: "Use of Kernel.rand", Patterns: []string{"- regex: '\\brand\\('"}},
				{ID: "RAND-C-01", Name: "c-rand", Language: "c", Severity: "medium", Description: "Use of rand()", Patterns: []string{"- regex: '\\brand\\(\\)'"}},
				{ID: "RAND-CPP-01", Name: "cpp-rand", Language: "cpp", Severity: "medium", Description: "Use of rand()", Patterns: []string{"- regex: '\\brand\\(\\)'"}},
				{ID: "RAND-RUST-01", Name: "rust-rand-thread_rng", Language: "rust", Severity: "medium", Description: "Use of thread_rng without crypto feature explicitly (heuristic)", Patterns: []string{"- regex: 'rand::thread_rng()'"}},
			},
		},
		{
			Dir:         "fs/path-traversal-extended",
			Name:        "Path Traversal (Extended)",
			Description: "Detects unsafe file writing or dynamic inclusions leading to LFI/path traversal.",
			Tags:        []string{"fs", "traversal", "lfi", "owasp-a01"},
			CWE:         []string{"CWE-22", "CWE-98"},
			OWASP:       []string{"A01:2025"},
			Severity:    "high",
			Remediation: "Sanitize paths and restrict inclusions to allowed directories.",
			Rules: []RuleDef{
				{ID: "LFI-PHP-01", Name: "php-include-require", Language: "php", Severity: "critical", Description: "Dynamic include/require", Patterns: []string{"- regex: '(?i)(include|require|include_once|require_once)\\s*\\(\\s*\\$.*\\)'"}},
				{ID: "LFI-JAVA-01", Name: "java-file-read", Language: "java", Severity: "high", Description: "Unsafe file read", Patterns: []string{"- regex: 'new\\s+File\\(.*\\+.*\\)'"}},
				{ID: "LFI-C-01", Name: "c-fopen", Language: "c", Severity: "high", Description: "Unsafe fopen", Patterns: []string{"- regex: 'fopen\\([a-zA-Z_]'"}},
				{ID: "LFI-JS-01", Name: "js-fs-readfile", Language: "javascript", Severity: "high", Description: "fs.readFile with variable", Patterns: []string{"- regex: 'fs\\.(readFile|readFileSync)\\([a-zA-Z_]'"}},
				{ID: "LFI-PY-01", Name: "py-open", Language: "python", Severity: "high", Description: "open with variable", Patterns: []string{"- regex: 'open\\([a-zA-Z_]'"}},
			},
		},
		{
			Dir:         "web/deserialization",
			Name:        "Insecure Deserialization",
			Description: "Detects deserialization of untrusted data which can lead to Remote Code Execution.",
			Tags:        []string{"deserialization", "rce", "owasp-a08"},
			CWE:         []string{"CWE-502"},
			OWASP:       []string{"A08:2025"},
			Severity:    "critical",
			Remediation: "Use safe data formats like JSON instead of native serialization. If native serialization is required, sign the objects.",
			Rules: []RuleDef{
				{ID: "DESER-PY-01", Name: "py-pickle", Language: "python", Severity: "critical", Description: "Pickle loads", Patterns: []string{"- regex: 'pickle\\.loads\\('"}},
				{ID: "DESER-PY-02", Name: "py-yaml-load", Language: "python", Severity: "critical", Description: "Unsafe yaml.load", Patterns: []string{"- regex: 'yaml\\.load\\([^,]*\\)'"}}, // load without Loader=SafeLoader
				{ID: "DESER-JAVA-01", Name: "java-object-input-stream", Language: "java", Severity: "critical", Description: "ObjectInputStream", Patterns: []string{"- regex: 'new\\s+ObjectInputStream\\('"}},
				{ID: "DESER-PHP-01", Name: "php-unserialize", Language: "php", Severity: "critical", Description: "unserialize()", Patterns: []string{"- regex: 'unserialize\\('"}},
				{ID: "DESER-RUBY-01", Name: "ruby-marshal-load", Language: "ruby", Severity: "critical", Description: "Marshal.load", Patterns: []string{"- regex: 'Marshal\\.load\\('"}},
				{ID: "DESER-JS-01", Name: "js-node-serialize", Language: "javascript", Severity: "critical", Description: "node-serialize", Patterns: []string{"- regex: 'unserialize\\('"}},
			},
		},
	}

	totalRules := 0
	for _, skill := range skills {
		skillDir := filepath.Join(baseDir, skill.Dir)
		_ = os.MkdirAll(skillDir, 0755)

		mdPath := filepath.Join(skillDir, "SKILL.md")
		writeMD(mdPath, skill)

		yamlPath := filepath.Join(skillDir, "patterns.yaml")
		writeYAML(yamlPath, skill.Rules)

		totalRules += len(skill.Rules)
	}

	fmt.Printf("Successfully generated %d SKILL directories containing a total of %d detailed patterns.\n", len(skills), totalRules)
}

func writeMD(path string, s SkillDef) {
	content := fmt.Sprintf(`---
name: %s
version: 1.0.0
description: %s
tags: [%s]
languages: [generic]
severity: %s
confidence: medium
cwe: [%s]
owasp: [%s]
---

# %s

## Overview
%s

## Remediation
%s
`, s.Name, s.Description, strings.Join(s.Tags, ", "), s.Severity, strings.Join(s.CWE, ", "), strings.Join(s.OWASP, ", "), s.Name, s.Description, s.Remediation)
	os.WriteFile(path, []byte(content), 0644)
}

func writeYAML(path string, rules []RuleDef) {
	var sb strings.Builder
	sb.WriteString("rules:\n")

	for _, r := range rules {
		sb.WriteString("  - id: " + r.ID + "\n")
		sb.WriteString("    name: " + r.Name + "\n")
		sb.WriteString(fmt.Sprintf("    description: %q\n", r.Description))
		sb.WriteString("    severity: " + r.Severity + "\n")
		sb.WriteString("    confidence: medium\n")
		sb.WriteString(fmt.Sprintf("    languages: [%s]\n", r.Language))
		sb.WriteString("    patterns:\n")
		for _, p := range r.Patterns {
			lines := strings.Split(p, "\n")
			for _, l := range lines {
				sb.WriteString("      " + l + "\n")
			}
		}
		sb.WriteString("\n")
	}

	os.WriteFile(path, []byte(sb.String()), 0644)
}
