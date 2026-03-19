---
confidence: medium
cwe:
    - CWE-532
description: Detects passwords, tokens, credit card numbers, and other sensitive data written to log files.
languages:
    - javascript
    - typescript
    - python
    - go
    - java
    - php
    - ruby
    - kotlin
    - dart
    - zig
    - elixir
name: Sensitive Data in Logs
owasp:
    - A09:2025
severity: high
tags:
    - logging
    - sensitive-data
    - pii
    - owasp-a09
version: 1.0.0
---

# Sensitive Data in Logs

## Overview
Logging sensitive information creates secondary exposure risks:
- Log files are often stored less securely than databases
- Log aggregation services (ELK, Splunk) may have broader access
- Logs may be shipped to third parties (monitoring vendors)
- Compliance violations: PCI-DSS (card data), GDPR (PII), HIPAA (health data)

Sensitive data that must not be logged:
- Passwords, PINs, secrets
- Credit card numbers, CVV
- Authentication tokens, session IDs
- Social Security Numbers, health data

## Remediation
- Use structured logging and explicitly list safe fields
- Implement a log sanitizer/redactor middleware
- Never log request bodies wholesale — extract only safe fields
- Use `[REDACTED]` placeholders for sensitive fields
