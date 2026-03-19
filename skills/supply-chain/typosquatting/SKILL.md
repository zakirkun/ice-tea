---
name: Typosquatting Package Names
version: 1.0.0
description: Detects imports of known typosquatting package names that mimic popular packages with malicious intent.
tags: [supply-chain, typosquatting, malware, owasp-a06]
languages: [javascript, typescript, python]
severity: critical
confidence: high
cwe: [CWE-494]
owasp: [A06:2025]
---

# Typosquatting Package Names

## Overview
Typosquatting packages mimic the names of popular packages with one-character differences (e.g., `lodahs` for `lodash`, `requets` for `requests`). They are published to public registries and contain malicious code to steal credentials or install backdoors.

Known historical examples: `event-stream`, `crossenv`, `babelcli`, `colourama`.

## Remediation
- Carefully verify package names before adding to dependencies
- Use `npm audit` and tools like Socket.dev, Snyk, or Dependabot
- Lock exact versions and verify package integrity hashes
