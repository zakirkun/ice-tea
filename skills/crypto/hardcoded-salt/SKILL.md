---
confidence: high
cwe:
    - CWE-760
description: Detects hardcoded or static salt values used in password hashing, allowing precomputed rainbow table attacks.
languages:
    - python
    - javascript
    - typescript
    - java
    - go
    - php
    - kotlin
    - dart
    - zig
    - elixir
name: Hardcoded Salt in Password Hashing
owasp:
    - A02:2025
severity: high
tags:
    - crypto
    - salt
    - password-hashing
    - owasp-a02
version: 1.0.0
---

# Hardcoded Salt

## Overview
A salt's purpose is to make each password hash unique, even when two users have identical passwords. A hardcoded static salt negates this purpose because:
- All users with the same password have the same hash
- Precomputed tables (rainbow tables) can be built for that specific salt
- A database breach exposes all passwords simultaneously

## Remediation
Generate a unique random salt for each user:
```python
import os
import hashlib
salt = os.urandom(32)  # Random 32-byte salt per user
hash = hashlib.pbkdf2_hmac('sha256', password.encode(), salt, 600000)
# Store both salt and hash per user
```
