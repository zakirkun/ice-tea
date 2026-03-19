---
confidence: high
cwe:
    - CWE-327
    - CWE-328
description: Detects use of weak or deprecated cryptographic algorithms
languages:
    - go
    - kotlin
    - dart
    - zig
    - elixir
name: Weak Cryptography Detection
owasp:
    - A04:2025
severity: high
tags:
    - crypto
    - hash
    - weak
    - owasp-a04
version: 1.0.0
---

# Weak Cryptography Detection

## Overview
Use of weak cryptographic algorithms (MD5, SHA1, DES) can lead to security vulnerabilities. These algorithms are considered broken for security purposes.

## Remediation
Use strong algorithms: SHA-256+, AES-256, bcrypt/scrypt/argon2 for passwords.
