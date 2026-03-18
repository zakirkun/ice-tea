---
name: Hardcoded Credentials Detection
version: 1.0.0
description: Detects hardcoded passwords, API keys, and tokens in source code
tags: [credentials, secrets, hardcoded, owasp-a07]
languages: [go, javascript, python, java, php]
severity: high
confidence: medium
cwe: [CWE-798]
owasp: [A07:2025]
---

# Hardcoded Credentials Detection

## Overview
Hardcoded credentials in source code can be extracted by attackers with access to the codebase or compiled binaries.

## Remediation
Use environment variables, secrets managers, or configuration files (excluded from version control) to manage sensitive credentials.
