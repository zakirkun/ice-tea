---
name: Timing Attack Vulnerability
version: 1.0.0
description: Detects comparison operations on secrets using non-constant-time functions, enabling timing side-channel attacks.
tags: [crypto, timing-attack, side-channel, owasp-a02]
languages: [python, javascript, typescript, go, java, php, ruby]
severity: high
confidence: high
cwe: [CWE-208]
owasp: [A02:2025]
---

# Timing Attack Vulnerability

## Overview
String comparison functions (`==`, `===`, `strcmp`) short-circuit on the first mismatched character. Measuring response times reveals how many characters of a secret match the attacker's guess. This is exploitable for:
- API key validation
- HMAC signature verification
- Session token comparison
- Password comparison (before hashing)

## Remediation
Use constant-time comparison functions:
- Python: `hmac.compare_digest()`
- Go: `subtle.ConstantTimeCompare()`
- Node.js: `crypto.timingSafeEqual()`
- PHP: `hash_equals()`
- Java: `MessageDigest.isEqual()`
