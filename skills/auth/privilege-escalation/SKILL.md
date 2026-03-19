---
confidence: medium
cwe:
    - CWE-269
description: Detects code patterns that allow users to elevate their own privileges through API manipulation or mass assignment.
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
name: Privilege Escalation Through Code Logic
owasp:
    - A01:2025
severity: critical
tags:
    - auth
    - privilege-escalation
    - owasp-a01
version: 1.0.0
---

# Privilege Escalation Through Code Logic

## Overview
Privilege escalation occurs when users can elevate their own access levels through:
- Mass assignment: User profile update endpoint accepts `role` or `isAdmin` fields
- Parameter tampering: Changing `user_id` to another user's ID
- API parameter forgery: Sending `{"role": "admin"}` in user update request

## Remediation
- Explicitly whitelist allowed fields in user update operations
- Never accept role, permission, or admin fields from user-controlled input
- Verify that any elevation action is performed by a sufficiently privileged user
