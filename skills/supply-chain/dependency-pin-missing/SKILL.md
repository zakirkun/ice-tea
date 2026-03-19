---
confidence: medium
cwe:
    - CWE-494
description: Detects dependencies specified with loose version constraints that may install different code between builds.
languages:
    - generic
    - python
    - javascript
    - typescript
    - go
    - kotlin
    - dart
    - zig
    - elixir
name: Unpinned Dependency Versions
owasp:
    - A06:2025
severity: medium
tags:
    - supply-chain
    - dependencies
    - version-pinning
    - owasp-a06
version: 1.0.0
---

# Unpinned Dependency Versions

## Overview
Using version ranges (`^1.2.3`, `~1.2.3`, `>=1.0.0`, `*`, `latest`) instead of exact pinned versions means builds can install different dependency versions over time. This allows:
- Accidentally upgrading to a version with new vulnerabilities
- Installing a maliciously updated package
- Non-reproducible builds

## Remediation
- Pin exact versions in production: `"express": "4.18.2"` not `"^4.18.2"`
- Use lockfiles to ensure reproducible installs
- Update dependencies deliberately with security review
