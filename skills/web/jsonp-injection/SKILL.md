---
confidence: high
cwe:
    - CWE-79
description: Detects JSONP endpoints that reflect user-controlled callback names without validation, enabling XSS.
languages:
    - javascript
    - typescript
    - python
    - php
    - go
    - java
    - kotlin
    - dart
    - zig
    - elixir
name: JSONP Injection
owasp:
    - A03:2025
severity: high
tags:
    - web
    - jsonp
    - xss
    - owasp-a03
version: 1.0.0
---

# JSONP Callback Injection

## Overview
JSONP (JSON with Padding) endpoints that reflect a user-supplied callback name without validation allow attackers to inject arbitrary JavaScript. JSONP is a legacy cross-origin technique that modern applications should replace with CORS.

**Attack:** `https://api.example.com/data?callback=alert(document.cookie)//`
**Response:** `alert(document.cookie)//({"user": "admin"})` — executes JS

## Remediation
- Replace JSONP with CORS
- If JSONP must be maintained, validate callback against `[a-zA-Z0-9._]+`
- Set `Content-Type: application/javascript` not `text/html`
