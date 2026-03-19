---
name: Concurrent Session Issues
version: 1.0.0
description: Detects missing concurrent session controls that allow multiple active sessions or fail to invalidate old sessions on new login.
tags: [auth, session, concurrency, owasp-a07]
languages: [javascript, typescript, python, php, java, go]
severity: medium
confidence: low
cwe: [CWE-613]
owasp: [A07:2025]
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
