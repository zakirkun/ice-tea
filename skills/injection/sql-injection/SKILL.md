---
confidence: high
cwe:
    - CWE-89
description: "Detects SQL injection via string concatenation and fmt.Sprintf in database queries. Use when reviewing code for SQLi, unsafe query construction, or auditing parameterized query usage."
languages:
    - go
    - javascript
    - python
    - java
    - php
    - kotlin
    - dart
    - zig
    - elixir
name: sql-injection-detection
owasp:
    - A05:2025
severity: critical
tags:
    - injection
    - sql
    - database
    - owasp-a05
version: 1.0.0
---

# SQL Injection Detection

Identifies SQL queries built with string concatenation or format strings where user input reaches execution sinks. See `patterns.yaml` for detection rules.

## Detection Workflow

1. **Import check** — look for database packages (`database/sql`, `gorm`, `sqlx`, ORMs)
2. **Call check** — find SQL execution functions (`Query`, `QueryRow`, `Exec`, `Raw`)
3. **Context check** — check if arguments include string concatenation or `fmt.Sprintf`
4. **Taint check** — trace whether the concatenated variable is user-controlled (request params, form data, headers)
5. **Validate** — confirm the variable is not a compile-time constant or allow-listed value before reporting

### Vulnerable

```go
db.Query("SELECT * FROM users WHERE id = " + userInput)
db.Query(fmt.Sprintf("SELECT * FROM users WHERE id = %s", userInput))
```

### Safe

```go
db.Query("SELECT * FROM users WHERE id = $1", userInput)
```
