# How to Create Custom SKILLs for Ice Tea

Ice Tea uses a declarative **SKILL** system to define security vulnerabilities, search patterns, and remediation guidance. You can add new detection capabilities without writing any Go code.

A SKILL is a directory containing exactly two files:

1. **`SKILL.md`** — Metadata frontmatter + human-readable vulnerability documentation
2. **`patterns.yaml`** — Detection rules (regex and/or AST node matchers)

---

## Directory Structure

Skills live inside the `skills/` directory, organized by category. Ice Tea automatically discovers all `patterns.yaml` files recursively.

```
skills/
├── auth/
│   ├── hardcoded-secrets/
│   │   ├── SKILL.md
│   │   └── patterns.yaml
│   └── session-fixation/
│       ├── SKILL.md
│       └── patterns.yaml
├── injection/
│   └── sql-injection/
│       ├── SKILL.md
│       └── patterns.yaml
└── myteam/                  ← You can create new top-level categories
    └── custom-vuln/
        ├── SKILL.md
        └── patterns.yaml
```

You can also point Ice Tea at a completely separate directory using:
```bash
./bin/ice-tea scan ./src --skills-dir ./my-company-skills
```

---

## Available Skill Categories

| Category | Focus Area |
|----------|-----------|
| `auth/` | Authentication & authorization flaws |
| `injection/` | All injection vulnerability classes |
| `web/` | Browser/HTTP-layer vulnerabilities |
| `crypto/` | Cryptographic algorithm misuse |
| `fs/` | Filesystem and file handling issues |
| `api/` | REST/GraphQL API security |
| `infra/` | Infrastructure & configuration |
| `logging/` | Security logging failures |
| `memory/` | Memory safety (C/C++) |
| `cloud/` | Cloud platform misconfigurations |
| `android/` | Android mobile security |
| `network/` | Network-level vulnerabilities |

---

## Step 1: Writing `SKILL.md`

This file uses **Markdown** with a **YAML Frontmatter** block. The frontmatter is parsed by Ice Tea's loader; the markdown body is passed to the LLM Engine as context for false-positive verification.

### Full Example

```markdown
---
name: Insecure Cookie Configuration
version: 1.0.0
description: Detects cookies set without Secure, HttpOnly, or SameSite attributes.
tags: [web, cookie, session, owasp-a05]
languages: [go, javascript, typescript, python, php, java, ruby]
severity: medium
confidence: high
cwe: [CWE-614, CWE-1004]
owasp: [A05:2025]
---

# Insecure Cookie Configuration

## Overview
Cookies that store session tokens must be configured with:
- **Secure**: Cookie only transmitted over HTTPS
- **HttpOnly**: Inaccessible to JavaScript — prevents XSS token theft
- **SameSite=Strict**: Prevents cross-site request forgery

## Detection Strategy
Look for `http.SetCookie` calls (Go) or `res.cookie()` calls (Express) that do not
include all three security attributes.

## Remediation
Always set all three security attributes on session cookies.

**Vulnerable (Go):**
```go
http.SetCookie(w, &http.Cookie{
    Name:  "session",
    Value: token,
    // Missing Secure, HttpOnly, SameSite
})
```

**Safe (Go):**
```go
http.SetCookie(w, &http.Cookie{
    Name:     "session",
    Value:    token,
    Secure:   true,
    HttpOnly: true,
    SameSite: http.SameSiteStrictMode,
})
```
```

### Frontmatter Field Reference

| Field | Required | Type | Description |
|-------|----------|------|-------------|
| `name` | ✅ | string | Human-readable SKILL name |
| `version` | ✅ | semver | Start at `1.0.0` |
| `description` | ✅ | string | One sentence starting with "Detects" |
| `tags` | ✅ | array | Lowercase kebab-case labels |
| `languages` | ✅ | array | Language identifiers (see list below), or `[generic]` |
| `severity` | ✅ | enum | `critical` / `high` / `medium` / `low` / `info` |
| `confidence` | ✅ | enum | `high` / `medium` / `low` |
| `cwe` | ✅ | array | `["CWE-XXX"]` — use official CWE identifiers |
| `owasp` | ✅ | array | `["AXX:2025"]` — OWASP Top 10 2025 reference |

