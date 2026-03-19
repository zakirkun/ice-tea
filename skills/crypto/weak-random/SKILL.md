---
confidence: medium
cwe:
    - CWE-338
description: Detects the use of PRNGs (Pseudo-Random Number Generators) that are not cryptographically secure.
languages:
    - generic
    - kotlin
    - dart
    - zig
    - elixir
name: Weak Random Number Generation
owasp:
    - A02:2025
severity: medium
tags:
    - crypto
    - random
    - prng
    - owasp-a02
version: 1.0.0
---

# Weak Random Number Generation

## Overview
Detects the use of PRNGs (Pseudo-Random Number Generators) that are not cryptographically secure.

## Remediation
Use cryptographically secure PRNGs (CSPRNG), like crypto/rand in Go, secrets in Python, or Crypto.getRandomValues in JS.
