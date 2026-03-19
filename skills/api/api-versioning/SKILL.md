---
confidence: medium
cwe:
    - CWE-1059
    - CWE-1188
description: Detects old API versions that may lack current security controls and deprecated endpoints still accessible in production.
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
name: Deprecated / Unprotected API Versions
owasp:
    - A09:2025
severity: medium
tags:
    - api-versioning
    - api
    - owasp-api9
version: 1.0.0
---

# Deprecated / Unprotected API Versions

## Overview
Old API versions are a common attack vector because:
- They often lack security controls added in newer versions
- Authentication requirements may differ between v1 and v2
- They may expose endpoints removed from newer versions for security reasons
- Rate limiting and input validation may not apply to legacy versions

## Detection Strategy
- Old API version routes (`/v1/`, `/v0/`) served alongside newer versions
- Deprecated endpoint paths still registered in routing
- Different authentication middleware applied to different version prefixes

## Remediation
- Maintain a clear API deprecation policy with sunset dates
- Apply the same security controls to all active API versions
- Log and monitor usage of deprecated endpoints
- Return HTTP 410 (Gone) for fully deprecated versions