### Supported Language Identifiers

```
go, javascript, typescript, python, java, php, ruby, rust, c, cpp, yaml, dart, kotlin, zig, perl, elixir, generic
```

Use `generic` for patterns (regex-only) that apply to any file type.

---

## Step 2: Writing `patterns.yaml`

This file contains the actual detection logic. Each file defines a list of `rules`, and each rule contains one or more `patterns`.

### Full Schema

```yaml
rules:
  - id: CATEGORY-ABBREV-NNN      # Unique ID — see naming convention below
    name: kebab-case-rule-name   # Short slug for the rule
    description: "One sentence describing what this rule detects"
    severity: high               # critical | high | medium | low | info
    confidence: medium           # high | medium | low
    cwe: [CWE-XXX]               # Array of CWE identifiers
    owasp: [AXX:2025]            # Array of OWASP Top 10 identifiers
    languages: [go, python]      # Languages this rule applies to
    patterns:
      # ── Pattern Type 1: Regex ──────────────────────────────────────────
      # Matched against the raw source code text of each file
      - regex: "(?i)dangerous_call\\s*\\("

      # ── Pattern Type 2: AST Node ───────────────────────────────────────
      # Matched against the Abstract Syntax Tree
      - ast_node_type: "call_expression"
        function: "dangerous_call"      # Optional: match specific function name
      
      # ── Pattern Type 3: Import Path (Go only) ─────────────────────────
      - import_path: "crypto/md5"
```

### Pattern Types In Depth

#### `regex` — Raw Source Matching

The regex is matched against the raw text of every file whose language is in the rule's `languages` list.

```yaml
patterns:
  # Case-insensitive match for MD5 in any context
  - regex: "(?i)\\bmd5\\s*\\("

  # Match SELECT with string concatenation (SQL injection indicator)
  - regex: "(?i)(SELECT|UPDATE|DELETE).*['\"]\\s*\\+\\s*\\w"

  # Match hardcoded AWS access key
  - regex: "AKIA[0-9A-Z]{16}"
```

