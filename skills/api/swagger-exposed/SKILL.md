---
name: Swagger / OpenAPI Documentation Exposed in Production
version: 1.0.0
description: Detects Swagger UI, Redoc, or OpenAPI documentation endpoints accessible without authentication in production.
tags: [api, swagger, information-disclosure, owasp-a05]
languages: [javascript, typescript, python, java, go]
severity: medium
confidence: high
cwe: [CWE-200]
owasp: [A05:2025]
---

# Swagger / OpenAPI Exposed in Production

## Overview
API documentation endpoints expose:
- Complete API schema including undocumented endpoints
- Authentication mechanisms and expected formats
- Example payloads that can be used for fuzzing
- Internal models and data structures

While useful in development, public Swagger in production gives attackers a full roadmap of the API.

## Remediation
- Restrict Swagger UI to internal networks or authenticated users
- Disable Swagger in production builds
- Use `NODE_ENV=production` check or similar to conditionally serve docs
