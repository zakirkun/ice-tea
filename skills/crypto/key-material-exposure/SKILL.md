---
confidence: high
cwe:
    - CWE-321
description: Detects private keys, certificates, and cryptographic key material hardcoded or committed in source code.
languages:
    - generic
    - javascript
    - typescript
    - python
    - go
    - java
    - php
    - kotlin
    - dart
    - zig
    - elixir
name: Cryptographic Key Material Exposure
owasp:
    - A07:2025
severity: critical
tags:
    - crypto
    - key-material
    - secrets
    - owasp-a07
version: 1.0.0
---

# Cryptographic Key Material Exposure

## Overview
Private keys and key material committed to version control or hardcoded in source permanently compromise:
- TLS certificates (can intercept all encrypted traffic)
- SSH keys (unauthorized server access)
- JWT signing keys (forge authentication tokens)
- PGP keys (decrypt private communications)

## Remediation
- Use `.gitignore` to exclude all key files
- Rotate any exposed keys immediately — treat them as compromised
- Use secrets management systems (Vault, AWS Secrets Manager)
- Use hardware security modules (HSM) for production key storage
