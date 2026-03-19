---
name: Mass Assignment
version: 1.0.0
description: Detects frameworks binding raw HTTP payload bodies directly to database models or internal objects without field allow-lists.
tags: [api, backend, mass-assignment, owasp-a08]
languages: [generic]
severity: high
confidence: medium
cwe: [CWE-915]
owasp: [A08:2025]
---

# Mass Assignment

## Overview
Detects frameworks binding raw HTTP payload bodies directly to database models or internal objects without field allow-lists.

## Remediation
Use explicitly defined Data Transfer Objects (DTOs) or field allow-lists rather than blindly binding raw JSON to models.
