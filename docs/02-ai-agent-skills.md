# AI Agent Skills & SKILL.md System

## Overview

The **SKILL.md** standard is an open format for packaging reusable AI agent capabilities into plain Markdown files. Originally introduced by Anthropic (October 2025) and made an open standard (December 2025), it is now supported by all major AI coding tools: OpenAI Codex, GitHub Copilot, Claude Code, Gemini CLI, Cursor, and more.

In Ice Tea, SKILLs serve as modular **vulnerability knowledge bases** that the scanner loads on-demand during analysis.

## SKILL.md Structure

Each SKILL is a folder containing at minimum a `SKILL.md` file:

```
skills/
├── sql-injection/
│   ├── SKILL.md           # Main instruction file
│   ├── patterns.yaml      # Detection patterns
│   ├── examples/
│   │   ├── vulnerable.go  # Example vulnerable code
│   │   └── safe.go        # Example safe code
│   └── remediation.md     # Fix guidance
├── xss/
│   ├── SKILL.md
│   └── patterns.yaml
├── path-traversal/
│   ├── SKILL.md
│   └── patterns.yaml
└── ...
```

### SKILL.md Format

```markdown
---
name: SQL Injection Detection
version: 1.0.0
description: Detects SQL injection vulnerabilities across multiple languages
tags: [injection, sql, database, owasp-a05]
languages: [go, python, javascript, java]
severity: critical
confidence: high
cwe: [CWE-89]
owasp: [A05:2025]
---

# SQL Injection Detection

## Detection Instructions

Look for patterns where user-controlled input is concatenated or
interpolated into SQL query strings without parameterization.

### Vulnerable Patterns

1. String concatenation in SQL queries
2. fmt.Sprintf used to build SQL queries with user input
3. Template literals with user data in SQL strings
4. Raw SQL execution without prepared statements

### AST Patterns

- Function calls to `db.Query()`, `db.Exec()` where argument is
  a concatenated string or fmt.Sprintf result
- String concatenation (`+` operator) near SQL keywords
  (SELECT, INSERT, UPDATE, DELETE, WHERE)

### Taint Sources
- HTTP request parameters (r.URL.Query(), r.FormValue())
- Request body data
- Path parameters
- Headers

### Taint Sinks
- database/sql: Query, QueryRow, Exec, QueryContext, ExecContext
- GORM: Raw, Exec, Where (with string argument)
- sqlx: Query, Get, Select

## Remediation

Use parameterized queries / prepared statements:

**Vulnerable:**
```go
db.Query("SELECT * FROM users WHERE id = " + userInput)
```

**Safe:**
```go
db.Query("SELECT * FROM users WHERE id = $1", userInput)
```
```

## Progressive Disclosure

To manage context efficiently, Ice Tea uses progressive disclosure:

1. **Index Phase**: At startup, only SKILL names, descriptions, and tags are loaded into memory (lightweight metadata).
2. **Match Phase**: When scanning a file, the engine matches the file's language and detected patterns against SKILL tags.
3. **Load Phase**: Only matched SKILLs have their full instructions loaded for detailed analysis.

This prevents loading all vulnerability knowledge into every scan, reducing memory footprint and LLM token costs.

```
┌─────────────────────┐
│   SKILL Index        │  (~50 bytes per skill)
│   name + tags only   │
└──────────┬──────────┘
           │ match against file language/patterns
           ▼
┌─────────────────────┐
│   Matched SKILLs     │  (load full SKILL.md)
│   full instructions  │
└──────────┬──────────┘
           │ feed to analysis engine
           ▼
┌─────────────────────┐
│   Detection Result   │
└─────────────────────┘
```

## Security Considerations

> **Warning**: External/community SKILLs are untrusted by default.

Malicious SKILLs could contain **prompt injection** attacks that hijack the LLM analysis process. Ice Tea must implement safeguards:

1. **SKILL Validation**: Check SKILL structure conforms to expected schema
2. **Content Sanitization**: Strip or escape potentially dangerous prompt patterns
3. **Trust Levels**: Differentiate between built-in (trusted) and external (untrusted) SKILLs
4. **Digital Signatures**: Optionally verify SKILL authenticity via cryptographic signatures
5. **Sandbox Execution**: If SKILLs contain scripts, execute in sandboxed environments

## Built-in SKILLs (Planned)

Ice Tea will ship with built-in SKILLs covering OWASP Top 10:2025:

| SKILL | OWASP Category | CWE |
|-------|---------------|-----|
| Broken Access Control | A01:2025 | CWE-284, CWE-639 |
| Security Misconfiguration | A02:2025 | CWE-16, CWE-732 |
| Supply Chain Failures | A03:2025 | CWE-829, CWE-506 |
| Cryptographic Failures | A04:2025 | CWE-327, CWE-328 |
| Injection (SQL, CMD, XSS) | A05:2025 | CWE-79, CWE-89, CWE-78 |
| Insecure Design | A06:2025 | CWE-209, CWE-502 |
| Authentication Failures | A07:2025 | CWE-287, CWE-384 |
| Data Integrity Failures | A08:2025 | CWE-494, CWE-502 |
| Logging & Alerting Failures | A09:2025 | CWE-778, CWE-223 |
| Mishandled Exceptions | A10:2025 | CWE-754, CWE-391 |
