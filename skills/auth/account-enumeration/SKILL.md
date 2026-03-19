---
confidence: high
cwe:
    - CWE-204
description: Detects authentication and registration endpoints that reveal whether a username or email exists, enabling targeted attacks.
languages:
    - javascript
    - typescript
    - python
    - go
    - java
    - php
    - kotlin
    - dart
    - zig
    - elixir
name: Account Enumeration
owasp:
    - A07:2025
severity: medium
tags:
    - auth
    - enumeration
    - information-disclosure
    - owasp-a07
version: 1.0.0
---

# Account Enumeration

## Overview
Applications that return different responses for valid vs invalid usernames allow attackers to enumerate valid accounts for:
- Targeted phishing campaigns
- Credential stuffing (knowing which accounts exist)
- Social engineering targeting known employees

## Detection Strategy
Look for error messages that distinguish "wrong password" from "user not found".

## Remediation
- Return identical error messages regardless of whether username exists
- Use constant-time comparison to prevent timing-based enumeration
- Implement CAPTCHA after multiple failed attempts
