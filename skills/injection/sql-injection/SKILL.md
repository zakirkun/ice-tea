---
name: SQL Injection Detection
version: 1.0.0
description: Detects SQL injection vulnerabilities where user input is concatenated into SQL queries
tags: [injection, sql, database, owasp-a05]
languages: [go, javascript, python, java, php]
severity: critical
confidence: high
cwe: [CWE-89]
owasp: [A05:2025]
---

# SQL Injection Detection

## Overview
SQL injection occurs when user-controlled input is incorporated into SQL queries without proper sanitization or parameterization.

## Detection Strategy
1. **Import check**: Look for database packages (database/sql, gorm, sqlx)
2. **Call check**: Find SQL execution functions (Query, Exec, Raw)
3. **Context check**: Check if arguments include string concatenation or fmt.Sprintf
4. **Taint check**: Trace if user input reaches SQL execution sinks

## Remediation
Use parameterized queries / prepared statements instead of string concatenation.

**Vulnerable:**
```go
db.Query("SELECT * FROM users WHERE id = " + userInput)
```

**Safe:**
```go
db.Query("SELECT * FROM users WHERE id = $1", userInput)
```
