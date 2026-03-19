---
confidence: high
cwe:
    - CWE-330
description: Detects initialization vectors generated using non-cryptographic random functions, compromising cipher security.
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
name: Insecure Random IV Generation
owasp:
    - A02:2025
severity: high
tags:
    - crypto
    - iv
    - random
    - owasp-a02
version: 1.0.0
---

# Insecure Random IV Generation

## Overview
Even when an IV is generated dynamically (not hardcoded), using a non-cryptographic PRNG makes it predictable:
- `Math.random()` in JavaScript (not cryptographic)
- `random.random()` in Python (Mersenne Twister, seeded predictably)
- `rand()` in C (LCG, predictable)

A predictable IV defeats the purpose of encryption for modes that rely on IV uniqueness.

## Remediation
Use OS-provided cryptographic random for IV generation:
- Python: `os.urandom(16)`
- Node.js: `crypto.randomBytes(16)`
- Go: `io.ReadFull(rand.Reader, iv)`
- Java: `new SecureRandom().nextBytes(iv)`
