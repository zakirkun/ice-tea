---
name: Server-Side Request Forgery (SSRF)
version: 1.0.0
description: Detects HTTP requests made with user-controlled URLs, potentially allowing internal network access.
tags: [ssrf, web, network, owasp-a10]
languages: [generic]
severity: high
confidence: medium
cwe: [CWE-918]
owasp: [A10:2025]
---

# Server-Side Request Forgery (SSRF)

## Overview
Detects HTTP requests made with user-controlled URLs, potentially allowing internal network access.

## Remediation
Validate URLs against a strict allowlist. Do not permit access to loopback or private IP ranges.
