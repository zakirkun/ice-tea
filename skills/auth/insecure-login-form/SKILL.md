---
name: Insecure Login Form
version: 1.0.0
description: Detects login forms served over HTTP, with autocomplete enabled for password fields, or without CSRF protection.
tags: [auth, login-form, https, owasp-a07]
languages: [javascript, typescript, php, python, generic]
severity: high
confidence: high
cwe: [CWE-319, CWE-352]
owasp: [A07:2025]
---

# Insecure Login Form

## Overview
Login forms must be secured from several angles:
1. **HTTP form submission**: Credentials transmitted in cleartext
2. **Password autocomplete enabled**: Stored credentials can be accessed by XSS
3. **No CSRF protection**: Login form submissions forged from other origins
4. **Remember credential prompt disabled**: Should be disabled in some contexts

## Remediation
- Serve login forms and POST targets exclusively over HTTPS
- Add `autocomplete="new-password"` to prevent unintended credential storage
- Add CSRF tokens to login forms
