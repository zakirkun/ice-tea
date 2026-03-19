---
confidence: medium
cwe:
    - CWE-532
description: Detects Personally Identifiable Information (PII) such as email addresses, phone numbers, and SSNs being written to log files.
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
name: PII (Personally Identifiable Information) in Logs
owasp:
    - A09:2025
severity: high
tags:
    - privacy
    - gdpr
    - pii
    - logging
    - owasp-a09
version: 1.0.0
---

# PII in Logs

## Overview
Logging PII (email addresses, phone numbers, names, SSNs, credit card numbers) creates compliance risks under GDPR, CCPA, HIPAA, and PCI-DSS. Log files:
- May be stored indefinitely beyond data retention requirements
- Are often shipped to third-party aggregators (Datadog, Splunk, ELK)
- May have broader access than production databases
- Can be subpoenaed in legal proceedings

## Remediation
- Implement a log sanitizer/redactor for all PII fields
- Use structured logging and explicitly list safe fields
- Mask PII in logs: `user@ex*****.com`, `+1-xxx-xxx-1234`
- Audit log destinations for GDPR data transfer compliance
