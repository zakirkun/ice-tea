---
confidence: high
cwe:
    - CWE-916
    - CWE-760
description: Detects passwords used directly as encryption keys or hashed with fast algorithms instead of proper KDFs (PBKDF2, bcrypt, Argon2, scrypt).
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
name: Insecure Key Derivation
owasp:
    - A02:2025
severity: high
tags:
    - crypto
    - kdf
    - pbkdf2
    - argon2
    - owasp-a02
version: 1.0.0
---

# Insecure Key Derivation

## Overview
Passwords must not be used directly as cryptographic keys or passed through fast hash functions. Proper Key Derivation Functions (KDFs) intentionally add computational cost and mandatory salting:
- **PBKDF2**: NIST-approved, configurable iteration count (minimum 600,000 for SHA-256)
- **bcrypt**: Adaptive work factor, built-in salt
- **scrypt**: Memory-hard, resistant to GPU cracking
- **Argon2id**: Winner of Password Hashing Competition, recommended

## Remediation
Always derive keys from passwords using a KDF with a random salt.

**Vulnerable:**
```python
key = hashlib.sha256(password.encode()).digest()
cipher = AES.new(key, AES.MODE_GCM)
```

**Safe:**
```python
import os
from cryptography.hazmat.primitives.kdf.pbkdf2 import PBKDF2HMAC
from cryptography.hazmat.primitives import hashes
salt = os.urandom(16)
kdf = PBKDF2HMAC(algorithm=hashes.SHA256(), length=32, salt=salt, iterations=600000)
key = kdf.derive(password.encode())
```
