---
name: Weak JWT Algorithm
version: 1.0.0
description: Detects JWT usage with weak or insecure signing algorithms including HS256 with short secrets, RS256 key confusion, and algorithm downgrade.
tags: [crypto, jwt, algorithm, owasp-a02]
languages: [javascript, typescript, python, java, go]
severity: high
confidence: high
cwe: [CWE-327, CWE-347]
owasp: [A02:2025]
---

# Weak JWT Algorithm

## Overview
JWT algorithm issues include:
1. **HS256 with short key**: HS256 keys shorter than 256 bits are brute-forceable
2. **Algorithm confusion (RS256 → HS256)**: Public key used as HMAC secret when server accepts both
3. **Explicit `none` algorithm**: No signature required
4. **Embedded JWK in header**: Attacker provides their own public key

## Remediation
- Use RS256 or ES256 for production systems (asymmetric keys)
- If using HS256, use a random 32+ byte secret
- Restrict accepted algorithms explicitly: `algorithms: ['RS256']`
- Never allow `none` algorithm
