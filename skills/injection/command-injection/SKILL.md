---
name: Command Injection Detection
version: 1.0.0
description: Detects command injection vulnerabilities where user input is passed to OS command execution
tags: [injection, command, os, owasp-a05]
languages: [go, python, javascript, php]
severity: critical
confidence: high
cwe: [CWE-78]
owasp: [A05:2025]
---

# Command Injection Detection

## Overview
Command injection occurs when an application passes user-supplied data to a system shell or OS command without proper sanitization.

## Remediation
Avoid using shell commands with user input. If necessary, use allowlists and never pass raw user input to command execution functions.
