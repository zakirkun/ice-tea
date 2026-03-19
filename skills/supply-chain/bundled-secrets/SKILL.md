---
name: Secrets Bundled in Published Packages
version: 1.0.0
description: Detects API keys, private keys, and credentials accidentally included in npm/pip packages via missing .npmignore or .pypiignore.
tags: [supply-chain, secrets, npm, owasp-a07]
languages: [generic]
severity: critical
confidence: medium
cwe: [CWE-312, CWE-798]
owasp: [A07:2025]
---

# Secrets Bundled in Published Packages

## Overview
When publishing packages without a proper `.npmignore` or explicit `files` field in `package.json`, sensitive files like `.env`, `*.pem`, and test credentials get included in the published artifact and become visible to anyone who installs the package.

## Remediation
- Add `.npmignore` with patterns for `.env`, `*.key`, `*.pem`, `test/`, `secrets/`
- OR use the `files` field in `package.json` to explicitly list only publishable files
- Run `npm publish --dry-run` to see what will be included
