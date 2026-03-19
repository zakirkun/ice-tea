---
confidence: medium
cwe:
    - CWE-80
description: Detects user-controlled input reflected in HTML without proper encoding, enabling HTML injection that may not execute scripts but can redirect or spoof content.
languages:
    - python
    - php
    - javascript
    - typescript
    - java
    - go
    - ruby
    - kotlin
    - dart
    - zig
    - elixir
name: HTML Injection
owasp:
    - A03:2025
severity: medium
tags:
    - injection
    - html-injection
    - xss
    - owasp-a03
version: 1.0.0
---

# HTML Injection

## Overview
HTML injection allows attackers to insert arbitrary HTML markup into web pages. Unlike XSS, HTML injection may not involve script execution (e.g., blocked by CSP) but can still:
- Redirect users via injected `<meta refresh>`
- Spoof content with injected forms (phishing)
- Manipulate page structure to mislead users
- In some contexts, escalate to XSS

## Remediation
- HTML-encode all user output: `htmlspecialchars()`, `html.escape()`, `template.HTMLEscapeString()`
- Use auto-escaping template engines
