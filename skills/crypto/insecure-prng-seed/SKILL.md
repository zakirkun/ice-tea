---
name: Insecure PRNG Seed
version: 1.0.0
description: Detects cryptographic operations where pseudo-random number generators are seeded with predictable values.
tags: [crypto, prng, random, owasp-a02]
languages: [python, javascript, typescript, java, go, php, c, cpp]
severity: high
confidence: high
cwe: [CWE-335, CWE-330]
owasp: [A02:2025]
---

# Insecure PRNG Seed

## Overview
A PRNG seeded with predictable values produces predictable output. Attackers who know or can guess the seed can predict all subsequent random values, including:
- Session tokens and CSRFs
- Password reset tokens
- Cryptographic keys
- Nonces

Common bad seeds: `time()`, `getpid()`, hardcoded integers, zero.

## Remediation
Use OS entropy sources for seeding or use CSPRNGs directly:
- Python: `secrets` module, `os.urandom()`
- Go: `crypto/rand`
- Node.js: `crypto.randomBytes()`
- Java: `SecureRandom()` (default seeding is safe)
- C/C++: `/dev/urandom` or `getrandom()`
