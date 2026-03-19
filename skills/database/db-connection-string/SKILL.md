---
confidence: high
cwe:
    - CWE-312
    - CWE-798
description: Detects database connection strings with embedded credentials hardcoded in source code.
languages:
    - javascript
    - typescript
    - python
    - go
    - java
    - php
    - ruby
    - generic
    - kotlin
    - dart
    - zig
    - elixir
name: Hardcoded Database Connection String
owasp:
    - A07:2025
severity: critical
tags:
    - database
    - secrets
    - hardcoded
    - owasp-a07
version: 1.0.0
---

# Hardcoded Database Connection String

## Overview
Database connection strings containing credentials embedded directly in source code are exposed whenever the code is shared, committed to version control, or viewed by anyone with code access. This is one of the most common causes of database credential leaks.

## Remediation
- Use environment variables for all credential components
- Use a secrets manager (AWS Secrets Manager, Vault, Azure Key Vault)
- Never commit `.env` files with real credentials

**Safe:**
```python
conn_str = os.environ["DATABASE_URL"]
```
