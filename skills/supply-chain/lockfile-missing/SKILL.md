---
name: Missing or Ignored Dependency Lockfile
version: 1.0.0
description: Detects projects where lockfiles are missing or ignored in version control, allowing non-deterministic builds with potentially different dependency versions.
tags: [supply-chain, lockfile, dependencies, owasp-a06]
languages: [javascript, typescript, python, go, java, ruby, generic]
severity: medium
confidence: medium
cwe: [CWE-494]
owasp: [A06:2025]
---

# Missing or Ignored Dependency Lockfile

## Overview
Lockfiles (`package-lock.json`, `yarn.lock`, `Pipfile.lock`, `go.sum`, `Gemfile.lock`) pin exact dependency versions. Without them, `npm install` may install newer versions that introduce vulnerabilities or malicious code.

Gitignoring lockfiles in application projects is a security anti-pattern.

## Remediation
- Commit lockfiles to version control for all application projects
- Use `npm ci` instead of `npm install` in CI/CD
- Verify lockfile integrity in CI with `--frozen-lockfile`
