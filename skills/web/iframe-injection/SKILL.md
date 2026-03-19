---
confidence: high
cwe:
    - CWE-79
description: Detects user-controlled content injected into iframe src attributes, enabling page embedding of malicious content.
languages:
    - javascript
    - typescript
    - php
    - python
    - kotlin
    - dart
    - zig
    - elixir
name: iframe Injection
owasp:
    - A03:2025
severity: high
tags:
    - web
    - iframe
    - xss
    - owasp-a03
version: 1.0.0
---

# iframe Injection

## Overview
Injecting user-controlled URLs into `<iframe src="...">` allows attackers to embed arbitrary external content in the application's page, enabling:
- Content spoofing (fake login forms within trusted domain)
- Clickjacking of inner content
- Cross-site cookie access in older browsers

## Remediation
- Validate and whitelist iframe src URLs against allowed domains
- Use `sandbox` attribute on iframes
- Set CSP `frame-src` to restrict allowed iframe sources
