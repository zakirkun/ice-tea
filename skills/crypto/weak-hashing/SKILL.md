---
name: Weak Hashing Algorithms
version: 1.0.0
description: Detects the use of cryptographically weak hashing algorithms like MD5 and SHA1.
tags: [crypto, hash, md5, sha1, owasp-a02]
languages: [generic]
severity: high
confidence: medium
cwe: [CWE-327, CWE-328]
owasp: [A02:2025]
---

# Weak Hashing Algorithms

## Overview
Detects the use of cryptographically weak hashing algorithms like MD5 and SHA1.

## Remediation
Use strong hashing algorithms like SHA-256, SHA-3, or argon2/bcrypt for passwords.
