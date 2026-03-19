---
confidence: medium
cwe:
    - CWE-89
description: Detects untrusted input concatenated directly into SQL queries.
languages:
    - generic
    - kotlin
    - dart
    - zig
    - elixir
name: SQL Injection
owasp:
    - A03:2021
    - A03:2025
severity: critical
tags:
    - sqli
    - web
    - injection
    - database
version: 1.0.0
---

# SQL Injection

## Overview
Detects untrusted input concatenated directly into SQL queries.

## Remediation
Use parameterized queries or prepared statements instead of string concatenation.
