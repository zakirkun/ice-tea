---
confidence: medium
cwe:
    - CWE-798
description: Detects API keys, passwords, and tokens embedded directly in the source code.
languages:
    - generic
    - kotlin
    - dart
    - zig
    - elixir
name: Hardcoded Secrets
owasp:
    - A07:2025
severity: critical
tags:
    - secrets
    - auth
    - hardcoded
    - owasp-a07
version: 1.0.0
---

# Hardcoded Secrets

## Overview
Detects API keys, passwords, and tokens embedded directly in the source code.

## Remediation
Use a secure secrets manager or environment variables to inject sensitive credentials at runtime.
