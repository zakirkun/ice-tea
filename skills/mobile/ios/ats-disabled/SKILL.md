---
confidence: high
cwe:
    - CWE-319
description: Detects iOS ATS exceptions that allow insecure HTTP connections or disable certificate validation.
languages:
    - generic
    - kotlin
    - dart
    - zig
    - elixir
name: iOS App Transport Security (ATS) Disabled
owasp:
    - A02:2025
severity: high
tags:
    - ios
    - mobile
    - ats
    - tls
    - owasp-m3
version: 1.0.0
---

# iOS App Transport Security Disabled

## Overview
iOS App Transport Security (ATS) enforces HTTPS and modern TLS for all network connections. Disabling ATS via `NSAppTransportSecurity` plist exceptions allows insecure HTTP connections and weakened TLS, enabling MITM attacks.

## Remediation
- Remove `NSAllowsArbitraryLoads: true`
- Fix server TLS configuration to be ATS-compliant
- Use specific domain exceptions only for legacy systems during migration
