---
name: Sensitive URL in Referrer Header
version: 1.0.0
description: Detects pages that expose sensitive parameters (tokens, IDs) in URLs that get leaked via the Referer header to external resources.
tags: [web, referrer, information-disclosure, owasp-a05]
languages: [javascript, typescript, python, go, java, php]
severity: medium
confidence: medium
cwe: [CWE-598]
owasp: [A05:2025]
---

# Sensitive URL Referrer Leakage

## Overview
When a page loads external resources (images, scripts, analytics) and the page URL contains sensitive parameters (tokens, session IDs), the `Referer` header sent to external servers exposes these parameters. Password reset tokens and session tokens in URLs are particularly dangerous.

## Remediation
- Never put sensitive tokens in URL query parameters
- Use `Referrer-Policy: no-referrer` or `strict-origin` meta tag/header
- Use POST requests for sensitive data submission
