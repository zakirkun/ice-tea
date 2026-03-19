---
name: Missing API Rate Limiting
version: 1.0.0
description: Detects sensitive API endpoints without rate limiting, enabling resource exhaustion, credential stuffing, and enumeration attacks.
tags: [rate-limiting, api, dos, owasp-api4]
languages: [javascript, typescript, python, go, java, php]
severity: high
confidence: low
cwe: [CWE-770, CWE-307]
owasp: [A04:2025]
---

# Missing API Rate Limiting

## Overview
APIs without rate limiting are vulnerable to:
- **Credential stuffing**: Automated login attempts with breached credentials
- **Resource exhaustion**: Expensive computation triggered repeatedly (DoS)
- **Data harvesting**: Scraping all records via automated enumeration
- **OTP brute force**: Guessing 6-digit codes in 1,000,000 requests

## Detection Strategy
Identify API endpoints handling authentication, password reset, OTP verification, or resource-intensive operations that lack rate limiting middleware.

## Remediation
Apply rate limiting at the API gateway or application level with per-IP or per-user quotas.
