---
confidence: medium
cwe:
    - CWE-530
description: Detects build artifacts, debug information, and development files committed to version control or accessible in production.
languages:
    - generic
    - yaml
    - kotlin
    - dart
    - zig
    - elixir
name: Exposed Build Artifacts and Debug Files
owasp:
    - A05:2025
severity: medium
tags:
    - devops
    - information-disclosure
    - owasp-a05
version: 1.0.0
---

# Exposed Build Artifacts and Debug Files

## Overview
Build artifacts, debug symbols, and development configuration files accidentally committed or deployed expose internal application structure and potentially sensitive logic. Common exposures:
- `.map` source map files in production (exposes original source)
- `*.pdb` debug symbol files
- `node_modules/` committed to git
- `.DS_Store` files leaking directory structure
- `TODO`, `FIXME` comments with security implications

## Remediation
- Add build artifacts to `.gitignore`
- Disable source map generation in production builds
- Serve content with proper `X-Content-Type-Options: nosniff` headers
