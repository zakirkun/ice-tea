---
confidence: low
cwe:
    - CWE-284
description: Detects multi-step workflows where step validation relies on client-side state or is insufficiently enforced server-side.
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
name: Workflow Step Bypass
owasp:
    - A01:2025
severity: high
tags:
    - business-logic
    - workflow
    - owasp-a01
version: 1.0.0
---

# Workflow Step Bypass

## Overview
Multi-step processes (checkout, registration, password reset, onboarding) where step progression is enforced only by client-side state or easily-manipulated tokens can be bypassed. Attackers jump directly to final steps (e.g., payment confirmation without paying, account activation without email verification).

## Remediation
- Store workflow step state server-side, not in client cookies/form fields
- Verify each step's prerequisites before proceeding
- Use server-side session to track workflow progress
