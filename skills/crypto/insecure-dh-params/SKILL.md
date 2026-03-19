---
confidence: high
cwe:
    - CWE-326
description: Detects use of weak Diffie-Hellman parameters (< 2048 bits, export-grade, or known broken groups).
languages:
    - python
    - go
    - java
    - javascript
    - typescript
    - kotlin
    - dart
    - zig
    - elixir
name: Insecure Diffie-Hellman Parameters
owasp:
    - A02:2025
severity: high
tags:
    - crypto
    - dh
    - tls
    - owasp-a02
version: 1.0.0
---

# Insecure Diffie-Hellman Parameters

## Overview
Diffie-Hellman key exchange with weak parameters is vulnerable to precomputation attacks (LogJam attack, 2015). Common issues:
- DH groups < 2048 bits (export-grade: 512/768/1024 bits)
- Using standard well-known small prime groups (precomputed NFS databases)
- Reusing the same DH parameters across many connections (static DH)

## Remediation
- Use DH groups ≥ 2048 bits or RFC 7919 FFDHE groups
- Prefer ECDH (Elliptic Curve DH) with P-256/P-384/X25519 — more efficient and secure
- Generate unique DH parameters per deployment
