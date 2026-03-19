---
confidence: low
cwe:
    - CWE-778
description: Detects authentication and authorization events that are not logged, hampering incident detection and forensics.
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
name: Insufficient Security Event Logging
owasp:
    - A09:2025
severity: medium
tags:
    - logging
    - audit
    - monitoring
    - owasp-a09
version: 1.0.0
---

# Insufficient Security Event Logging

## Overview
Security events that must be logged for compliance and incident response:
- Failed login attempts (with username and IP)
- Successful logins and logouts
- Password changes and resets
- Admin operations (user creation, deletion, privilege changes)
- Access control failures (403 responses)

Without adequate logging, attackers can conduct long-running attacks undetected, and post-incident forensics is impossible.

## Detection Strategy
- Authentication handlers that don't call any logger on failure
- Admin operations without audit log entries

## Remediation
Log all security-relevant events with: timestamp, user ID, IP address, action performed, success/failure.
