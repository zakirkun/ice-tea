---
confidence: high
cwe:
    - CWE-329
    - CWE-330
description: Detects static or hardcoded Initialization Vectors (IV) or nonces in symmetric encryption, breaking confidentiality.
languages:
    - javascript
    - typescript
    - python
    - go
    - java
    - php
    - ruby
    - kotlin
    - dart
    - zig
    - elixir
name: Hardcoded IV / Nonce
owasp:
    - A02:2025
severity: high
tags:
    - crypto
    - iv
    - nonce
    - aes
    - owasp-a02
version: 1.0.0
---

# Hardcoded IV / Nonce

## Overview
A static or predictable Initialization Vector (IV) in block cipher modes (CBC, CTR, GCM) critically weakens encryption:
- **CBC with fixed IV**: Two identical plaintexts produce identical ciphertexts, leaking information
- **CTR/GCM with reused nonce**: Nonce reuse in CTR/GCM allows key recovery and plaintext XOR

The IV/nonce must be **randomly generated** for each encryption operation and stored alongside the ciphertext.

## Detection Strategy
- Zero-filled IV: `\x00\x00...\x00` or `iv = bytes(16)`
- Hardcoded IV bytes: `iv = b'\x01\x02\x03...'`
- Derived from static string: `iv = "hardcoded_iv____".encode()`

## Remediation
Generate a cryptographically random IV for every encryption operation.

**Vulnerable (Python):**
```python
from Crypto.Cipher import AES
iv = b'\x00' * 16  # Static zero IV!
cipher = AES.new(key, AES.MODE_CBC, iv)
```

**Safe (Python):**
```python
import os
from Crypto.Cipher import AES
iv = os.urandom(16)  # Random IV
cipher = AES.new(key, AES.MODE_CBC, iv)
ciphertext = iv + cipher.encrypt(plaintext)  # Prepend IV to output
```
