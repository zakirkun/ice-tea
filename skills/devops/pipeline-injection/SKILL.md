---
name: CI/CD Pipeline Code Injection
version: 1.0.0
description: Detects CI/CD pipeline scripts that incorporate untrusted input into shell commands, enabling pipeline injection attacks.
tags: [devops, ci-cd, injection, owasp-a03]
languages: [yaml, generic]
severity: critical
confidence: medium
cwe: [CWE-77]
owasp: [A03:2025]
---

# CI/CD Pipeline Code Injection

## Overview
CI/CD pipelines that execute shell commands incorporating untrusted values (branch names, commit messages, PR metadata) are vulnerable to injection. Attackers can create branches or issues with malicious names to execute arbitrary commands in the CI environment, accessing secrets.

## Remediation
- Quote all shell variables in CI scripts
- Use CI-specific secret mechanisms instead of environment variable interpolation
- Validate and sanitize branch names and other git metadata before use in scripts
