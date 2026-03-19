---
name: Insufficient Business Logic Input Validation
version: 1.0.0
description: Detects missing or inadequate server-side validation of business-critical fields that could be manipulated.
tags: [business-logic, validation, owasp-a03]
languages: [javascript, typescript, python, php, java, go]
severity: medium
confidence: low
cwe: [CWE-20]
owasp: [A03:2025]
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
