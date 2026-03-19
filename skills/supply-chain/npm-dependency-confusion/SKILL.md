---
name: NPM Dependency Confusion Attack
version: 1.0.0
description: Detects package.json configurations vulnerable to dependency confusion attacks where internal package names could be hijacked via public npm registry.
tags: [supply-chain, npm, dependency-confusion, owasp-a06]
languages: [javascript, typescript, generic]
severity: high
confidence: medium
cwe: [CWE-494]
owasp: [A06:2025]
---

# NPM Dependency Confusion Attack

## Overview
Dependency confusion exploits package manager precedence: if an internal package is published to the public npm registry with a higher version, it gets installed instead of the internal one. Attackers scan job postings and error messages for internal package names.

## Remediation
- Use scoped packages: `@company/internal-package` (scoped packages can only conflict in the same scope)
- Configure `npmrc` to always use internal registry for internal packages
- Use `package-lock.json` integrity checks
- Add `publishConfig` to prevent accidental public publishing
