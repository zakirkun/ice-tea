# Ice Tea — Complete Usage Guide

Ice Tea is a multi-engine Static Application Security Testing (SAST) scanner. This guide covers every way to run it — from a quick one-liner to full CI/CD pipeline integration.

---

## Table of Contents

1. [Installation & Build](#1-installation--build)
2. [Quick Start](#2-quick-start)
3. [CLI Reference](#3-cli-reference)
4. [Output Formats](#4-output-formats)
5. [Configuration File](#5-configuration-file)
6. [LLM Deep Reasoning](#6-llm-deep-reasoning)
7. [Scanning Example Projects](#7-scanning-example-projects)
8. [CI/CD Integration](#8-cicd-integration)
9. [MCP Agent Mode](#9-mcp-agent-mode)
10. [Exit Codes](#10-exit-codes)
11. [Troubleshooting](#11-troubleshooting)

---

## 1. Installation & Build

### Prerequisites

| Requirement | Version | Notes |
|------------|---------|-------|
| Go | 1.21+ | [Download](https://go.dev/dl/) |
| C Compiler | Any | Required for Tree-sitter CGO bindings |
| Git | Any | |

**Linux / macOS:**
```bash
# Install GCC (if not present)
# Ubuntu/Debian: sudo apt install build-essential
# macOS: xcode-select --install

git clone https://github.com/zakirkun/ice-tea.git
cd ice-tea
make build

./bin/ice-tea version
```

**Windows (PowerShell):**
```powershell
# Install GCC via MSYS2 (https://www.msys2.org/) or TDM-GCC
# Then add the bin directory to PATH

git clone https://github.com/zakirkun/ice-tea.git
cd ice-tea
go build -o bin/ice-tea.exe ./cmd/ice-tea

.\bin\ice-tea.exe version
```

### Verify Installation

```
ice-tea version
# ice-tea v1.0.0 (go1.21, linux/amd64)
```

---

## 2. Quick Start

### Scan a Directory

```bash
# Scan current project
./bin/ice-tea scan .

# Scan a specific path
./bin/ice-tea scan ./src

# Scan a single file
./bin/ice-tea scan ./handlers/auth.go
```

### Scan with Severity Filter

```bash
# Only show high and critical findings (default: medium+)
./bin/ice-tea scan . --severity high

# Show everything including info-level findings
./bin/ice-tea scan . --severity info
```

### Scan a Specific Language

```bash
# Only scan Go files
./bin/ice-tea scan . --language go

# Only scan JavaScript and TypeScript
./bin/ice-tea scan . --language javascript,typescript
```

### Scan the Included Vulnerable Examples

```bash
# Scan all example projects
./bin/ice-tea scan ./examples --severity low

# Scan only the Go vulnerable app
./bin/ice-tea scan ./examples/vulnerable-go --severity low

# Scan only the PHP app
./bin/ice-tea scan ./examples/vulnerable-php --language php

# Scan Python app with verbose output
./bin/ice-tea scan ./examples/vulnerable-python --verbose
```

---

## 3. CLI Reference

### Global Flags

These flags work with any subcommand:

| Flag | Description | Default |
|------|-------------|---------|
| `--config <path>` | Path to config file | `.ice-tea.yaml` |
| `--verbose, -v` | Enable verbose debug logging | `false` |
| `--log-level <level>` | Log level: `debug`, `info`, `warn`, `error`, `fatal` | `info` |
| `--no-color` | Disable ANSI color output (for pipes/logs) | `false` |

### `scan` Subcommand

```
./bin/ice-tea scan [path] [flags]
```

| Flag | Short | Description | Default |
|------|-------|-------------|---------|
| `--format` | `-f` | Output format: `console`, `json`, `sarif`, `gitlab` | `console` |
| `--output` | `-o` | Write output to file (empty = stdout) | `""` |
| `--severity` | `-s` | Minimum severity: `info`, `low`, `medium`, `high`, `critical` | `medium` |
| `--confidence` | | Minimum confidence: `low`, `medium`, `high` | `medium` |
| `--exclude-dir` | | Comma-separated directories to skip | `node_modules,vendor,.git` |
| `--exclude-file` | | Comma-separated file glob patterns to skip | `*.test` |
| `--language` | | Restrict to specific languages (comma-separated) | auto-detect |
| `--concurrency` | `-c` | Parallel worker count | `4` |
| `--skills-dir` | | Load additional SKILL definitions from this path | `""` |
| `--enable-llm` | | Enable LLM Engine for false-positive filtering | `false` |

### `version` Subcommand

```bash
./bin/ice-tea version
```

### `mcp` Subcommand

```bash
# Start as MCP (Model Context Protocol) server for AI agent integrations
./bin/ice-tea mcp
```

---

## 4. Output Formats

### Console (Default)

Human-readable, color-coded terminal output. Best for local development.

```bash
./bin/ice-tea scan ./src
```

```
[CRITICAL] [CWE-89] SQL Injection
  File: src/handlers/user.go:47
  Rule: GO-SQLI-001 (injection/sql-injection)
  Code: query := "SELECT * FROM users WHERE id = " + userID
  ──────────────────────────────────────────────────────
[HIGH]     [CWE-798] Hardcoded Secret
  File: src/config/config.go:12
  Rule: AUTH-SEC-003 (auth/hardcoded-secrets)
  Code: const jwtSecret = "supersecret123"

  Summary: 2 findings (1 critical, 1 high)
```

### JSON

Machine-readable output for custom tooling and integrations.

```bash
./bin/ice-tea scan ./src --format json --output findings.json
```

```json
{
  "version": "1.0.0",
  "summary": { "total": 2, "critical": 1, "high": 1 },
  "findings": [
    {
      "id": "GO-SQLI-001",
      "rule_id": "GO-SQLI-001",
      "name": "sql-injection-concat",
      "severity": "critical",
      "confidence": "high",
      "file": "src/handlers/user.go",
      "line": 47,
      "cwe": ["CWE-89"],
      "owasp": ["A03:2025"],
      "description": "SQL query built via string concatenation with user input"
    }
  ]
}
```

### SARIF (GitHub Advanced Security)

Standard SARIF 2.1.0 format, supported by GitHub, Azure DevOps, and VS Code.

```bash
./bin/ice-tea scan ./src --format sarif --output results.sarif
```

Upload to GitHub:
```yaml
# .github/workflows/security.yml
- name: Upload SARIF results
  uses: github/codeql-action/upload-sarif@v3
  with:
    sarif_file: results.sarif
```

### GitLab SAST JSON

Compatible with GitLab's Security Dashboard.

```bash
./bin/ice-tea scan ./src --format gitlab --output gl-sast-report.json
```

---

## 5. Configuration File

Create `.ice-tea.yaml` in your project root to avoid passing flags every time. The CLI flags always take precedence over the config file.

### Full Configuration Reference

```yaml
# .ice-tea.yaml

# Output Settings
output:
  format: "sarif"          # console | json | sarif | gitlab
  file: "results.sarif"    # Empty string = stdout
  color: true              # ANSI colors (disable for CI logs)

# Scanner Settings
scan:
  severity: "medium"       # Minimum severity to report
  confidence: "medium"     # Minimum confidence to report
  concurrency: 8           # Worker goroutines (tune to CPU count)
  languages: []            # Empty = auto-detect all supported languages
  skills_dir: ""           # Additional custom skills directory

# Exclusions
exclude:
  dirs:
    - "vendor"
    - "node_modules"
    - ".git"
    - "testdata"
    - "dist"
    - "build"
    - "coverage"
  files:
    - "*.min.js"
    - "*_test.go"
    - "*.generated.go"
    - "*.pb.go"
    - "mocks/**"

# LLM False-Positive Engine
llm:
  enabled: false           # Set to true to activate
  provider: "openai"       # openai (more providers planned)
  model: "gpt-4o"          # Any OpenAI model with function calling
  # API key is read from OPENAI_API_KEY environment variable
```

### Minimal Config for CI

```yaml
# .ice-tea.yaml (minimal CI config)
output:
  format: "sarif"
  file: "results.sarif"
scan:
  severity: "high"
  concurrency: 8
exclude:
  dirs: ["vendor", "node_modules", "testdata"]
```

---

## 6. LLM Deep Reasoning

Engine 3 uses an LLM to verify static pattern matches against their full code context, dramatically reducing false positives.

### Setup

```bash
# Linux/macOS
export OPENAI_API_KEY="sk-proj-..."

# Windows PowerShell
$env:OPENAI_API_KEY = "sk-proj-..."
```

### Running with LLM

```bash
# Enable LLM for a local scan
./bin/ice-tea scan ./src --enable-llm

# LLM + SARIF output for CI
./bin/ice-tea scan ./src --enable-llm --format sarif --output results.sarif

# LLM with higher severity threshold (reduces LLM API calls)
./bin/ice-tea scan ./src --enable-llm --severity high
```

### How It Works

1. Engine 1 (Pattern Matcher) finds a potential match
2. Engine 2 (Taint Tracker) validates data flow from source to sink
3. Engine 3 (LLM) is invoked with:
   - The finding's code snippet (with surrounding context)
   - The `SKILL.md` content as system context
   - A structured prompt asking: *"Is this a real vulnerability or false positive?"*
4. The LLM response filters out false positives before reporting

### Cost Optimization Tips

- Use `--severity high` to reduce the number of findings sent to the LLM
- Use `--confidence high` to only send high-confidence matches
- Use `--language go` to limit scanning scope in polyglot projects
- The `gpt-4o-mini` model works well for most patterns and costs ~10x less

---

## 7. Scanning Example Projects

The `examples/` directory contains four intentionally vulnerable applications for testing scanner rules.

### Vulnerable Go REST API

```bash
# Expected: SQL injection, command injection, path traversal, SSRF,
#           hardcoded JWT secret, insecure cookie, weak crypto, open redirect
./bin/ice-tea scan ./examples/vulnerable-go --severity low
```

### Vulnerable PHP Web Application

```bash
# Expected: SQL injection, XSS, LFI, file upload bypass,
#           XXE, command injection, MD5 passwords
./bin/ice-tea scan ./examples/vulnerable-php --severity low --language php
```

### Vulnerable Python (Flask) Application

```bash
# Expected: SSTI, SQL injection, pickle deserialization, 
#           open redirect, CORS misconfiguration, debug mode
./bin/ice-tea scan ./examples/vulnerable-python --severity low --language python
```

### Vulnerable Node.js (Express) Application

```bash
# Expected: Prototype pollution, NoSQL injection, SSRF,
#           JWT weakness, ReDoS, XSS, command injection
./bin/ice-tea scan ./examples/vulnerable-nodejs --severity low --language javascript,typescript
```

### Expected Detection Summary

Running `./bin/ice-tea scan ./examples --severity low --format json --output report.json` should detect **30+ distinct findings** across all four projects, covering 20+ unique vulnerability categories.

---

## 8. CI/CD Integration

### GitHub Actions

```yaml
# .github/workflows/ice-tea.yml
name: Ice Tea Security Scan

on:
  push:
    branches: [main, develop]
  pull_request:
    branches: [main]

jobs:
  security-scan:
    name: SAST Scan
    runs-on: ubuntu-latest
    permissions:
      contents: read
      security-events: write

    steps:
      - name: Checkout repository
        uses: actions/checkout@v4

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.21'

      - name: Install GCC
        run: sudo apt-get install -y gcc

      - name: Build Ice Tea
        run: make build

      - name: Run Security Scan
        run: |
          ./bin/ice-tea scan ./src \
            --format sarif \
            --output results.sarif \
            --severity high \
            --exclude-dir testdata,vendor
        # Exit code 1 means findings found — decide whether to fail the build
        continue-on-error: true

      - name: Upload SARIF to GitHub Security
        uses: github/codeql-action/upload-sarif@v3
        if: always()
        with:
          sarif_file: results.sarif
```

### GitLab CI

```yaml
# .gitlab-ci.yml
ice-tea-sast:
  stage: test
  image: golang:1.21
  script:
    - apt-get update -qq && apt-get install -y gcc
    - make build
    - ./bin/ice-tea scan . --format gitlab --output gl-sast-report.json --severity high
  artifacts:
    reports:
      sast: gl-sast-report.json
    paths:
      - gl-sast-report.json
    expire_in: 1 week
  allow_failure: true
```

### Pre-commit Hook

```bash
# .git/hooks/pre-commit
#!/bin/bash
set -e

echo "Running Ice Tea security scan..."
./bin/ice-tea scan . --severity critical --format console

if [ $? -eq 1 ]; then
  echo "Critical vulnerabilities found. Commit blocked."
  exit 1
fi
```

---

## 9. MCP Agent Mode

Ice Tea can run as a Model Context Protocol (MCP) server, allowing Claude, ChatGPT, and other AI agents to call it as a tool.

```bash
# Start the MCP server (listens on stdio by default)
./bin/ice-tea mcp
```

When used with Claude Desktop, add to `claude_desktop_config.json`:
```json
{
  "mcpServers": {
    "ice-tea": {
      "command": "/path/to/ice-tea",
      "args": ["mcp"]
    }
  }
}
```

The MCP server exposes a `scan_project` tool that the agent can call to audit code on demand.

---

## 10. Exit Codes

| Code | Meaning |
|------|---------|
| `0` | Scan completed. No findings above the severity threshold. |
| `1` | Scan completed. One or more findings above the threshold were found. |
| `2` | Fatal error — bad arguments, target not found, configuration error. |

In CI/CD pipelines, you can use these codes as quality gates:

```bash
# Fail the pipeline only on critical findings
./bin/ice-tea scan . --severity critical
if [ $? -eq 1 ]; then
    echo "Critical security issues found — blocking deployment"
    exit 1
fi
```

---

## 11. Troubleshooting

### Build Error: `cgo: C compiler "gcc" not found`

Install a C compiler:
```bash
# Ubuntu/Debian
sudo apt-get install build-essential

# macOS
xcode-select --install

# Windows: Install MSYS2 from https://www.msys2.org/
# Then: pacman -S mingw-w64-x86_64-gcc
# Add C:\msys64\mingw64\bin to PATH
```

### No Findings on Known-Vulnerable Code

1. Check the language is supported: `--verbose` shows which files are parsed
2. Verify the severity threshold — default is `medium`, use `--severity low`
3. Check the skill covers that language: review `skills/<category>/<vuln>/patterns.yaml`
4. The file may be excluded: check `--exclude-dir` and `--exclude-file`

### LLM Not Filtering False Positives

1. Verify `OPENAI_API_KEY` is exported in the environment
2. Run with `--verbose` to see LLM request/response logs
3. Check your OpenAI account quota
4. Try a simpler model: `--config` with `llm.model: "gpt-4o-mini"`

### Too Many False Positives (Without LLM)

1. Raise the confidence threshold: `--confidence high`
2. Raise the severity threshold: `--severity high`
3. Enable LLM filtering: `--enable-llm`
4. If a specific rule is noisy, consider filtering by rule ID in your config

### Slow Scan on Large Projects

1. Increase concurrency: `--concurrency 16`
2. Limit to changed files in CI (use `git diff --name-only`)
3. Restrict language scope: `--language go` (skips parsing unsupported files)
4. Exclude large generated directories: `--exclude-dir dist,build,coverage`

---

## Supported Languages

| Language | Extension | Parser |
|----------|-----------|--------|
| Go | `.go` | `go/ast` (stdlib) |
| JavaScript | `.js`, `.mjs` | Tree-sitter |
| TypeScript | `.ts`, `.tsx` | Tree-sitter |
| Python | `.py` | Tree-sitter |
| Java | `.java` | Tree-sitter |
| PHP | `.php` | Tree-sitter |
| Ruby | `.rb` | Tree-sitter |
| Rust | `.rs` | Tree-sitter |
| C | `.c`, `.h` | Tree-sitter |
| C++ | `.cpp`, `.cc`, `.cxx` | Tree-sitter |
| YAML | `.yaml`, `.yml` | Regex |
| Generic | Any | Regex-only |

---

## Skill Coverage Summary

Ice Tea ships with **82 SKILL definitions** across **12 categories**, covering **456+ detection rules**:

| Category | Skills | Coverage Highlights |
|----------|--------|---------------------|
| `auth/` | 10 | JWT, hardcoded secrets, session fixation, OAuth, MFA bypass |
| `web/` | 16 | XSS, SSRF, CSRF, CORS, clickjacking, prototype pollution, WebSocket |
| `injection/` | 9 | SQLi, CMDi, XXE, SSTI, LDAP, XPath, NoSQL, log injection |
| `crypto/` | 8 | Weak hash, weak cipher, insecure TLS, hardcoded IV, insecure RSA |
| `api/` | 7 | BOLA/IDOR, BFLA, rate limiting, GraphQL, mass assignment |
| `fs/` | 5 | Path traversal, file upload, zip slip, unsafe temp files |
| `infra/` | 5 | Debug mode, default creds, Docker, Kubernetes, hardcoded IPs |
| `logging/` | 3 | Sensitive data in logs, insufficient logging, verbose errors |
| `memory/` | 5 | Buffer overflow, use-after-free, null deref, integer overflow, format string |
| `cloud/` | 5 | AWS, GCP, Azure misconfigs, IaC security, secrets in env |
| `android/` | 5 | Insecure storage, exported components, WebView JS, intent redirect |
| `network/` | 4 | Cleartext traffic, insecure socket, DNS rebinding, extended SSRF |
