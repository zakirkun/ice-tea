---
name: JWT Without Expiration
version: 1.0.0
description: Detects JWTs issued without expiration claims, creating tokens that remain valid indefinitely.
tags: [api, jwt, session, owasp-a07]
languages: [javascript, typescript, python, go, java]
severity: high
confidence: high
cwe: [CWE-613]
owasp: [A07:2025]
---

# JWT Without Expiration

## Overview
JWTs without expiration (`exp` claim) or with very long expiration never become invalid unless the secret is rotated. This means:
- Stolen tokens remain valid forever
- Deprovisioned users can still authenticate
- No way to force re-authentication without key rotation

## Remediation
- Set short-lived access tokens: 15 minutes to 1 hour
- Use refresh tokens with longer validity for user experience
- Implement token revocation list for immediate invalidation
