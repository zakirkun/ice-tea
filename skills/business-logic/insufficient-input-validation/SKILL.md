---
confidence: low
cwe:
    - CWE-20
description: Detects missing or inadequate server-side validation of business-critical fields that could be manipulated.
languages:
    - javascript
    - typescript
    - python
    - php
    - java
    - go
    - kotlin
    - dart
    - zig
    - elixir
name: Insufficient Business Logic Input Validation
owasp:
    - A03:2025
severity: medium
tags:
    - business-logic
    - validation
    - owasp-a03
version: 1.0.0
---

# Insufficient Business Logic Input Validation

## Overview
Beyond injection and type checking, business logic requires semantic validation:
- Age/date fields must be realistic (no future birthdays for adults)
- Email domain validation (company domains for B2B)
- Geolocation consistency (billing address matches IP country)
- Field length limits on free-text fields to prevent DoS

## Remediation
- Implement domain-specific validation for each business field
- Define and enforce explicit validation rules at the controller layer
- Use validation libraries with schema definitions
