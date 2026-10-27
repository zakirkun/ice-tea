---
confidence: high
cwe:
    - CWE-327
    - CWE-328
description: "Detects usage of weak or deprecated cryptographic algorithms (MD5, SHA1, DES) and recommends modern replacements. Use when auditing code for insecure hashing, weak encryption, or crypto vulnerabilities."
languages:
    - go
    - kotlin
    - dart
    - zig
    - elixir
name: weak-cryptography-detection
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

Identifies imports and function calls for deprecated crypto algorithms and suggests modern replacements. See `patterns.yaml` for detection rules.

## Detection Workflow

1. **Scan imports** — flag `crypto/md5`, `crypto/sha1`, `crypto/des`, `hashlib.md5`, `hashlib.sha1`
2. **Scan calls** — flag `md5.New()`, `md5.Sum()`, `sha1.New()`, `sha1.Sum()`, `DES.new()`
3. **Classify** — hashing (MD5/SHA1) vs encryption (DES/RC4) vs password storage
4. **Recommend** — SHA-256+ for hashing, AES-256 for encryption, bcrypt/scrypt/argon2 for passwords

### Vulnerable

```go
import "crypto/md5"
hash := md5.Sum(data)
```

### Safe

```go
import "crypto/sha256"
hash := sha256.Sum256(data)
```
