---
confidence: high
cwe:
    - CWE-78
description: Detects command injection vulnerabilities where user input is passed to OS command execution
languages:
    - go
    - python
    - javascript
    - php
    - kotlin
    - dart
    - zig
    - elixir
name: Command Injection Detection
owasp:
    - A05:2025
severity: critical
tags:
    - injection
    - command
    - os
    - owasp-a05
version: 1.0.0
---

# Command Injection Detection

## Overview
Command injection occurs when an application passes user-supplied data to a system shell or OS command without proper sanitization.

## Remediation
Avoid using shell commands with user input. If necessary, use allowlists and never pass raw user input to command execution functions.
