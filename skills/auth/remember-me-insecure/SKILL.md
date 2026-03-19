---
name: Insecure Remember-Me / Persistent Session
version: 1.0.0
description: Detects insecure persistent login implementations using predictable tokens or insufficient expiration policies.
tags: [auth, remember-me, session, owasp-a07]
languages: [javascript, typescript, python, php, java, go]
severity: high
confidence: medium
cwe: [CWE-613]
owasp: [A07:2025]
---

# Insecure Remember-Me

## Overview
"Remember Me" functionality that uses predictable tokens or stores sensitive data in persistent cookies enables long-lived session hijacking:
- Tokens derived from username+timestamp (predictable)
- Tokens stored without server-side revocation capability
- Tokens valid even after password change
- Long-lived tokens with no rotation on use

## Remediation
- Use cryptographically random 32+ byte tokens stored server-side
- Store only a token reference in the cookie, not user data
- Invalidate remember-me tokens on password change and logout
- Rotate the token on each use to prevent replay attacks
