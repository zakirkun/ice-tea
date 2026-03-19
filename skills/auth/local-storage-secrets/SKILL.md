---
confidence: medium
cwe:
    - CWE-312
description: Detects potential storage of sensitive credentials or JWTs directly in the browser's localStorage or sessionStorage.
languages:
    - generic
    - kotlin
    - dart
    - zig
    - elixir
name: Sensitive Data in LocalStorage
owasp:
    - A02:2025
severity: medium
tags:
    - frontend
    - storage
    - secrets
    - owasp-a02
version: 1.0.0
---

# Sensitive Data in LocalStorage

## Overview
Detects potential storage of sensitive credentials or JWTs directly in the browser's localStorage or sessionStorage.

## Remediation
Use HttpOnly/Secure cookies for session tokens. Local storage is easily accessible via any XSS vulnerability.
