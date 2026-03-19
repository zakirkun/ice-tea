---
confidence: high
cwe:
    - CWE-526
    - CWE-798
description: Detects sensitive secrets hardcoded in .env files, configuration files, and infrastructure definitions that may be committed to version control.
languages:
    - generic
    - yaml
    - kotlin
    - dart
    - zig
    - elixir
name: Secrets in Environment Files and Configuration
owasp:
    - A07:2025
severity: critical
tags:
    - secrets
    - env
    - configuration
    - owasp-a07
version: 1.0.0
---

# Secrets in Environment Files and Configuration

## Overview
Secrets in `.env`, `config.yaml`, `docker-compose.yml`, and similar files are frequently committed to Git repositories (including public ones), exposing credentials, API keys, and tokens. This is one of the most common causes of cloud breaches.

## Detection Strategy
- `.env` files with actual values (not placeholders)
- `docker-compose.yml` with hardcoded credentials
- CI/CD configuration files with secrets

## Remediation
- Use `.gitignore` to exclude `.env` files from version control
- Use secret management services (AWS Secrets Manager, HashiCorp Vault, GCP Secret Manager)
- Use `${VARIABLE}` placeholders in config files, with actual values injected at runtime
- Rotate any secrets that have been committed — assume they are compromised
