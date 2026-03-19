---
name: Android Insecure Certificate Validation
version: 1.0.0
description: Detects Android TrustManager implementations that accept all certificates, disabling SSL/TLS security.
tags: [android, mobile, certificate-pinning, ssl, owasp-m3]
languages: [java]
severity: critical
confidence: high
cwe: [CWE-295]
owasp: [A02:2025]
---

# Android Insecure Certificate Validation

## Overview
Android applications that implement custom `X509TrustManager` with empty or no-op validation methods accept any certificate, including self-signed and attacker-controlled certificates. This completely negates TLS security and enables trivial man-in-the-middle attacks.

Common vulnerable pattern: `checkServerTrusted()` method with empty body.

## Detection Strategy
- `X509TrustManager` with empty `checkServerTrusted()`
- `HostnameVerifier` that always returns true
- `setSSLSocketFactory` with custom factory that accepts all certs

## Remediation
- Use the default system TrustManager (trusts only CA-signed certs)
- Implement certificate pinning for sensitive applications using OkHttp `CertificatePinner` or Android Network Security Config
- Use the `networkSecurityConfig` XML for declarative pinning
