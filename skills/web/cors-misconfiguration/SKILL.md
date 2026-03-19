---
confidence: high
cwe:
    - CWE-942
description: Detects insecure Cross-Origin Resource Sharing configurations that allow unauthorized cross-origin access.
languages:
    - javascript
    - typescript
    - python
    - go
    - java
    - php
    - kotlin
    - dart
    - zig
    - elixir
name: CORS Misconfiguration
owasp:
    - A05:2025
severity: high
tags:
    - cors
    - web
    - owasp-a05
version: 1.0.0
---

# CORS Misconfiguration

## Overview
CORS misconfigurations allow attackers to make cross-origin requests that read sensitive API responses from a victim's browser. Common issues:
1. **Wildcard with credentials**: `Access-Control-Allow-Origin: *` combined with credentials is rejected by browsers but misconfiguring the origin reflection is not
2. **Reflected origin**: Origin header value reflected directly without validation
3. **Null origin**: Allowing `null` origin (sandbox iframes)
4. **Subdomain wildcard**: Allowing `*.example.com` which includes attacker-controlled subdomains

## Remediation
- Maintain an explicit allowlist of trusted origins
- Never reflect the `Origin` header directly without validation
- Never allow `null` origin in production
- Do not combine `Access-Control-Allow-Credentials: true` with broad origin policies
