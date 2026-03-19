---
confidence: medium
cwe:
    - CWE-1021
description: Detects missing X-Frame-Options or Content-Security-Policy frame-ancestors directives.
languages:
    - javascript
    - typescript
    - python
    - go
    - java
    - php
    - ruby
    - kotlin
    - dart
    - zig
    - elixir
name: Clickjacking (Missing Frame Protection)
owasp:
    - A05:2025
severity: medium
tags:
    - clickjacking
    - web
    - security-headers
    - owasp-a05
version: 1.0.0
---

# Clickjacking (Missing Frame Protection)

## Overview
Clickjacking embeds a target website in a transparent iframe on an attacker's page. The victim is tricked into clicking UI elements on the invisible iframe (e.g., approve a transfer, change settings, delete account).

## Detection Strategy
Look for HTTP responses that do not set either:
- `X-Frame-Options: DENY` or `X-Frame-Options: SAMEORIGIN`
- `Content-Security-Policy: frame-ancestors 'none'` or `'self'`

## Remediation
Use CSP `frame-ancestors` (modern, preferred) or `X-Frame-Options` (legacy).

```
X-Frame-Options: DENY
Content-Security-Policy: frame-ancestors 'none';
```
