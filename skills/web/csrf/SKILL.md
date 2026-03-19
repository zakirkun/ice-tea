---
confidence: medium
cwe:
    - CWE-352
description: Detects web forms and state-changing endpoints that lack CSRF token protection.
languages:
    - javascript
    - typescript
    - python
    - php
    - java
    - go
    - ruby
    - kotlin
    - dart
    - zig
    - elixir
name: Cross-Site Request Forgery (CSRF)
owasp:
    - A01:2025
severity: high
tags:
    - csrf
    - web
    - auth
    - owasp-a01
version: 1.0.0
---

# Cross-Site Request Forgery (CSRF)

## Overview
CSRF tricks authenticated users into unknowingly submitting requests to a web application they're logged into. Without CSRF tokens, an attacker can craft a malicious webpage that triggers state-changing actions (password change, fund transfer, account deletion) in the victim's session.

## Detection Strategy
- HTML forms without CSRF token hidden input
- POST endpoints in Express/Flask/Go without CSRF middleware
- CORS policy that allows cross-origin requests with credentials

## Remediation
- Use framework CSRF protection middleware (csurf, Flask-WTF, Django CSRF, etc.)
- Verify `Origin` or `Referer` header for same-origin
- Use `SameSite=Strict` cookies as a defense-in-depth measure

**Vulnerable (HTML form):**
```html
<form method="POST" action="/transfer">
    <input name="amount" value="1000">
    <input type="submit">
</form>
```

**Safe:**
```html
<form method="POST" action="/transfer">
    <input type="hidden" name="_csrf" value="{{ csrf_token }}">
    <input name="amount" value="1000">
    <input type="submit">
</form>
```
