---
name: Missing Security Headers
version: 1.0.0
description: Detects missing or misconfigured HTTP security headers that protect against common browser-based attacks.
tags: [security-headers, web, csp, hsts, owasp-a05]
languages: [javascript, typescript, python, go, java, php, ruby]
severity: medium
confidence: medium
cwe: [CWE-16]
owasp: [A05:2025]
---

# Missing Security Headers

## Overview
HTTP security headers instruct the browser to enable additional protections. Missing headers expose applications to:
- **Missing HSTS**: Downgrade attacks, MITM on first connection
- **Missing CSP**: XSS amplification
- **Missing X-Content-Type-Options**: MIME sniffing attacks
- **Missing Referrer-Policy**: Sensitive URL leakage in Referer header
- **Missing Permissions-Policy**: Feature abuse (camera, microphone, geolocation)

## Recommended Headers
```
Strict-Transport-Security: max-age=31536000; includeSubDomains
Content-Security-Policy: default-src 'self'
X-Content-Type-Options: nosniff
X-Frame-Options: DENY
Referrer-Policy: strict-origin-when-cross-origin
Permissions-Policy: geolocation=(), microphone=(), camera=()
```
