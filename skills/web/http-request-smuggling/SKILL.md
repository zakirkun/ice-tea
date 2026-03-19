---
confidence: medium
cwe:
    - CWE-444
description: Detects HTTP response headers and server configurations vulnerable to request smuggling via Transfer-Encoding/Content-Length discrepancies.
languages:
    - javascript
    - typescript
    - python
    - go
    - java
    - kotlin
    - dart
    - zig
    - elixir
name: HTTP Request Smuggling
owasp:
    - A03:2025
severity: critical
tags:
    - web
    - http-smuggling
    - owasp-a03
version: 1.0.0
---

# HTTP Request Smuggling

## Overview
HTTP Request Smuggling exploits discrepancies in how front-end (proxy/CDN) and back-end servers parse HTTP request boundaries. By crafting requests with ambiguous `Content-Length` and `Transfer-Encoding` headers, attackers can:
- Bypass security controls at the proxy layer
- Poison the request queue affecting other users
- Achieve XSS, redirect stealing, and cache poisoning

## Detection Strategy
Look for custom HTTP header handling that processes both Content-Length and Transfer-Encoding, or proxy configurations that forward headers without normalization.

## Remediation
- Use HTTP/2 end-to-end where possible
- Normalize HTTP headers at the proxy layer
- Configure backend to reject ambiguous requests
