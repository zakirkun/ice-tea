---
confidence: high
cwe:
    - CWE-93
description: Detects user-controlled input in SMTP commands that can inject malicious SMTP commands via CRLF sequences.
languages:
    - python
    - javascript
    - typescript
    - java
    - php
    - go
    - kotlin
    - dart
    - zig
    - elixir
name: SMTP Command Injection
owasp:
    - A03:2025
severity: high
tags:
    - network
    - smtp
    - injection
    - owasp-a03
version: 1.0.0
---

# SMTP Command Injection

## Overview
SMTP command injection (distinct from email header injection) occurs when user input is placed directly in raw SMTP protocol commands. By injecting CRLF sequences, attackers can:
- Add additional RCPT TO commands (spam relay)
- Send emails to unintended recipients
- Inject arbitrary SMTP commands (VRFY, EXPN, DATA)

## Remediation
- Use email library abstractions that handle SMTP protocol internally
- Strip CRLF characters from all email fields before SMTP transmission
- Validate sender/recipient addresses with strict RFC-compliant regex
