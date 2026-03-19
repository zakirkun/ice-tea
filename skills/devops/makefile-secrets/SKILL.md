---
name: Hardcoded Secrets in Makefiles
version: 1.0.0
description: Detects API keys, tokens, and passwords hardcoded in Makefile targets and variables.
tags: [devops, makefile, secrets, owasp-a07]
languages: [generic]
severity: critical
confidence: high
cwe: [CWE-798]
owasp: [A07:2025]
---

# Hardcoded Secrets in Makefiles

## Overview
Makefiles often contain deployment commands that require credentials. These credentials hardcoded in Makefiles are committed to version control and visible in CI/CD logs.

## Remediation
- Use environment variables for all credentials: `$(API_KEY)` from shell environment
- Use `.env` files loaded before `make` commands (not committed)
- Use vault agents or cloud secrets in production pipelines
