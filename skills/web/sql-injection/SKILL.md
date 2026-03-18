---
name: SQL Injection
version: 1.0.0
description: Detects untrusted input concatenated directly into SQL queries.
tags: [sqli, web, injection, database]
languages: [generic]
severity: critical
confidence: medium
cwe: [CWE-89]
owasp: [A03:2021, A03:2025]
---

# SQL Injection

## Overview
Detects untrusted input concatenated directly into SQL queries.

## Remediation
Use parameterized queries or prepared statements instead of string concatenation.
