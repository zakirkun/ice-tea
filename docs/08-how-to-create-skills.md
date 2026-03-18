# How to Create Custom SKILLs for Ice Tea

Ice Tea uses a declarative **SKILL** system to define security vulnerabilities, search patterns, and remediation instructions. This system allows you to add detection capabilities without writing any Go code!

A "SKILL" is simply a directory containing two files:
1. `SKILL.md`: The documentation and metadata.
2. `patterns.yaml`: The actual AST or Regex detection rules.

## Directory Structure

You can create a new SKILL anywhere inside the `skills/` folder. For example, to detect insecure cookie configurations:

```
skills/
└── web/
    └── insecure-cookies/
        ├── SKILL.md
        └── patterns.yaml
```

## 1. Writing `SKILL.md`

This file uses **Markdown** with **YAML Frontmatter**. 

- The **YAML Frontmatter** defines the metadata Ice Tea parses (severity, CWEs, tags)
- The **Markdown Body** defines the explanation and remediation advice that will be displayed to users, and serves as the context for the **LLM Engine**.

```markdown
---
name: Insecure Cookie Configuration
version: 1.0.0
description: Detects cookies created without the 'Secure' and 'HttpOnly' flags.
tags: [web, cookie, session, owasp-a05]
languages: [go, javascript]
severity: high
confidence: medium
cwe: [CWE-614, CWE-1004]
owasp: [A05:2025]
---

# Insecure Cookie

## Overview
Cookies used for sensitive configuration or session management must be marked with the `Secure` flag (so they are only sent over HTTPS) and the `HttpOnly` flag (so they cannot be accessed by client-side Javascript).

## Remediation
Always set `Secure: true` and `HttpOnly: true` when defining cookies.
```

## 2. Writing `patterns.yaml`

This file contains the rules Ice Tea's **Pattern Matching Engine** (Engine 1) will execute against the source code syntax tree.

You define a list of `rules`, each containing one or more `patterns`.

```yaml
rules:
  - id: COOKIE-GO-01
    name: go-insecure-cookie
    description: "Go http.Cookie missing Secure or HttpOnly flags"
    severity: high
    confidence: high
    languages: [go]
    patterns:
      # Use ast_node_type to match structural syntax
      - ast_node_type: composite_literal
        # For Go AST, we can look for specific structs
        function: "http.Cookie"
      
      # Or you can use a regular expression!
      # This regex looks for cookies being constructed.
      - regex: '(?i)http\.Cookie\s*\{[^\}]*(?:Secure:\s*false|HttpOnly:\s*false)[^\}]*\}'

  - id: COOKIE-JS-01
    name: js-insecure-cookie-express
    description: "Express res.cookie missing secure options"
    severity: medium
    languages: [javascript, typescript]
    patterns:
      - ast_node_type: call_expression
        function: "res.cookie"
      - regex: 'res\.cookie\s*\(\s*[^,]+,\s*[^,]+,\s*\{[^\}]*secure\s*:\s*false'
```

### Supported Pattern Matchers

Ice Tea supports two primary types of matchers:

1. **`regex`**: A standard PCRE regular expression string. It will match against the raw source code text. Use `(?i)` to make the regex case-insensitive.
2. **`ast_node_type`**: Matches a specific node in the Abstract Syntax Tree (AST). 
   - *Example Types*: `call_expression`, `binary_expression`, `composite_literal`.
   - You can pair this with `function` to specifically match function calls (e.g. `function: "eval"`).

### How It Works Together
When Ice Tea runs `ice-tea scan`:
1. It walks the directory and parses files into ASTs.
2. It loads `SKILL.md` metadata into the rule index.
3. It evaluates the `patterns.yaml` AST matchers against the codebase.
4. If a match is found, the metadata and markdown description are bundled up and sent to the **LLM Engine** (Engine 3) to ask: *"Is this a true positive based on this code context?"*
5. The LLM responds, and the final filtered finding is printed in the SARIF or Console output!
