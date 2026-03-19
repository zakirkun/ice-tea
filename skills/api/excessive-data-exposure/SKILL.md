---
confidence: medium
cwe:
    - CWE-200
description: Detects API endpoints that return more data than required, including sensitive fields that clients should not receive.
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
name: Excessive Data Exposure in API Responses
owasp:
    - A01:2025
severity: high
tags:
    - api
    - information-disclosure
    - owasp-api3
version: 1.0.0
---

# Excessive Data Exposure

## Overview
APIs that return complete database objects rely on the frontend to filter sensitive fields. This is a common vulnerability (OWASP API3) where:
- Password hashes returned in user objects
- Internal IDs, audit timestamps exposed
- Admin-only fields visible to all users
- Credit card details partially exposed

## Remediation
- Use Data Transfer Objects (DTOs) that explicitly define what fields to return
- Never return entire database models directly
- Use field-level serialization control (`@JsonIgnore`, `exclude`, `select`)
