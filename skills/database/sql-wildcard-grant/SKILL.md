---
confidence: high
cwe:
    - CWE-732
description: Detects SQL GRANT statements with wildcard privileges that over-privilege database users.
languages:
    - generic
    - sql
    - kotlin
    - dart
    - zig
    - elixir
name: SQL Wildcard GRANT Privilege
owasp:
    - A01:2025
severity: high
tags:
    - database
    - sql
    - privilege
    - owasp-a01
version: 1.0.0
---

# SQL Wildcard GRANT Privilege

## Overview
GRANT ALL PRIVILEGES or GRANT * in SQL scripts gives a database user excessive permissions. Application database users should only have the minimum permissions required (SELECT, INSERT, UPDATE, DELETE on specific tables).

## Remediation
Follow the principle of least privilege for database users. Grant only the specific operations needed on specific tables.

**Vulnerable:**
```sql
GRANT ALL PRIVILEGES ON *.* TO 'app_user'@'%';
```

**Safe:**
```sql
GRANT SELECT, INSERT, UPDATE, DELETE ON myapp.users TO 'app_user'@'%';
```
