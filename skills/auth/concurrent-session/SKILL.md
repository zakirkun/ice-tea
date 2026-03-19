---
confidence: low
cwe:
    - CWE-613
description: Detects missing concurrent session controls that allow multiple active sessions or fail to invalidate old sessions on new login.
languages:
    - javascript
    - typescript
    - python
    - php
    - java
    - go
    - kotlin
    - dart
    - zig
    - elixir
name: Concurrent Session Issues
owasp:
    - A07:2025
severity: medium
tags:
    - auth
    - session
    - concurrency
    - owasp-a07
version: 1.0.0
---

# Concurrent Session Issues

## Overview
Applications without concurrent session controls allow:
- Multiple simultaneous logins from different locations (credential sharing detection fails)
- Old sessions remaining valid after logout from one device
- Session token theft goes undetected because both sessions appear valid

## Remediation
- Implement concurrent session limit (configurable per user tier)
- On new login, optionally invalidate all other sessions
- Provide users with a "log out all other sessions" feature
- Track sessions per user in the database
