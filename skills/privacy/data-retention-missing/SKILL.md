---
confidence: low
cwe:
    - CWE-272
description: Detects data storage without associated TTL, expiry, or cleanup mechanisms, indicating missing data retention policy.
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
name: Missing Data Retention Policy Implementation
owasp:
    - A05:2025
severity: medium
tags:
    - privacy
    - gdpr
    - data-retention
    - owasp-a05
version: 1.0.0
---

# Missing Data Retention Policy Implementation

## Overview
GDPR Article 5(1)(e) requires data to be kept no longer than necessary. Storing user data indefinitely without purging mechanisms violates the storage limitation principle.

Common issues:
- Log files without rotation or TTL
- Database records never deleted
- Caches that accumulate PII indefinitely
- Backups stored without expiration

## Remediation
- Implement automated data purging for records beyond retention period
- Set TTL on Redis/Memcached keys containing personal data
- Configure log rotation with appropriate retention
- Document and enforce data retention policy
