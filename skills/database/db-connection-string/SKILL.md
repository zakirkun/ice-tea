---
name: Hardcoded Database Connection String
version: 1.0.0
description: Detects database connection strings with embedded credentials hardcoded in source code.
tags: [database, secrets, hardcoded, owasp-a07]
languages: [javascript, typescript, python, go, java, php, ruby, generic]
severity: critical
confidence: high
cwe: [CWE-312, CWE-798]
owasp: [A07:2025]
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
