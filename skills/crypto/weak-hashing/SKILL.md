---
confidence: medium
cwe:
    - CWE-327
    - CWE-328
description: Detects the use of cryptographically weak hashing algorithms like MD5 and SHA1.
languages:
    - generic
    - kotlin
    - dart
    - zig
    - elixir
name: Weak Hashing Algorithms
owasp:
    - A02:2025
severity: high
tags:
    - crypto
    - hash
    - md5
    - sha1
    - owasp-a02
version: 1.0.0
---

# Weak Hashing Algorithms

## Overview
Detects the use of cryptographically weak hashing algorithms like MD5 and SHA1.

## Remediation
Use strong hashing algorithms like SHA-256, SHA-3, or argon2/bcrypt for passwords.
