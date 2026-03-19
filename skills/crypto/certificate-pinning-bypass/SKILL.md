---
confidence: high
cwe:
    - CWE-295
description: Detects network configurations and code patterns that disable or bypass certificate pinning, weakening TLS security.
languages:
    - java
    - javascript
    - typescript
    - go
    - python
    - kotlin
    - dart
    - zig
    - elixir
name: Certificate Pinning Bypass
owasp:
    - A02:2025
severity: critical
tags:
    - crypto
    - certificate-pinning
    - tls
    - owasp-a02
version: 1.0.0
---

# Certificate Pinning Bypass

## Overview
Certificate pinning enhances TLS by restricting which certificates are trusted for specific connections. Bypassing or disabling pinning allows MITM attacks even on apps that previously validated certificates.

Common bypass methods:
- Setting `checkValidity = false` in pinning configuration
- Using a catch-all TrustManager
- OkHttp pinning disabled
- Node.js `NODE_TLS_REJECT_UNAUTHORIZED=0`

## Remediation
- Implement pinning using the native OS keystore or a dedicated library
- Monitor for pinning bypass attempts
- Test pinning resistance against tools like Frida, Objection
