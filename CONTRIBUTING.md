# Contributing to Ice Tea 🍵

Thank you for your interest in contributing to Ice Tea! This document covers everything you need to know — from submitting a bug report to adding a full new vulnerability SKILL.

---

## Table of Contents

1. [Code of Conduct](#1-code-of-conduct)
2. [Ways to Contribute](#2-ways-to-contribute)
3. [Getting Started (Development Setup)](#3-getting-started-development-setup)
4. [Contributing SKILL Definitions](#4-contributing-skill-definitions)
5. [Contributing Go Code](#5-contributing-go-code)
6. [Reporting Bugs & False Positives](#6-reporting-bugs--false-positives)
7. [Pull Request Process](#7-pull-request-process)
8. [Style Guidelines](#8-style-guidelines)
9. [Skill ID Naming Convention](#9-skill-id-naming-convention)

---

## 1. Code of Conduct

Please engage with maintainers and other contributors respectfully. We expect professional, patient, and constructive communication. Harassment, personal attacks, or discriminatory language will not be tolerated.

---

## 2. Ways to Contribute

Contributions come in many forms — all are valued:

| Type | Description |
|------|-------------|
| 🛡️ **New SKILL** | Add a `SKILL.md` + `patterns.yaml` for a new vulnerability class |
| 🐛 **Bug fix** | Fix a false positive, false negative, or parser issue |
| 📖 **Documentation** | Improve docs, fix typos, add usage examples |
| 🌐 **Language support** | Add Tree-sitter parser support for a new language |
| 🤖 **LLM prompts** | Improve the LLM Engine prompt quality for a specific SKILL |
| 🧪 **Test cases** | Add vulnerable code samples to `testdata/` or `examples/` |
| 🏗️ **Core features** | New output formats, performance improvements, new CLI flags |

---

## 3. Getting Started (Development Setup)

### Prerequisites

- **Go 1.21+**: [https://go.dev/dl/](https://go.dev/dl/)
- **GCC or Clang**: Required for Tree-sitter CGO bindings
  - Linux: `sudo apt-get install build-essential`
  - macOS: `xcode-select --install`
  - Windows: Install MSYS2 + `pacman -S mingw-w64-x86_64-gcc`
- **Git**

### Fork & Clone

```bash
# 1. Fork the repository on GitHub (click "Fork" button)

# 2. Clone your fork
git clone https://github.com/YOUR_USERNAME/ice-tea.git
cd ice-tea

# 3. Add the upstream remote
git remote add upstream https://github.com/zakirkun/ice-tea.git
```

### Build & Test

```bash
# Build the binary
make build

# Verify build
./bin/ice-tea version

# Run all unit tests
go test ./...

# Run tests with race detection
go test -race ./...

# Run tests with coverage
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Create a Feature Branch

```bash
git checkout -b feat/my-new-skill
# or
git checkout -b fix/false-positive-in-sqli
```

---

## 4. Contributing SKILL Definitions

Adding SKILLs is the **most impactful** and **easiest** contribution. No Go knowledge required.

### Step 1: Check for Existing Coverage

Before starting, check if the vulnerability is already covered:

```bash
# Search existing skill names
ls skills/*/
ls skills/*/*/

# Search pattern content
grep -r "cwe: \[CWE-89\]" skills/
```

### Step 2: Choose the Right Category

Place your SKILL in the most appropriate existing category, or create a new top-level directory if none fits:

| Directory | Used For |
|-----------|----------|
| `skills/auth/` | Authentication & authorization flaws |
| `skills/injection/` | All injection vulnerability classes |
| `skills/web/` | Browser/HTTP-layer web vulnerabilities |
| `skills/crypto/` | Cryptographic misuse |
| `skills/fs/` | Filesystem and file handling vulnerabilities |
| `skills/api/` | REST/GraphQL API-specific issues |
| `skills/infra/` | Infrastructure & configuration security |
| `skills/logging/` | Logging and monitoring failures |
| `skills/memory/` | Memory safety (C/C++/Rust) |
| `skills/cloud/` | Cloud platform misconfigurations |
| `skills/android/` | Android mobile security |
| `skills/network/` | Network-level vulnerabilities |

### Step 3: Create the SKILL Directory

```bash
mkdir -p skills/<category>/<vuln-name>
```

Example: `mkdir -p skills/web/clickjacking`

### Step 4: Write `SKILL.md`

```markdown
---
name: <Human-readable name>
version: 1.0.0
description: <One sentence: what this skill detects>
tags: [<tag1>, <tag2>, owasp-aXX]
languages: [go, python, javascript, typescript, java, php, ruby]
severity: critical | high | medium | low | info
confidence: high | medium | low
cwe: [CWE-XXX]
owasp: [AXX:2025]
---

# <Vulnerability Name>

## Overview
<2-4 sentences explaining the vulnerability, why it's dangerous, and what attackers can do with it.>

## Detection Strategy
<Describe what code patterns to look for. This is the context the LLM Engine uses.>

Key patterns:
- <Pattern 1>
- <Pattern 2>

## Remediation
<Clear, actionable advice. Include before/after code examples in at least one language.>

**Vulnerable:**
```lang
// example vulnerable code
```

**Safe:**
```lang
// example safe code
```
```

**Frontmatter field reference:**

| Field | Required | Values |
|-------|----------|--------|
| `name` | ✅ | Human-readable string |
| `version` | ✅ | Semantic version (start with `1.0.0`) |
| `description` | ✅ | One sentence, ends with period |
| `tags` | ✅ | Lowercase kebab-case array |
| `languages` | ✅ | Array of supported languages, or `[generic]` for any |
| `severity` | ✅ | `critical`, `high`, `medium`, `low`, `info` |
| `confidence` | ✅ | `high`, `medium`, `low` |
| `cwe` | ✅ | Array of `CWE-XXX` identifiers |
| `owasp` | ✅ | Array of `AXX:2025` OWASP Top 10 identifiers |

### Step 5: Write `patterns.yaml`

```yaml
rules:
  - id: <CATEGORY>-<VULN>-<NNN>
    name: <kebab-case-rule-name>
    description: "<One sentence describing what this rule detects>"
    severity: high
    confidence: medium
    cwe: [CWE-XXX]
    owasp: [AXX:2025]
    languages: [javascript, typescript]
    patterns:
      # Regex pattern — matches raw source text
      - regex: "(?i)dangerous_pattern\\s*\\("

      # AST pattern — matches syntax tree nodes
      - ast_node_type: "call_expression"
        function: "dangerous_function"
```

**Pattern writing tips:**

- Use `(?i)` prefix for case-insensitive regex
- Escape backslashes: `\\.` in YAML = `\.` in regex
- Test your regex at [regex101.com](https://regex101.com/) with PCRE flavor
- Write multiple patterns per rule — they are OR'd together (any match triggers the rule)
- Prefer specific patterns over broad ones to reduce false positives
- Add a comment explaining non-obvious regex

### Step 6: Add a Test Case

Add a file demonstrating the vulnerability to `testdata/vulnerable/<language>/`:

```bash
# Example: testdata/vulnerable/javascript/prototype-pollution.js
```

### Step 7: Validate Your Skill

```bash
# Build and test against your example
make build
./bin/ice-tea scan testdata/vulnerable/<language>/<your-file> --severity low --verbose

# Should see a finding from your new rule
```

---

## 5. Contributing Go Code

### Project Layout

```
cmd/ice-tea/          → CLI entry point
internal/
  analyzer/
    pattern/          → Engine 1: AST/Regex pattern matching
    taint/            → Engine 2: Data flow / taint tracking
    llm/              → Engine 3: LLM reasoning engine
  cli/                → Cobra command definitions
  config/             → Config loading and validation
  finding/            → Finding struct and deduplication
  mcp/                → MCP server implementation
  parser/
    goparser/         → Go native AST parser
    treesitter/       → Tree-sitter multi-language parser
  reporter/           → Output formatters (console, JSON, SARIF, GitLab)
  scanner/            → Scan orchestration and file walking
  skill/              → SKILL loader and rule index
skills/               → Built-in SKILL definitions
examples/             → Intentionally vulnerable test projects
testdata/             → Unit test fixtures
docs/                 → Documentation
```

### Writing Tests

Every new function or code path should include unit tests:

```go
// internal/analyzer/pattern/matcher_test.go
func TestMyNewFeature(t *testing.T) {
    t.Run("detects vulnerable pattern", func(t *testing.T) {
        // Arrange
        src := []byte(`vulnerable_code_here`)
        
        // Act
        findings := matcher.Match(src)
        
        // Assert
        assert.Len(t, findings, 1)
        assert.Equal(t, "MY-RULE-001", findings[0].RuleID)
    })
    
    t.Run("does not flag safe pattern", func(t *testing.T) {
        src := []byte(`safe_code_here`)
        findings := matcher.Match(src)
        assert.Empty(t, findings)
    })
}
```

Run tests before submitting:
```bash
go test ./...
go vet ./...
```

---

## 6. Reporting Bugs & False Positives

### Bug Reports

Open a GitHub Issue with the following template:

```markdown
**Bug Type:** False Positive / False Negative / Crash / Other

**Ice Tea Version:** (run `./bin/ice-tea version`)

**Operating System:** (e.g. Ubuntu 22.04, macOS 14, Windows 11)

**Command Used:**
```
./bin/ice-tea scan <path> <flags>
```

**Expected Behavior:**
<What should happen>

**Actual Behavior:**
<What actually happens — include full output/error>

**Minimal Reproduction Code:**
```<language>
// Paste the smallest code snippet that triggers the issue
```
```

### Reporting a False Positive

A false positive is when Ice Tea flags code as vulnerable when it is actually safe. To report:

1. Create an issue with the `false-positive` label
2. Include the flagged code snippet
3. Explain why it is safe
4. Optionally, submit a PR tightening the regex or AST pattern

### Reporting a False Negative

A false negative is when vulnerable code is not detected. To report:

1. Create an issue with the `false-negative` label
2. Include the vulnerable code snippet
3. Identify which rule _should_ have caught it (or propose a new rule)

---

## 7. Pull Request Process

1. **Sync with upstream** before starting work:
   ```bash
   git fetch upstream
   git rebase upstream/main
   ```

2. **Make your changes** on a feature branch.

3. **Ensure all tests pass**:
   ```bash
   go test ./...
   go vet ./...
   ```

4. **For new SKILLs**: Confirm your skill fires against a test case:
   ```bash
   ./bin/ice-tea scan testdata/vulnerable/ --severity low
   ```

5. **Write a clear PR description** using this template:

   ```markdown
   ## Summary
   Brief description of what this PR adds or fixes.

   ## Type of Change
   - [ ] New SKILL definition
   - [ ] Bug fix (false positive / false negative)
   - [ ] New feature
   - [ ] Documentation update
   - [ ] Other: ___

   ## Skill Details (if applicable)
   - **Vulnerability**: (e.g. Prototype Pollution)
   - **Category**: (e.g. web/)
   - **CWE(s)**: CWE-1321
   - **OWASP**: A08:2025
   - **Languages**: JavaScript, TypeScript
   - **Rules added**: 6

   ## Testing
   - [ ] New rule fires on `testdata/vulnerable/<lang>/<file>`
   - [ ] New rule does NOT fire on safe code
   - [ ] `go test ./...` passes
   ```

6. **Request a review** — a maintainer will review within a few days.

7. **Address review comments** by pushing new commits to the same branch.

---

## 8. Style Guidelines

### SKILL Files

- Use **American English** spelling in all SKILL content
- `description` field: start with a verb (e.g., "Detects...", "Identifies...")
- Keep regex patterns focused — prefer specific patterns over broad `.*` catches
- Always include both a vulnerable and safe code example in `SKILL.md`
- Version starts at `1.0.0`, bump minor version on significant pattern changes

### Go Code

- Follow standard Go formatting: `gofmt -w .`
- Variable and function names: `camelCase` (local), `PascalCase` (exported)
- Error handling: always check errors; never use `_` for errors from security-relevant functions
- Prefer table-driven tests
- Comments should explain *why*, not *what*

### Commit Messages

Use [Conventional Commits](https://www.conventionalcommits.org/) format:

```
feat(skill): add prototype-pollution detection for JavaScript
fix(pattern): tighten SQLi regex to reduce false positives in Go
docs(usage): update CLI reference with new --confidence flag
test(sqli): add Python SQLAlchemy false-positive test case
```

---

## 9. Skill ID Naming Convention

Each rule in `patterns.yaml` must have a unique `id` field. Use this format:

```
<CATEGORY>-<ABBREV>-<NNN>
```

| Part | Description | Example |
|------|-------------|---------|
| `CATEGORY` | 3-6 char category prefix (all caps) | `AUTH`, `INJ`, `WEB`, `CRYPTO`, `FS`, `API`, `INFRA`, `LOG`, `MEM`, `CLOUD`, `AND`, `NET` |
| `ABBREV` | 2-5 char vulnerability abbreviation | `SQL`, `XSS`, `CSRF`, `TLS`, `BOF`, `SSRF` |
| `NNN` | 3-digit sequential number within the skill | `001`, `002`, `003` |

**Examples:**

| ID | Meaning |
|----|---------|
| `AUTH-JWT-001` | Auth category, JWT vulnerability, first rule |
| `INJ-SSTI-003` | Injection category, SSTI, third rule |
| `CLOUD-AWS-002` | Cloud category, AWS, second rule |
| `MEM-BOF-001` | Memory category, buffer overflow, first rule |

**Uniqueness:** IDs must be globally unique across all `patterns.yaml` files in the repository. Check existing IDs before adding new ones:

```bash
grep -r "^\s*- id:" skills/ | awk -F': ' '{print $2}' | sort | uniq -d
# Should output nothing (no duplicates)
```

---

## Questions?

Open a [GitHub Discussion](https://github.com/zakirkun/ice-tea/discussions) or a GitHub Issue labeled `question`. We're happy to help you get started!
