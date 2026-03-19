---
confidence: high
cwe:
    - CWE-1392
    - CWE-798
description: Detects use of common default usernames and passwords in application configuration and code.
languages:
    - generic
    - javascript
    - typescript
    - python
    - go
    - java
    - php
    - yaml
    - kotlin
    - dart
    - zig
    - elixir
name: Default Credentials
owasp:
    - A07:2025
severity: critical
tags:
    - default-credentials
    - auth
    - infra
    - owasp-a07
version: 1.0.0
---

# Default Credentials

## Overview
Default credentials are pre-configured usernames and passwords that ship with software, databases, or infrastructure components. Attackers maintain databases of default credentials and routinely attempt them. Common examples:
- `admin:admin`, `admin:password`, `root:root`
- Database defaults: `mysql root:""`, `postgres postgres:postgres`
- IoT/Router defaults: `admin:1234`, `admin:admin`

## Remediation
- Change all default credentials immediately upon installation
- Enforce strong, unique passwords for administrative accounts
- Use secrets management systems for credentials in code
- Detect default credentials in CI/CD pipelines
