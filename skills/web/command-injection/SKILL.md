---
name: Command Injection
version: 1.0.0
description: Detects untrusted input passed directly to operating system shell commands.
tags: [cmdi, rce, injection, os]
languages: [generic]
severity: critical
confidence: medium
cwe: [CWE-78]
owasp: [A03:2025]
---

# Command Injection

## Overview
Detects untrusted input passed directly to operating system shell commands.

## Remediation
Avoid calling OS commands directly. Use built-in language APIs. If necessary, use exec arrays instead of shell strings.
