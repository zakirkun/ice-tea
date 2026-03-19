---
name: CI/CD Pipeline Secrets Exposure
version: 1.0.0
description: Detects secrets hardcoded in CI/CD configuration files (GitLab CI, GitHub Actions, CircleCI, Jenkins).
tags: [devops, ci-cd, secrets, owasp-a07]
languages: [yaml, generic]
severity: critical
confidence: high
cwe: [CWE-312, CWE-798]
owasp: [A07:2025]
---

# CI/CD Pipeline Secrets Exposure

## Overview
CI/CD configuration files committed to version control often contain hardcoded API keys, tokens, or passwords. These are frequently visible in:
- Public repositories
- Git history (even after removal)
- Third-party services that access the repository

## Remediation
- Use CI/CD secret management (GitHub Secrets, GitLab CI Variables, CircleCI Contexts)
- Reference secrets as `${{ secrets.MY_SECRET }}` or `$MY_VAR`
- Scan git history with tools like `trufflehog` or `gitleaks`
