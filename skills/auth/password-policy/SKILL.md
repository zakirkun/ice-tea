---
confidence: high
cwe:
    - CWE-521
    - CWE-916
    - CWE-327
description: Detects passwords stored in plaintext or hashed with weak/unsalted algorithms instead of bcrypt/argon2/scrypt.
languages:
    - php
    - python
    - javascript
    - java
    - go
    - ruby
    - kotlin
    - dart
    - zig
    - elixir
name: Insecure Password Storage
owasp:
    - A02:2025
severity: critical
tags:
    - password
    - auth
    - hashing
    - owasp-a02
version: 1.0.0
---

# Insecure Password Storage

## Overview
Storing passwords in plaintext or with weak hashing algorithms like MD5 or SHA-1 (even with a salt) is insecure because:
- Plaintext passwords are immediately exposed in a database breach
- MD5/SHA-1 are too fast — GPU cracking renders them ineffective
- Missing salts allow rainbow table attacks

The correct approach is to use a purpose-built password hashing algorithm: **bcrypt**, **Argon2id**, or **scrypt**.

## Detection Strategy
Look for:
- Direct string comparison of passwords (no hashing at all)
- Use of MD5/SHA1/SHA256 for password storage (wrong tool for the job)
- Missing bcrypt/argon2/scrypt usage in authentication code

## Remediation
Use `bcrypt` with a minimum cost factor of 12, or `Argon2id`.

**Vulnerable (Python):**
```python
import hashlib
stored = hashlib.md5(password.encode()).hexdigest()
```

**Safe (Python):**
```python
import bcrypt
hashed = bcrypt.hashpw(password.encode(), bcrypt.gensalt(rounds=12))
```
