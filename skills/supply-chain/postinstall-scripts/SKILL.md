---
confidence: high
cwe:
    - CWE-506
description: Detects npm/pip package lifecycle scripts that download and execute code, a common malicious package technique.
languages:
    - javascript
    - typescript
    - generic
    - kotlin
    - dart
    - zig
    - elixir
name: Dangerous Package Lifecycle Scripts
owasp:
    - A06:2025
severity: high
tags:
    - supply-chain
    - npm
    - lifecycle-scripts
    - owasp-a06
version: 1.0.0
---

# Dangerous Package Lifecycle Scripts

## Overview
npm package `postinstall`, `preinstall`, and `install` scripts run automatically when a package is installed. Malicious packages abuse this to:
- Download and execute a remote payload
- Exfiltrate environment variables (API keys, AWS credentials)
- Install persistent backdoors

## Remediation
- Audit all `postinstall` scripts in `node_modules`
- Use `npm install --ignore-scripts` for packages that don't need build steps
- Use tools like `npm audit`, `socket.dev`, or `snyk` to scan packages
