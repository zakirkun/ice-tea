---
name: Insecure RSA Configuration
version: 1.0.0
description: Detects RSA usage with insufficient key sizes, PKCS#1 v1.5 padding (vulnerable to padding oracle), or direct message encryption without hybrid scheme.
tags: [crypto, rsa, oaep, owasp-a02]
languages: [python, javascript, typescript, java, go, php]
severity: high
confidence: high
cwe: [CWE-780, CWE-326]
owasp: [A02:2025]
---

# Insecure RSA Configuration

## Overview
Common RSA vulnerabilities:
1. **Key size < 2048 bits**: Keys below 2048 bits are factorable; minimum is 3072 bits by 2030 (NIST guidance)
2. **PKCS#1 v1.5 padding**: Vulnerable to Bleichenbacher's padding oracle attack; OAEP must be used
3. **Direct message encryption**: RSA should only encrypt symmetric keys (hybrid encryption), not arbitrary messages
4. **Public exponent e=1 or e=3**: Trivial to break with small exponents

## Remediation
- Use RSA key size ≥ 2048 bits (prefer 4096 for long-lived keys)
- Always use OAEP padding (`PKCS1_OAEP` in Python, `RSA/ECB/OAEPWithSHA-256AndMGF1Padding` in Java)
- Use hybrid encryption (encrypt data with AES, encrypt AES key with RSA)

**Vulnerable (Python):**
```python
from Crypto.PublicKey import RSA
from Crypto.Cipher import PKCS1_v1_5  # Vulnerable padding!
key = RSA.generate(1024)  # Too small!
```

**Safe (Python):**
```python
from Crypto.PublicKey import RSA
from Crypto.Cipher import PKCS1_OAEP
key = RSA.generate(4096)
cipher = PKCS1_OAEP.new(key.publickey())
```
