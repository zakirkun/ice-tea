---
name: Hardcoded Secrets
version: 1.0.0
description: Detects API keys, passwords, and tokens embedded directly in the source code.
tags: [secrets, auth, hardcoded, owasp-a07]
languages: [generic]
severity: critical
confidence: medium
cwe: [CWE-798]
owasp: [A07:2025]
---

# Hardcoded Secrets

## Overview
Detects API keys, passwords, and tokens embedded directly in the source code.

## Remediation
Use a secure secrets manager or environment variables to inject sensitive credentials at runtime.
