---
confidence: medium
cwe:
    - CWE-502
description: Detects deserialization of untrusted data which can lead to Remote Code Execution.
languages:
    - generic
    - kotlin
    - dart
    - zig
    - elixir
name: Insecure Deserialization
owasp:
    - A08:2025
severity: critical
tags:
    - deserialization
    - rce
    - owasp-a08
version: 1.0.0
---

# Insecure Deserialization

## Overview
Detects deserialization of untrusted data which can lead to Remote Code Execution.

## Remediation
Use safe data formats like JSON instead of native serialization. If native serialization is required, sign the objects.
