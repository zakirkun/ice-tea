---
name: Insecure TLS Configuration
version: 1.0.0
description: Detects insecure TLS settings including disabled certificate verification, outdated protocol versions, and weak cipher suites.
tags: [tls, ssl, crypto, owasp-a02]
languages: [go, python, javascript, typescript, java, php]
severity: critical
confidence: high
cwe: [CWE-295, CWE-326, CWE-327]
owasp: [A02:2025]
---

# Insecure TLS Configuration

## Overview
TLS misconfigurations expose communications to interception and tampering:
1. **InsecureSkipVerify**: Disables certificate validation entirely — trivial MITM
2. **TLS 1.0/1.1**: Deprecated protocols with known weaknesses (BEAST, POODLE)
3. **Weak cipher suites**: RC4, DES, 3DES, EXPORT ciphers
4. **Self-signed cert acceptance**: `verify=False` in Python requests

## Remediation
- Never set `InsecureSkipVerify: true` in production
- Enforce TLS 1.2 minimum (TLS 1.3 preferred)
- Use only AEAD cipher suites (AES-GCM, ChaCha20-Poly1305)
- Always verify certificates in HTTP clients

**Vulnerable (Go):**
```go
tr := &http.Transport{
    TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
}
```

**Safe (Go):**
```go
tr := &http.Transport{
    TLSClientConfig: &tls.Config{
        MinVersion: tls.VersionTLS13,
    },
}
```
