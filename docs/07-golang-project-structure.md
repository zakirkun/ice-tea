# Golang Project Structure & Best Practices

## Overview

Ice Tea is implemented in Go for its advantages:
- **Single binary deployment** — no runtime dependencies
- **Cross-platform compilation** with `GOOS`/`GOARCH` environment variables
- **Built-in concurrency** via goroutines and channels
- **Strong standard library** for AST parsing, networking, and crypto
- **Fast compilation and execution**

## Project Layout

```
ice-tea/
├── cmd/                          # Application entry points
│   └── ice-tea/
│       └── main.go               # CLI entry point
│
├── internal/                     # Private application code
│   ├── cli/                      # CLI commands (Cobra)
│   │   ├── root.go               # Root command & global flags
│   │   ├── scan.go               # Scan command
│   │   ├── mcp.go                # MCP server command
│   │   └── version.go            # Version command
│   │
│   ├── config/                   # Configuration management (Viper)
│   │   └── config.go             # Config loading & validation
│   │
│   ├── scanner/                  # Core scan engine
│   │   ├── engine.go             # Scan orchestrator
│   │   ├── walker.go             # File discovery & filtering
│   │   └── worker.go             # Concurrent file processing
│   │
│   ├── parser/                   # AST parsing backends
│   │   ├── parser.go             # Common parser interface
│   │   ├── goparser/             # Go native parser
│   │   │   └── parser.go
│   │   └── treesitter/           # Tree-Sitter multi-language parser
│   │       ├── parser.go
│   │       └── languages.go      # Language registry
│   │
│   ├── analyzer/                 # Vulnerability analysis engines
│   │   ├── analyzer.go           # Common analyzer interface
│   │   ├── pattern/              # Engine 1: Static pattern matching
│   │   │   ├── matcher.go
│   │   │   └── rules.go
│   │   ├── taint/                # Engine 2: Taint/data flow analysis
│   │   │   ├── tracker.go
│   │   │   ├── sources.go
│   │   │   └── sinks.go
│   │   └── llm/                  # Engine 3: LLM deep reasoning
│   │       ├── engine.go
│   │       ├── prompt.go
│   │       └── providers/        # LLM provider adapters
│   │           ├── openai.go
│   │           ├── anthropic.go
│   │           └── ollama.go     # Local LLM support
│   │
│   ├── skill/                    # SKILL management
│   │   ├── loader.go             # SKILL discovery & loading
│   │   ├── index.go              # SKILL metadata indexing
│   │   ├── validator.go          # SKILL validation & sanitization
│   │   └── types.go              # SKILL data types
│   │
│   ├── finding/                  # Finding management
│   │   ├── finding.go            # Finding data types
│   │   ├── aggregator.go         # Multi-engine result aggregation
│   │   └── deduplicator.go       # Finding deduplication
│   │
│   ├── reporter/                 # Output report generation
│   │   ├── reporter.go           # Common reporter interface
│   │   ├── sarif.go              # SARIF v2.1.0 output
│   │   ├── gitlab.go             # GitLab SAST JSON output
│   │   ├── json.go               # Raw JSON output
│   │   └── console.go            # Terminal output (colored)
│   │
│   └── mcp/                      # MCP server implementation
│       ├── server.go             # MCP server main logic
│       ├── transport.go          # stdio/HTTP transport
│       ├── tools.go              # MCP tool definitions
│       ├── resources.go          # MCP resource definitions
│       └── prompts.go            # MCP prompt templates
│
├── skills/                       # Built-in vulnerability SKILLs
│   ├── injection/
│   │   ├── sql-injection/
│   │   │   └── SKILL.md
│   │   ├── command-injection/
│   │   │   └── SKILL.md
│   │   └── xss/
│   │       └── SKILL.md
│   ├── crypto/
│   │   ├── weak-hash/
│   │   │   └── SKILL.md
│   │   └── insecure-tls/
│   │       └── SKILL.md
│   ├── auth/
│   │   ├── hardcoded-creds/
│   │   │   └── SKILL.md
│   │   └── weak-auth/
│   │       └── SKILL.md
│   └── ...
│
├── docs/                         # Documentation
│   ├── 01-architecture-overview.md
│   ├── 02-ai-agent-skills.md
│   └── ...
│
├── testdata/                     # Test fixtures
│   ├── vulnerable/               # Known vulnerable code samples
│   │   ├── go/
│   │   ├── javascript/
│   │   └── python/
│   └── safe/                     # Known safe code samples
│       ├── go/
│       ├── javascript/
│       └── python/
│
├── .github/                      # GitHub Actions
│   └── workflows/
│       ├── ci.yml                # Build, test, lint
│       └── release.yml           # Release binaries
│
├── go.mod                        # Go module definition
├── go.sum                        # Dependency checksums
├── Makefile                      # Build automation
├── .goreleaser.yml               # Cross-platform release config
├── .ice-tea.yaml                 # Default configuration
└── README.md
```

