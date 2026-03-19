---
confidence: medium
cwe:
    - CWE-798
description: Detects hardcoded passwords, API keys, and tokens in source code
languages:
    - go
    - javascript
    - python
    - java
    - php
    - kotlin
    - dart
    - zig
    - elixir
name: Hardcoded Credentials Detection
owasp:
    - A07:2025
severity: high
tags:
    - credentials
    - secrets
    - hardcoded
    - owasp-a07
version: 1.0.0
---

# Hardcoded Credentials Detection

## Overview
Hardcoded credentials in source code can be extracted by attackers with access to the codebase or compiled binaries.

## Remediation
Use environment variables, secrets managers, or configuration files (excluded from version control) to manage sensitive credentials.
