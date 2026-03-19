---
name: iOS URL Scheme Hijacking
version: 1.0.0
description: Detects iOS custom URL scheme handling that processes sensitive data without sender verification.
tags: [ios, mobile, url-scheme, owasp-m1]
languages: [generic, java]
severity: high
confidence: medium
cwe: [CWE-940]
owasp: [A01:2025]
---

# iOS URL Scheme Hijacking

## Overview
Custom URL schemes (e.g., `myapp://`) can be registered by any app on the device. A malicious app registering the same scheme can intercept deep links, OAuth callbacks, and payment confirmations.

## Remediation
- Use Universal Links (HTTPS-based) instead of custom URL schemes for sensitive flows
- Validate the `sourceApplication` parameter in URL handler
- Never pass sensitive tokens via URL scheme parameters
- For OAuth, use ASWebAuthenticationSession with system-managed callback
