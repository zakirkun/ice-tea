---
name: API Key Exposed in URL
version: 1.0.0
description: Detects API keys and tokens passed as URL query parameters, which are logged in server logs, browser history, and Referer headers.
tags: [auth, api-key, information-disclosure, owasp-a07]
languages: [javascript, typescript, python, go, java, php, ruby]
severity: high
confidence: high
cwe: [CWE-598]
owasp: [A07:2025]
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
