---
confidence: high
cwe:
    - CWE-598
description: Detects API keys and tokens passed as URL query parameters, which are logged in server logs, browser history, and Referer headers.
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
name: API Key Exposed in URL
owasp:
    - A07:2025
severity: high
tags:
    - auth
    - api-key
    - information-disclosure
    - owasp-a07
version: 1.0.0
---

# API Key Exposed in URL

## Overview
API keys in URL query parameters appear in:
- Server access logs (nginx, Apache, CloudFront)
- Browser history
- Referer headers sent to third-party analytics
- Shared URLs (when users copy the URL from their browser)
- Proxy logs and CDN access logs

## Remediation
- Pass API keys in HTTP headers: `Authorization: Bearer <token>` or `X-API-Key: <key>`
- Never log or store full URLs with API keys
- Rotate any keys that appeared in URLs
