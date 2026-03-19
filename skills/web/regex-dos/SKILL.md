---
confidence: medium
cwe:
    - CWE-1333
description: Detects catastrophically backtracking regular expressions applied to user-controlled input, causing CPU-intensive denial of service.
languages:
    - javascript
    - typescript
    - python
    - java
    - go
    - kotlin
    - dart
    - zig
    - elixir
name: Regular Expression Denial of Service (ReDoS)
owasp:
    - A06:2025
severity: high
tags:
    - redos
    - regex
    - dos
    - owasp-a06
version: 1.0.0
---

# Regular Expression Denial of Service (ReDoS)

## Overview
Certain regex patterns with nested quantifiers (e.g., `(a+)+`, `(a|aa)+`, `([a-zA-Z]+)*`) cause catastrophic backtracking when matched against specially crafted inputs. A single request with a malicious string can cause the regex engine to run for seconds, minutes, or hours, exhausting server CPU.

Classic vulnerable patterns:
- `(a+)+` — exponential backtracking
- `([a-z]+)*` — polynomial backtracking  
- `(a|a?)+` — ambiguous alternation

## Detection Strategy
Look for regular expressions with:
- Nested quantifiers: `(x+)+`, `(x*)*`
- Alternation inside quantifiers: `(a|b|ab)+`
- Applied to user-controlled strings

## Remediation
- Rewrite regex to avoid ambiguous patterns
- Use linear-time regex engines (RE2 via `re2` npm package or Go's `regexp`)
- Apply regex only to length-limited input
- Use `timeout` options where available

**Vulnerable:**
```js
const emailRegex = /^([a-zA-Z0-9])(([a-zA-Z0-9])*([._-])?)+@.../;
req.body.email.match(emailRegex); // ReDoS with crafted email
```
