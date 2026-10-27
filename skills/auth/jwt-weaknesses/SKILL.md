---
confidence: medium
cwe:
    - CWE-347
    - CWE-287
description: "Detects insecure JWT implementations including none-algorithm acceptance, hardcoded signing secrets, and disabled verification. Use when reviewing authentication code, token signing, or JWT validation logic."
languages:
    - generic
    - kotlin
    - dart
    - zig
    - elixir
name: json-web-token-jwt-weaknesses
owasp:
    - A07:2025
severity: critical
tags:
    - jwt
    - auth
    - crypto
    - owasp-a07
version: 1.0.0
---

# JSON Web Token (JWT) Weaknesses

Scans for insecure JWT patterns across common libraries (PyJWT, jsonwebtoken, java-jwt, go-jwt). See `patterns.yaml` for detection rules.

## Detection Workflow

1. **Find JWT usage** — search for `jwt.sign`, `jwt.verify`, `jwt.decode`, `Jwts.parser`
2. **Check algorithm** — flag `algorithms: ['none']` or `SigningMethodNone`
3. **Check secrets** — flag hardcoded string literals passed as signing keys
4. **Check verification** — flag `verify_signature=False` or missing verification

### Vulnerable

```javascript
jwt.verify(token, 'hardcoded-secret');
jwt.decode(token, { algorithms: ['none'] });
```

```python
jwt.decode(token, options={"verify_signature": False})
```

### Safe

```javascript
jwt.verify(token, process.env.JWT_SECRET, { algorithms: ['RS256'] });
```
