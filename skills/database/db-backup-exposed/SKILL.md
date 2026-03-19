---
confidence: medium
cwe:
    - CWE-530
description: Detects database backup files stored in web-accessible directories or referenced with predictable names.
languages:
    - generic
    - php
    - python
    - javascript
    - go
    - kotlin
    - dart
    - zig
    - elixir
name: Database Backup File Exposure
owasp:
    - A05:2025
severity: critical
tags:
    - database
    - backup
    - information-disclosure
    - owasp-a05
version: 1.0.0
---

# Database Backup File Exposure

## Overview
Database backup files (`.sql`, `.dump`, `.bak`, `.db`, `.sqlite`) stored in web-accessible directories can be directly downloaded by attackers, exposing the full database contents including all user data and credentials.

## Remediation
- Store database backups outside the web root
- Restrict backup file access with web server rules
- Encrypt backup files at rest
- Use access-controlled cloud storage for backups
