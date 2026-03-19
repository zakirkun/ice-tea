---
confidence: medium
cwe:
    - CWE-78
description: Detects untrusted input passed directly to operating system shell commands.
languages:
    - generic
    - kotlin
    - dart
    - zig
    - elixir
name: Command Injection
owasp:
    - A03:2025
severity: critical
tags:
    - cmdi
    - rce
    - injection
    - os
version: 1.0.0
---

# Command Injection

## Overview
Detects untrusted input passed directly to operating system shell commands.

## Remediation
Avoid calling OS commands directly. Use built-in language APIs. If necessary, use exec arrays instead of shell strings.
