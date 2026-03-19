---
confidence: medium
cwe:
    - CWE-494
description: Detects package.json configurations vulnerable to dependency confusion attacks where internal package names could be hijacked via public npm registry.
languages:
    - javascript
    - typescript
    - generic
    - kotlin
    - dart
    - zig
    - elixir
name: NPM Dependency Confusion Attack
owasp:
    - A06:2025
severity: high
tags:
    - supply-chain
    - npm
    - dependency-confusion
    - owasp-a06
version: 1.0.0
---

# NPM Dependency Confusion Attack

## Overview
Dependency confusion exploits package manager precedence: if an internal package is published to the public npm registry with a higher version, it gets installed instead of the internal one. Attackers scan job postings and error messages for internal package names.

## Remediation
- Use scoped packages: `@company/internal-package` (scoped packages can only conflict in the same scope)
- Configure `npmrc` to always use internal registry for internal packages
- Use `package-lock.json` integrity checks
- Add `publishConfig` to prevent accidental public publishing
