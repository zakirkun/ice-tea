---
name: Weak Cipher Algorithm Usage
version: 1.0.0
description: Detects use of broken or weak symmetric encryption algorithms (DES, 3DES, RC4, Blowfish, ECB mode).
tags: [crypto, cipher, des, rc4, owasp-a02]
languages: [javascript, typescript, python, go, java, php, ruby]
severity: high
confidence: high
cwe: [CWE-327]
owasp: [A02:2025]
---

# Weak Cipher Algorithm Usage

## Overview
Several symmetric encryption algorithms are cryptographically broken and should not be used:
- **DES**: 56-bit key — broken since 1998, trivially brute-forced
- **3DES/TDEA**: 112-bit effective security, SWEET32 attack, deprecated since 2018
- **RC4**: Statistical biases, NOMORE attack, prohibited in TLS (RFC 7465)
- **ECB mode**: Each block encrypted independently — patterns visible in ciphertext
- **Blowfish**: 64-bit block size — vulnerable to SWEET32 with large data volumes

## Remediation
Use **AES-256-GCM** (authenticated encryption) for most use cases.

**Vulnerable:**
```python
from Crypto.Cipher import DES
cipher = DES.new(key, DES.MODE_ECB)
```

**Safe:**
```python
from Crypto.Cipher import AES
import os
key = os.urandom(32)  # AES-256
nonce = os.urandom(12)
cipher = AES.new(key, AES.MODE_GCM, nonce=nonce)
```
