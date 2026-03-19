---
name: Database Backup File Exposure
version: 1.0.0
description: Detects database backup files stored in web-accessible directories or referenced with predictable names.
tags: [database, backup, information-disclosure, owasp-a05]
languages: [generic, php, python, javascript, go]
severity: critical
confidence: medium
cwe: [CWE-530]
owasp: [A05:2025]
---

# Database Backup File Exposure

## Overview
Database backup files (`.sql`, `.dump`, `.bak`, `.db`, `.sqlite`) stored in web-accessible directories can be directly downloaded by attackers, exposing the full database contents including all user data and credentials.

## Remediation
- Store database backups outside the web root
- Restrict backup file access with web server rules
- Encrypt backup files at rest
- Use access-controlled cloud storage for backups
