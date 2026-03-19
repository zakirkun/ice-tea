---
name: Information Exposure via Errors
version: 1.0.0
description: Detects raw error stacks or generic exception details being directly returned in HTTP responses.
tags: [api, backend, information-exposure, owasp-a05]
languages: [generic]
severity: low
confidence: medium
cwe: [CWE-209]
owasp: [A05:2025]
---

# Information Exposure via Errors

## Overview
Detects raw error stacks or generic exception details being directly returned in HTTP responses.

## Remediation
Catch exceptions and return generic (e.g. 500 Internal Server Error) responses in production. Log the actual stack trace securely on the backend.
