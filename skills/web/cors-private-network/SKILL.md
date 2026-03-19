---
confidence: medium
cwe:
    - CWE-942
description: Detects CORS configurations that allow access from public origins to private network endpoints.
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
name: CORS Private Network Access Misconfiguration
owasp:
    - A05:2025
severity: high
tags:
    - web
    - cors
    - private-network
    - owasp-a05
version: 1.0.0
---

# CORS Private Network Access Misconfiguration

## Overview
Chrome's Private Network Access (PNA) restrictions protect internal services. However, misconfigured CORS headers (`Access-Control-Allow-Private-Network: true` without origin restriction) allow malicious public websites to make requests to internal APIs running on private IP ranges. This enables CSRF attacks against private services.

## Remediation
- Only allow specific trusted origins for private network access
- Implement authentication on all internal APIs
- Do not set `Access-Control-Allow-Private-Network: true` for wildcard origins
