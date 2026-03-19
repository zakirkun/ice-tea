---
name: WebSocket Security Issues
version: 1.0.0
description: Detects insecure WebSocket implementations including missing origin validation, lack of authentication, and message injection risks.
tags: [websocket, web, auth, owasp-a01]
languages: [javascript, typescript, python, go, java]
severity: high
confidence: medium
cwe: [CWE-345, CWE-284]
owasp: [A01:2025]
---

# WebSocket Security Issues

## Overview
WebSocket connections bypass same-origin policy — browsers send cookies with WebSocket upgrade requests, making them vulnerable to:
1. **Cross-Site WebSocket Hijacking (CSWSH)**: No `Origin` header validation allows cross-origin connections
2. **Missing authentication**: WebSocket handlers that don't verify the user is logged in
3. **Message injection**: User-controlled messages echoed back without sanitization

## Detection Strategy
- WebSocket `upgrade` handlers that don't verify the `Origin` header
- WebSocket handlers that don't check session/token before processing messages
- Message data broadcast directly without sanitization

## Remediation
- Validate `Origin` header against an allowlist
- Require valid session/JWT before accepting WebSocket connection
- Sanitize all messages before broadcasting
