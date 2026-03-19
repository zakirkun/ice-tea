---
confidence: medium
cwe:
    - CWE-915
description: Detects frameworks binding raw HTTP payload bodies directly to database models or internal objects without field allow-lists.
languages:
    - generic
    - kotlin
    - dart
    - zig
    - elixir
name: Mass Assignment
owasp:
    - A08:2025
severity: high
tags:
    - api
    - backend
    - mass-assignment
    - owasp-a08
version: 1.0.0
---

# Mass Assignment

## Overview
Detects frameworks binding raw HTTP payload bodies directly to database models or internal objects without field allow-lists.

## Remediation
Use explicitly defined Data Transfer Objects (DTOs) or field allow-lists rather than blindly binding raw JSON to models.
