<div align="center">
  <h1>🍵 Ice Tea Security Scanner</h1>
  <p><strong>AI-Powered Static Application Security Testing (SAST) — written in Go</strong></p>
  <p>
    <a href="LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="MIT License"></a>
    <a href="go.mod"><img src="https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go" alt="Go 1.21+"></a>
    <img src="https://img.shields.io/badge/Languages-10-brightgreen" alt="10 Languages">
    <img src="https://img.shields.io/badge/Rules-456+-red" alt="456+ Rules">
    <img src="https://img.shields.io/badge/SKILLs-82-orange" alt="82 SKILLs">
  </p>
</div>

---

Ice Tea is an advanced SAST tool that combines **lightning-fast AST pattern matching** (via Tree-sitter) with **AI-powered deep reasoning** (via OpenAI) to find security vulnerabilities in your source code — with dramatically fewer false positives than traditional scanners.

```bash
$ ice-tea scan ./src --format console --severity high

[CRITICAL] [CWE-89] SQL Injection — src/handlers/user.go:47
  Rule: GO-SQLI-001 | Confidence: high
  Code: query := "SELECT * FROM users WHERE id = " + userID

[CRITICAL] [CWE-798] Hardcoded JWT Secret — src/config/config.go:12
  Rule: AUTH-SEC-003 | Confidence: high
  Code: const jwtSecret = "supersecret123"

[HIGH]     [CWE-295] InsecureSkipVerify TLS — src/client/http.go:31
  Rule: CRYPTO-TLS-001 | Confidence: high
  Code: TLSClientConfig: &tls.Config{InsecureSkipVerify: true}

Summary: 3 findings (2 critical, 1 high) in 0.4s
```

---

## Features

- **3-Engine Architecture**:
  1. **Pattern Matching Engine** — AST + Regex rules via Tree-sitter and Go's native `go/ast`
  2. **Taint Tracker** — Traces data flow from user-controlled sources to dangerous sinks
  3. **LLM Reasoning Engine** — Optionally consults OpenAI to verify each finding and eliminate false positives

- **10 Languages**: Go, JavaScript, TypeScript, Python, Java, PHP, Ruby, Rust, C, C++

- **82 Built-in SKILLs / 456+ Detection Rules** across 12 security domains:
  - Authentication, Injection, Web/API, Cryptography, Filesystem, Infrastructure, Logging, Memory Safety, Cloud, Android, Network

- **4 Output Formats**: Console, JSON, SARIF 2.1.0, GitLab SAST JSON

- **MCP Server**: Model Context Protocol integration for Claude / ChatGPT agentic workflows

- **CI/CD Ready**: Native SARIF output, configurable exit codes, `.ice-tea.yaml` config file

- **Extensible SKILL System**: Add custom detection rules in Markdown + YAML — no Go code required

---

## Installation

**Requirements:** Go 1.21+, GCC or Clang (for Tree-sitter CGO bindings)

```bash
git clone https://github.com/zakirkun/ice-tea.git
cd ice-tea
make build

# Verify
./bin/ice-tea version
```

**Windows (PowerShell):**
```powershell
# Install GCC via MSYS2: https://www.msys2.org/
git clone https://github.com/zakirkun/ice-tea.git
cd ice-tea
go build -o bin/ice-tea.exe ./cmd/ice-tea
.\bin\ice-tea.exe version
```

---

## Quick Start

### 1. Scan a Project

```bash
# Scan with default settings (medium+ severity)
./bin/ice-tea scan ./your-project

# Show all findings including low severity
./bin/ice-tea scan ./your-project --severity low

# Only critical findings to keep CI fast
./bin/ice-tea scan ./your-project --severity critical
```

### 2. SARIF Output for GitHub

```bash
./bin/ice-tea scan ./src --format sarif --output results.sarif
```

```yaml
# .github/workflows/security.yml
- name: Run Ice Tea
  run: ./bin/ice-tea scan ./src --format sarif --output results.sarif
  continue-on-error: true

- name: Upload SARIF
  uses: github/codeql-action/upload-sarif@v3
  with:
    sarif_file: results.sarif
```

### 3. AI False-Positive Filtering

```bash
export OPENAI_API_KEY="sk-proj-..."
./bin/ice-tea scan ./src --enable-llm --severity medium
```

### 4. Test Against Vulnerable Example Projects

```bash
# All 4 example apps (Go, PHP, Python, Node.js)
./bin/ice-tea scan ./examples --severity low

# Each individually
./bin/ice-tea scan ./examples/vulnerable-go
./bin/ice-tea scan ./examples/vulnerable-php --language php
./bin/ice-tea scan ./examples/vulnerable-python --language python
./bin/ice-tea scan ./examples/vulnerable-nodejs --language javascript
```

---

## SKILL Coverage

