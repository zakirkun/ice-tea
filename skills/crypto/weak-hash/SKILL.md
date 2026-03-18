---
name: Weak Cryptography Detection
version: 1.0.0
description: Detects use of weak or deprecated cryptographic algorithms
tags: [crypto, hash, weak, owasp-a04]
languages: [go]
severity: high
confidence: high
cwe: [CWE-327, CWE-328]
owasp: [A04:2025]
---

# Weak Cryptography Detection

## Overview
Use of weak cryptographic algorithms (MD5, SHA1, DES) can lead to security vulnerabilities. These algorithms are considered broken for security purposes.

## Remediation
Use strong algorithms: SHA-256+, AES-256, bcrypt/scrypt/argon2 for passwords.
