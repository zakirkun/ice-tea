---
confidence: medium
cwe:
    - CWE-200
description: Detects API endpoints that expose lists without pagination limits, enabling automated data harvesting.
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
name: Mass Enumeration / Data Harvesting
owasp:
    - A01:2025
severity: high
tags:
    - business-logic
    - enumeration
    - api
    - owasp-a01
version: 1.0.0
---

# Mass Enumeration / Data Harvesting

## Overview
APIs that return unbounded lists or allow scraping of all records enable attackers to harvest user databases, email lists, phone numbers, and other sensitive data. This is especially critical for:
- User listing endpoints
- Search APIs with no result limits
- Sequential ID enumeration

## Remediation
- Enforce maximum page size (e.g., 100 records per page)
- Rate limit list endpoints per user/IP
- Require authentication for all list endpoints
- Use opaque cursors instead of offset-based pagination
