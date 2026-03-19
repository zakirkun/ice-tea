---
confidence: medium
cwe:
    - CWE-311
description: Detects database configurations and connections missing encryption at rest and in transit.
languages:
    - python
    - javascript
    - typescript
    - go
    - java
    - php
    - yaml
    - kotlin
    - dart
    - zig
    - elixir
name: Unencrypted Database
owasp:
    - A02:2025
severity: high
tags:
    - database
    - encryption
    - owasp-a02
version: 1.0.0
---

# Unencrypted Database

## Overview
Databases storing sensitive data must be encrypted both at rest (disk encryption) and in transit (TLS). Without encryption at rest, physical access or cloud storage access exposes all data. Without TLS, credentials and data can be intercepted on the network.

## Remediation
- Enable TLS/SSL in database connections (`sslmode=require` for PostgreSQL)
- Use encrypted database volumes (AWS RDS encryption, cloud disk encryption)
- Use SQLCipher for SQLite encryption
