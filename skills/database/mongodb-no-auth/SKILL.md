---
confidence: high
cwe:
    - CWE-306
description: Detects MongoDB connections without authentication credentials or configurations that disable authorization.
languages:
    - javascript
    - typescript
    - python
    - go
    - java
    - kotlin
    - dart
    - zig
    - elixir
name: MongoDB Without Authentication
owasp:
    - A05:2025
severity: critical
tags:
    - database
    - mongodb
    - auth
    - owasp-a05
version: 1.0.0
---

# MongoDB Without Authentication

## Overview
MongoDB instances without authentication allow any client that reaches the port to read, write, or delete all databases. This has led to numerous mass data breaches. Default MongoDB installations prior to 3.0 had no authentication.

## Remediation
- Enable `--auth` flag or set `security.authorization: enabled` in `mongod.conf`
- Create admin user with strong password before exposing MongoDB
- Bind to localhost or use VPC/firewall rules
- Use TLS for connections
