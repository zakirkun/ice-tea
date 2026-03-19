---
confidence: medium
cwe:
    - CWE-359
description: Detects CSS-based browser history sniffing patterns and sensitive data in browser history.
languages:
    - javascript
    - typescript
    - kotlin
    - dart
    - zig
    - elixir
name: Browser History / Visited Link Theft
owasp:
    - A05:2025
severity: medium
tags:
    - web
    - privacy
    - information-disclosure
    - owasp-a05
version: 1.0.0
---

# Browser History Theft

## Overview
CSS-based history theft uses `:visited` pseudo-class timing differences to detect which URLs a user has visited. While modern browsers have mitigated this, JavaScript APIs can still expose visited URLs through timing channels or getComputedStyle attacks. Additionally, sensitive data in URLs persists in browser history.

## Remediation
- Avoid putting sensitive data in URLs
- Use `window.history.replaceState()` to remove sensitive parameters after processing
- Set `Cache-Control: no-store` for pages with sensitive URL parameters
