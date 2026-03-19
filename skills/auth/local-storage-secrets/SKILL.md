---
name: Sensitive Data in LocalStorage
version: 1.0.0
description: Detects potential storage of sensitive credentials or JWTs directly in the browser's localStorage or sessionStorage.
tags: [frontend, storage, secrets, owasp-a02]
languages: [generic]
severity: medium
confidence: medium
cwe: [CWE-312]
owasp: [A02:2025]
---

# Sensitive Data in LocalStorage

## Overview
Detects potential storage of sensitive credentials or JWTs directly in the browser's localStorage or sessionStorage.

## Remediation
Use HttpOnly/Secure cookies for session tokens. Local storage is easily accessible via any XSS vulnerability.
