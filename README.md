<div align="center">
  <h1>🍵 Ice Tea Security Scanner</h1>
  <p>An AI-Powered DevOps Security Tool Agent written in Go.</p>
</div>

Ice Tea is an advanced Static Application Security Testing (SAST) tool that combines lightning-fast **AST Pattern Matching** via Tree-sitter with **Deep Reasoning AI** support (via OpenAI).

By identifying vulnerabilities accurately and filtering them through localized AI agents, Ice Tea drastically reduces false positives while integrating seamlessly into your CI/CD pipelines.

---

## 🚀 Features

- **Multi-Engine Scanning**:
  1. *Quick Pattern Matching (AST & Regex)* using Go's Native parser and bindings to Tree-sitter.
  2. *LLM Deep Reasoning* automatically evaluates static matches using an LLM (OpenAI) to confirm if the match is a true vulnerability or a false positive.
- **Language Support**: 
  - Go, JavaScript, TypeScript, Python, Ruby, PHP, Java, Rust, C, and C++.
- **Progressive SKILL System**: Declare vulnerability rules and context effortlessly using `SKILL.md` (Markdown) and `patterns.yaml`.
- **Outputs for Every Environment**: Pretty-formatted Console, standard JSON, GitLab SAST JSON, or SARIF.
- **Agentic Workflows**: Integrated Model Context Protocol (MCP) server for Claude / ChatGPT agent integrations!

---

## 🛠️ Installation

Requirements: `go 1.21+` and a C compiler (for tree-sitter CGO bindings).

```bash
git clone https://github.com/zakirkun/ice-tea.git
cd ice-tea

# Build the binary
make build

# Check the version
./bin/ice-tea version
```

---

## 📖 Quick Start

### Basic Scan
Scan a project directory to find vulnerabilities using the built-in SKILL rules:
```bash
./bin/ice-tea scan ./path/to/project
```

### Scan with AI False-Positive Filtering
Provide your API Key and pass the `--enable-llm` flag. Ice Tea will first execute static pattern matching, and then independently consult the LLM to verify each finding against its deeper AST context.
```bash
export OPENAI_API_KEY="sk-proj-xyz..."
./bin/ice-tea scan ./src --enable-llm
```

### SARIF Output for CI/CD
Generate SARIF files compatible with GitHub Advanced Security.
```bash
./bin/ice-tea scan ./src --format sarif --output results.sarif
```
---

## 🧠 Building Custom SKILLs

Ice Tea includes an ecosystem of built-in SKILLs located in the `skills/` directory covering SQLi, XSS, Path Traversal, Weak Crypto, Hardcoded Secrets, SSRF, Command Injection, and Deserialization vulnerabilities.

You can create your own custom detection rules and integrate them cleanly. Please read our documentation on creating rules:
- 👉 **[How to Create SKILLs](docs/08-how-to-create-skills.md)**

## 🤝 Contributing

We welcome contributions! Whether it's adding a new SKILL pattern, implementing another language's tree-sitter parser, or improving the LLM prompt contexts, please refer to the [CONTRIBUTING.md](CONTRIBUTING.md) guide.

## 📄 License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
