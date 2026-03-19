---
confidence: medium
cwe:
    - CWE-918
description: Detects HTTP requests made with user-controlled URLs, potentially allowing internal network access.
languages:
    - generic
    - kotlin
    - dart
    - zig
    - elixir
name: Server-Side Request Forgery (SSRF)
owasp:
    - A10:2025
severity: high
tags:
    - ssrf
    - web
    - network
    - owasp-a10
version: 1.0.0
---

# Server-Side Request Forgery (SSRF)

## Overview
Detects HTTP requests made with user-controlled URLs, potentially allowing internal network access.

## Remediation
Validate URLs against a strict allowlist. Do not permit access to loopback or private IP ranges.
