---
confidence: medium
cwe:
    - CWE-359
description: Detects health and medical information logged, transmitted without encryption, or insufficiently protected, violating HIPAA and similar regulations.
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
name: Health / Medical Data Exposure
owasp:
    - A02:2025
severity: critical
tags:
    - privacy
    - hipaa
    - health-data
    - owasp-a02
version: 1.0.0
---

# Health / Medical Data Exposure

## Overview
Protected Health Information (PHI) under HIPAA includes diagnoses, medications, treatment history, and health identifiers. Improper handling creates legal liability and severe breach consequences.

## Remediation
- Encrypt PHI at rest and in transit
- Implement audit logging for all PHI access
- Apply minimum necessary access principles
- Never log PHI without encryption and access controls
