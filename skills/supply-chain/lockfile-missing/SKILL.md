---
confidence: medium
cwe:
    - CWE-494
description: Detects projects where lockfiles are missing or ignored in version control, allowing non-deterministic builds with potentially different dependency versions.
languages:
    - javascript
    - typescript
    - python
    - go
    - java
    - ruby
    - generic
    - kotlin
    - dart
    - zig
    - elixir
name: Missing or Ignored Dependency Lockfile
owasp:
    - A06:2025
severity: medium
tags:
    - supply-chain
    - lockfile
    - dependencies
    - owasp-a06
version: 1.0.0
---

# Missing or Ignored Dependency Lockfile

## Overview
Lockfiles (`package-lock.json`, `yarn.lock`, `Pipfile.lock`, `go.sum`, `Gemfile.lock`) pin exact dependency versions. Without them, `npm install` may install newer versions that introduce vulnerabilities or malicious code.

Gitignoring lockfiles in application projects is a security anti-pattern.

## Remediation
- Commit lockfiles to version control for all application projects
- Use `npm ci` instead of `npm install` in CI/CD
- Verify lockfile integrity in CI with `--frozen-lockfile`