## CLI Framework: Cobra + Viper

### Cobra Command Structure

```
ice-tea
├── scan        # Scan code for vulnerabilities
├── mcp         # Start MCP server
├── skills      # Manage vulnerability skills
│   ├── list    # List available skills
│   ├── info    # Show skill details
│   └── validate # Validate a skill
├── rules       # Manage detection rules
│   └── list    # List all rules
├── version     # Show version info
└── help        # Help information
```

### Viper Configuration Precedence

```
Priority (highest to lowest):
1. Command-line flags       --severity=high
2. Environment variables    ICE_TEA_SEVERITY=high
3. Config file              .ice-tea.yaml
4. Defaults                 severity: medium
```

### Configuration File (`.ice-tea.yaml`)

```yaml
# .ice-tea.yaml - Ice Tea configuration
scan:
  target: "."
  severity: medium
  confidence: medium
  concurrency: 4
  timeout: 300s

exclude:
  dirs:
    - vendor
    - node_modules
    - .git
    - build
    - dist
    - testdata
  files:
    - "*.min.js"
    - "*.min.css"
    - "*.generated.go"
  extensions:
    - .md
    - .txt
    - .png
    - .jpg
    - .svg

languages:
  - go
  - javascript
  - typescript
  - python

skills:
  dir: "./skills"
  external_dirs: []
  trust_external: false

llm:
  enabled: false
  provider: "openai"       # openai, anthropic, ollama
  model: "gpt-4o-mini"
  api_key_env: "ICE_TEA_LLM_API_KEY"
  max_tokens: 4096
  rate_limit: 10           # requests per minute
  timeout: 30s
  cache: true

output:
  format: "console"        # console, sarif, gitlab, json
  file: ""                 # output file (empty = stdout)
  verbose: false
  color: true
```

## Key Go Dependencies

| Package | Purpose | Version |
|---------|---------|---------|
| `github.com/spf13/cobra` | CLI framework | v1.8+ |
| `github.com/spf13/viper` | Configuration management | v1.18+ |
| `github.com/tree-sitter/go-tree-sitter` | Multi-language AST parsing | latest |
| `github.com/fatih/color` | Terminal color output | v1.16+ |
| `go.uber.org/zap` | Structured logging | v1.27+ |
| `golang.org/x/tools` | Go SSA/analysis packages | latest |
| `github.com/stretchr/testify` | Testing assertions | v1.9+ |

## Cross-Platform Build

### Makefile Targets

```makefile
VERSION := $(shell git describe --tags --always)
LDFLAGS := -ldflags "-X main.version=$(VERSION)"

.PHONY: build
build:
	go build $(LDFLAGS) -o bin/ice-tea ./cmd/ice-tea

.PHONY: build-all
build-all:
	GOOS=linux   GOARCH=amd64 go build $(LDFLAGS) -o bin/ice-tea-linux-amd64 ./cmd/ice-tea
	GOOS=linux   GOARCH=arm64 go build $(LDFLAGS) -o bin/ice-tea-linux-arm64 ./cmd/ice-tea
	GOOS=darwin  GOARCH=amd64 go build $(LDFLAGS) -o bin/ice-tea-darwin-amd64 ./cmd/ice-tea
	GOOS=darwin  GOARCH=arm64 go build $(LDFLAGS) -o bin/ice-tea-darwin-arm64 ./cmd/ice-tea
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o bin/ice-tea-windows-amd64.exe ./cmd/ice-tea

.PHONY: test
test:
	go test -race -cover ./...

.PHONY: lint
lint:
	golangci-lint run ./...
```

## Security Best Practices for the Tool Itself

1. **No Code Execution**: Never execute scanned code — only parse ASTs
2. **API Key Protection**: Read LLM API keys from environment variables only
3. **Input Sanitization**: Validate all file paths against path traversal
4. **Dependency Auditing**: Run `govulncheck` regularly on our own dependencies
5. **Memory Safety**: Use proper goroutine lifecycle management; prevent goroutine leaks
6. **Error Handling**: Never expose internal errors to untrusted callers
7. **SKILL Sandboxing**: External SKILLs are loaded as data-only, never executed as code
8. **Race Condition Prevention**: Use `go test -race` in CI; protect shared state with mutexes/channels
9. **Log Sanitization**: Never log API keys, tokens, or sensitive code content
10. **Minimal Dependencies**: Keep the dependency tree small to reduce supply chain risk
