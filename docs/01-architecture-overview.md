# Ice Tea - Architecture Overview

## What is Ice Tea?

**Ice Tea** is an AI-powered DevOps Security Tools Agent built in Golang. It performs static application security testing (SAST) by combining traditional AST-based code analysis with AI/LLM-powered deep reasoning. The tool is designed to integrate seamlessly into CI/CD pipelines (GitHub Actions, GitLab Runner) and can also operate as an MCP (Model Context Protocol) tool for agentic AI workflows.

## High-Level Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                        ICE TEA CLI                          │
│  (Cobra/Viper CLI Framework)                                │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  ┌──────────┐  ┌──────────────┐  ┌────────────────────┐    │
│  │  Config   │  │  File Walker │  │   SKILL Loader     │    │
│  │  Manager  │  │  & Filter    │  │   (Knowledge Mgr)  │    │
│  └────┬─────┘  └──────┬───────┘  └─────────┬──────────┘    │
│       │               │                     │               │
│  ┌────▼─────────────────▼─────────────────────▼──────────┐  │
│  │              SCAN ENGINE (Orchestrator)                │  │
│  │                                                       │  │
│  │  ┌─────────────┐  ┌──────────────┐  ┌──────────────┐ │  │
│  │  │  AST Parser  │  │  Pattern     │  │  LLM Engine  │ │  │
│  │  │  (Go native  │  │  Matcher     │  │  (AI Deep    │ │  │
│  │  │  + Tree-     │  │  (Rules,     │  │   Reasoning) │ │  │
│  │  │  Sitter)     │  │  Regex)      │  │              │ │  │
│  │  └──────┬──────┘  └──────┬───────┘  └──────┬───────┘ │  │
│  │         │                │                  │         │  │
│  │  ┌──────▼────────────────▼──────────────────▼───────┐ │  │
│  │  │           FINDING AGGREGATOR & DEDUPLICATOR      │ │  │
│  │  └──────────────────────┬───────────────────────────┘ │  │
│  └─────────────────────────┼─────────────────────────────┘  │
│                            │                                │
│  ┌─────────────────────────▼─────────────────────────────┐  │
│  │              REPORT GENERATOR                         │  │
│  │  (SARIF, JSON, GitLab SAST, Console)                  │  │
│  └───────────────────────────────────────────────────────┘  │
│                                                             │
├─────────────────────────────────────────────────────────────┤
│               INTEGRATION LAYER                             │
│  ┌────────────┐  ┌───────────┐  ┌─────────────────────┐    │
│  │  GitHub     │  │  GitLab   │  │  MCP Server         │    │
│  │  Actions    │  │  CI       │  │  (stdio/HTTP)       │    │
│  └────────────┘  └───────────┘  └─────────────────────┘    │
└─────────────────────────────────────────────────────────────┘
```

## Core Design Principles

### 1. Multi-Engine Detection
Ice Tea uses a **three-layer detection approach**:
- **Layer 1 — Static Pattern Matching**: Fast rule-based checks using AST traversal, regex patterns, and known vulnerability signatures. This catches common, well-defined vulnerabilities instantly.
- **Layer 2 — Data Flow / Taint Analysis**: Tracks how untrusted input propagates through the code from sources to sinks, detecting injection vulnerabilities that require understanding data flow.
- **Layer 3 — LLM Deep Reasoning**: AI-powered semantic analysis using chain-of-thought prompting. The LLM examines code context, business logic flaws, and subtle vulnerabilities that static rules cannot catch.

### 2. SKILL-Based Knowledge Architecture
Vulnerability knowledge is packaged as **SKILLs** — modular folders containing:
- `SKILL.md`: Instructions, detection patterns, and remediation guidance
- Optional scripts, test cases, or reference materials

SKILLs are loaded on-demand (progressive disclosure) based on the file type and patterns detected, keeping memory usage efficient.

### 3. Language-Agnostic Parsing
- **Go code**: Parsed natively using `go/parser` and `go/ast` (zero-dependency)
- **Other languages** (JS, Python, Java, etc.): Parsed via **Tree-Sitter** with language-specific grammars
- Common AST traversal interface abstracts away parser differences

### 4. CI/CD Native
- Outputs in **SARIF 2.1.0** format for GitHub Code Scanning
- Outputs in **GitLab SAST JSON** format for GitLab merge request widgets
- Single binary with no runtime dependencies — easy to install in any CI runner
- Cross-platform: Linux, macOS, Windows (amd64, arm64)

### 5. MCP Integration
- Can run as an **MCP Server** (via stdio or HTTP transport)
- Exposes scan capabilities as MCP Tools
- Vulnerability SKILLs exposed as MCP Resources
- Enables agentic AI workflows where LLM agents can invoke security scans

## Scan Flow

```
1. CLI invoked with target directory/file
2. Config loaded (flags > env vars > config file > defaults)
3. File walker discovers files, applying exclusion filters
4. For each file:
   a. Detect language from extension
   b. Select appropriate AST parser
   c. Parse into AST
   d. Load relevant SKILLs based on language + tags
   e. Run Layer 1: Pattern matching against AST
   f. Run Layer 2: Taint analysis (source → sink tracking)
   g. Run Layer 3: LLM reasoning (if enabled, with CoT prompting)
   h. Aggregate findings, deduplicate, assign severity/confidence
5. Generate report (SARIF, JSON, console, GitLab format)
6. Exit with appropriate code (0 = clean, 1 = findings)
```

## Concurrency Model

Go goroutines are used for parallel file processing:
- Worker pool pattern with configurable concurrency level
- Each worker processes files independently
- Findings are collected via channels into the aggregator
- LLM calls are rate-limited to respect API quotas
