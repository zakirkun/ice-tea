---
name: Web Cache Deception
version: 1.0.0
description: Detects server configurations that may serve personalized content under cacheable URLs, enabling web cache deception attacks.
tags: [web, cache-deception, owasp-a05]
languages: [javascript, typescript, python, go, java, php]
severity: high
confidence: low
cwe: [CWE-524]
owasp: [A05:2025]
---

# Web Cache Deception

## Overview
Web cache deception exploits a discrepancy between how the server and cache interpret URL paths. An attacker tricks a victim into visiting `https://example.com/profile/nonexistent.css` — the server ignores the `.css` extension and serves the authenticated profile page, but the CDN caches it as a public CSS file. The attacker then retrieves the cached authenticated response.

## Remediation
- Cache responses based on their Content-Type, not URL extension
- Never cache authenticated API responses
- Set `Cache-Control: no-store` for all authenticated/personalized content