**Tips:**
- Use `(?i)` prefix for case-insensitive matching
- Double-escape backslashes in YAML strings: `\\.` = one literal dot
- Test your regex on [regex101.com](https://regex101.com/) with PCRE flavor selected
- Anchor patterns (`^`, `$`) when you want line-level precision
- Use `\\b` word boundaries to avoid partial-word matches

#### `ast_node_type` — Syntax Tree Matching

Matches a specific type of AST node. Optionally restrict to a specific function name.

```yaml
patterns:
  # Match any call to eval()
  - ast_node_type: "call_expression"
    function: "eval"

  # Match Go composite literal for http.Cookie struct
  - ast_node_type: "composite_literal"
    function: "http.Cookie"

  # Match any binary expression (useful with taint tracking)
  - ast_node_type: "binary_expression"
```

**Common AST node types:**

| Node Type | Description |
|-----------|-------------|
| `call_expression` | Function/method call |
| `binary_expression` | `a + b`, `a && b`, etc. |
| `composite_literal` | Go struct/slice/map literal `T{...}` |
| `assignment_statement` | `x = y` |
| `import_declaration` | Import statement |
| `return_statement` | Return statement |

#### `import_path` — Go Import Matching (Go only)

Detects when a specific Go package is imported. Useful for flagging dangerous library usage.

```yaml
patterns:
  - import_path: "crypto/md5"
  - import_path: "math/rand"     # Non-cryptographic random
  - import_path: "net/http"      # Pair with function match for context
```

### Rule ID Naming Convention

```
<CATEGORY>-<ABBREV>-<NNN>
```

| Segment | Format | Examples |
|---------|--------|---------|
| CATEGORY | 2-6 uppercase chars | `AUTH`, `INJ`, `WEB`, `CRYPTO`, `FS`, `API`, `INFRA`, `LOG`, `MEM`, `CLOUD`, `AND`, `NET` |
| ABBREV | 2-5 uppercase chars | `SQL`, `XSS`, `TLS`, `BOF`, `SSRF`, `JWT`, `SSTI` |
| NNN | 3-digit number | `001`, `002`, `003` |

IDs must be **globally unique** across all skills. Check for duplicates:
```bash
grep -rh "^\s*- id:" skills/ | awk '{print $3}' | sort | uniq -d
```

### Complete Example: XSS SKILL

`skills/web/xss/patterns.yaml`:

```yaml
rules:
  - id: WEB-XSS-001
    name: xss-innerhtml
    description: "Assignment to innerHTML with non-literal value — potential DOM XSS"
    severity: high
    confidence: medium
    cwe: [CWE-79]
    owasp: [A03:2025]
    languages: [javascript, typescript]
    patterns:
      - regex: "\\.innerHTML\\s*=\\s*(?!\\s*['\"])"

  - id: WEB-XSS-002
    name: xss-eval
    description: "Call to eval() with non-literal argument — code injection risk"
    severity: critical
    confidence: high
    cwe: [CWE-94, CWE-79]
    owasp: [A03:2025]
    languages: [javascript, typescript]
    patterns:
      - ast_node_type: "call_expression"
        function: "eval"
      - regex: "\\beval\\s*\\([^)]+\\)"

  - id: WEB-XSS-003
    name: xss-document-write
    description: "document.write() with user-controlled content"
    severity: high
    confidence: medium
    cwe: [CWE-79]
    owasp: [A03:2025]
    languages: [javascript, typescript]
    patterns:
      - ast_node_type: "call_expression"
        function: "document.write"
      - regex: "document\\.write\\s*\\([^)]*(?:location|hash|search|query|input|param)"
```

---

## Step 3: Test Your SKILL

### Verify Detection (True Positive)

```bash
# Build with your new skill
make build

# Scan your vulnerable test file
./bin/ice-tea scan testdata/vulnerable/<lang>/<file> \
  --severity low \
  --verbose

# Expected: at least one finding from your new rule ID
```

### Verify No False Positives

Create a safe version of the code and confirm it's not flagged:

```bash
./bin/ice-tea scan testdata/safe/<lang>/<file> --severity low
# Expected: zero findings from your rule
```

### Run Unit Tests

```bash
go test ./...
```

---

## Step 4: Add to `testdata/`

Add a minimal code snippet demonstrating the vulnerability to `testdata/vulnerable/<language>/`:

```bash
# Example file name: testdata/vulnerable/python/ssti.py
cat > testdata/vulnerable/python/ssti.py << 'EOF'
from flask import Flask, render_template_string, request
app = Flask(__name__)

@app.route("/greet")
def greet():
    name = request.args.get("name")
    return render_template_string(f"Hello {name}!")  # VULN: SSTI
EOF
```

---

## How the Engines Use Your SKILL

When `ice-tea scan` runs:

```
File parsed by Go/Tree-sitter parser
          ↓
Engine 1: Pattern Matcher
  └─ Reads patterns.yaml
  └─ Runs regex against raw source
  └─ Matches AST nodes in parsed tree
  └─ Finding generated if pattern matches
          ↓
Engine 2: Taint Tracker (optional)
  └─ Traces if matched source has data flow from user input to dangerous sink
          ↓
Engine 3: LLM Reasoning (if --enable-llm)
  └─ Sends code snippet + SKILL.md body as context
  └─ Asks: "Is this a real vulnerability or false positive?"
  └─ Filters out false positives before reporting
          ↓
Reporter: Output to console / JSON / SARIF / GitLab
```

---

## Checklist Before Submitting

- [ ] `SKILL.md` frontmatter is complete with all required fields
- [ ] `description` field starts with "Detects" or "Identifies"
- [ ] At least one vulnerable code example in `SKILL.md`
- [ ] At least one safe code example in `SKILL.md`
- [ ] All rule IDs in `patterns.yaml` follow the `CATEGORY-ABBREV-NNN` format
- [ ] Rule IDs are unique (no duplicates in the whole `skills/` tree)
- [ ] `languages` field is set to the correct language(s)
- [ ] Regex tested on [regex101.com](https://regex101.com/) with PCRE
- [ ] Vulnerable test file added to `testdata/vulnerable/<lang>/`
- [ ] `./bin/ice-tea scan testdata/vulnerable/` detects your new rule
- [ ] `go test ./...` passes
