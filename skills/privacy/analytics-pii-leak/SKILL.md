---
confidence: medium
cwe:
    - CWE-359
description: Detects PII being sent to analytics platforms like Google Analytics, Mixpanel, or Amplitude in event properties.
languages:
    - javascript
    - typescript
    - kotlin
    - dart
    - zig
    - elixir
name: PII in Analytics Events
owasp:
    - A09:2025
severity: high
tags:
    - privacy
    - analytics
    - pii
    - owasp-a09
version: 1.0.0
---

# PII in Analytics Events

## Overview
Sending PII to analytics platforms violates GDPR (data minimization principle) and the Terms of Service of most analytics vendors (e.g., Google Analytics prohibits sending PII). Common violations:
- Email addresses in user identification
- Full names in event properties
- Health conditions in custom dimensions

## Remediation
- Use anonymized or pseudonymized identifiers in analytics
- Hash email addresses before sending (but this is still PII under GDPR)
- Audit all analytics events for PII before going live
