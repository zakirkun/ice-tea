---
confidence: low
cwe:
    - CWE-522
description: Detects API key implementations without expiration dates or rotation mechanisms, creating long-lived credentials.
languages:
    - javascript
    - typescript
    - python
    - go
    - java
    - php
    - kotlin
    - dart
    - zig
    - elixir
name: API Key Without Rotation or Expiration
owasp:
    - A07:2025
severity: medium
tags:
    - api
    - api-key
    - rotation
    - owasp-a07
version: 1.0.0
---

# API Key Without Rotation

## Overview
Static API keys without expiration or rotation remain valid indefinitely after compromise. Best practices require:
- Short-lived keys with automatic expiration
- Easy rotation mechanism
- Key usage monitoring and anomaly detection
- Immediate revocation capability

## Remediation
- Add expiration dates to all API keys
- Implement key rotation workflows
- Notify users before expiration
- Monitor for unusual usage patterns
