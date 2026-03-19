---
confidence: medium
cwe:
    - CWE-347
    - CWE-287
description: Detects insecure JWT implementations, such as accepting 'none' algorithms or using hardcoded secrets.
languages:
    - generic
    - kotlin
    - dart
    - zig
    - elixir
name: JSON Web Token (JWT) Weaknesses
owasp:
    - A07:2025
severity: critical
tags:
    - jwt
    - auth
    - crypto
    - owasp-a07
version: 1.0.0
---

# JSON Web Token (JWT) Weaknesses

## Overview
Detects insecure JWT implementations, such as accepting 'none' algorithms or using hardcoded secrets.

## Remediation
Always enforce cryptographic signatures (e.g., HS256, RS256). Do not allow the 'none' algorithm. Always load secrets from environment variables.