Ice Tea ships with **82 SKILL definitions** (456+ rules) covering:

| Category | Skills | Key Detections |
|----------|--------|----------------|
| **auth/** | 10 | Hardcoded secrets, JWT weaknesses, session fixation, insecure cookies, OAuth flaws, MFA bypass |
| **web/** | 16 | XSS, SSRF, CSRF, CORS misconfiguration, clickjacking, prototype pollution, ReDoS, WebSocket |
| **injection/** | 9 | SQL, Command, XXE, Template (SSTI), LDAP, XPath, NoSQL, Header/CRLF, Log injection |
| **crypto/** | 8 | Weak hash (MD5/SHA1), weak cipher (DES/RC4/ECB), insecure TLS, hardcoded IV, insecure RSA |
| **api/** | 7 | BOLA/IDOR, broken function-level auth, missing rate limiting, GraphQL security, mass assignment |
| **fs/** | 5 | Path traversal, unsafe file upload, Zip Slip, unsafe temp files |
| **infra/** | 5 | Debug mode in production, default credentials, insecure Docker/Kubernetes, hardcoded IPs |
| **logging/** | 3 | Sensitive data in logs, insufficient security logging, verbose error responses |
| **memory/** | 5 | Buffer overflow, use-after-free, null pointer dereference, integer overflow, format string |
| **cloud/** | 5 | AWS/GCP/Azure misconfigurations, secrets in env files, IaC security (Terraform) |
| **android/** | 5 | Insecure data storage, exported components, WebView JS bridge, intent redirection |
| **network/** | 4 | Cleartext HTTP, insecure sockets, DNS rebinding, extended SSRF (cloud metadata) |

---

## Creating Custom SKILLs

Adding a new detection rule requires only two files — no Go code needed:

```
skills/
└── myteam/
    └── my-custom-vuln/
        ├── SKILL.md      ← Metadata + remediation guidance
        └── patterns.yaml ← Detection rules (regex or AST)
```

**`SKILL.md`** (frontmatter + markdown):
```yaml
---
name: My Custom Vulnerability
version: 1.0.0
description: Detects unsafe use of dangerous_function()
tags: [custom, injection]
languages: [python]
severity: high
confidence: high
cwe: [CWE-78]
owasp: [A03:2025]
---
# My Custom Vulnerability
## Overview
...
## Remediation
...
```

**`patterns.yaml`**:
```yaml
rules:
  - id: CUSTOM-001
    name: dangerous-function-call
    description: "dangerous_function() called with user input"
    severity: high
    confidence: high
    languages: [python]
    patterns:
      - regex: "dangerous_function\\s*\\(.*request\\."
      - ast_node_type: "call_expression"
        function: "dangerous_function"
```

Then scan with your custom skills:
```bash
./bin/ice-tea scan ./src --skills-dir ./my-skills
```

📖 Read the full guide: [docs/08-how-to-create-skills.md](docs/08-how-to-create-skills.md)

---

## Configuration File

Place `.ice-tea.yaml` in your project root:

```yaml
output:
  format: "sarif"
  file: "results.sarif"

scan:
  severity: "medium"
  confidence: "medium"
  concurrency: 8

exclude:
  dirs: ["vendor", "node_modules", "testdata", "dist"]
  files: ["*.min.js", "*_test.go", "*.pb.go"]

llm:
  enabled: false
  provider: "openai"
  model: "gpt-4o"
```

---

## Exit Codes

| Code | Meaning |
|------|---------|
| `0` | No findings above threshold |
| `1` | Findings found above threshold |
| `2` | Fatal error (bad args, config issue, etc.) |

---

## Documentation

| Document | Description |
|----------|-------------|
| [Architecture Overview](docs/01-architecture-overview.md) | System design and engine pipeline |
| [AI Agent Skills](docs/02-ai-agent-skills.md) | How the SKILL system works |
| [AST Parsing](docs/03-ast-parsing.md) | Tree-sitter and Go AST internals |
| [Vulnerability Detection](docs/04-vulnerability-detection.md) | Pattern matching and taint tracking |
| [CI/CD Integration](docs/05-cicd-integration.md) | GitHub Actions, GitLab CI, Jenkins |
| [MCP Integration](docs/06-mcp-integration.md) | Claude / ChatGPT agent workflows |
| [How to Create SKILLs](docs/08-how-to-create-skills.md) | Step-by-step guide for custom rules |
| [**Usage Guide**](docs/09-usage-guide.md) | **Complete CLI reference and examples** |

---

## Contributing

Contributions are welcome — especially new SKILL definitions! See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines on:
- Adding new vulnerability patterns
- Reporting false positives
- Improving parser coverage
- Submitting pull requests

---

## License

This project is licensed under the **MIT License** — see the [LICENSE](LICENSE) file for details.

Copyright (c) 2026 Ice Tea Scanner Contributors
