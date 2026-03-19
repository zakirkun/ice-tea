---
name: ORM Raw Query Injection
version: 1.0.0
description: Detects raw SQL queries embedded within ORM frameworks that bypass parameterization, reintroducing SQL injection risk.
tags: [database, sql-injection, orm, owasp-a03]
languages: [javascript, typescript, python, java, go, php, ruby]
severity: high
confidence: medium
cwe: [CWE-89]
owasp: [A03:2025]
---

# ORM Raw Query Injection

## Overview
ORM frameworks like SQLAlchemy, Hibernate, Sequelize, and GORM provide parameterized query builders that prevent SQL injection. However, they also expose `raw()`, `execute()`, and `query()` escape hatches that, when used with string concatenation or f-strings, reintroduce SQL injection vulnerabilities.

## Remediation
Use the ORM's parameterized query API instead of raw query methods.

**Vulnerable (Python/SQLAlchemy):**
```python
db.execute(f"SELECT * FROM users WHERE username = '{username}'")
```

**Safe (Python/SQLAlchemy):**
```python
db.execute(text("SELECT * FROM users WHERE username = :username"), {"username": username})
```
