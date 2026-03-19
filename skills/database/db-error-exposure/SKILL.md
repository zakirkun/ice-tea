---
name: Database Error Information Exposure
version: 1.0.0
description: Detects database errors and query details sent directly to HTTP responses, leaking schema information.
tags: [database, information-disclosure, error-handling, owasp-a05]
languages: [javascript, typescript, python, go, java, php, ruby]
severity: high
confidence: high
cwe: [CWE-209]
owasp: [A05:2025]
---

# Database Error Information Exposure

## Overview
Database error messages sent to users reveal internal schema details:
- Table and column names
- SQL query structure
- Database version and type
- Internal file paths (SQLite, PostgreSQL data directory)

This information directly assists attackers in crafting SQL injection payloads.

## Remediation
Catch all database exceptions and return generic error messages. Log the actual error server-side.
