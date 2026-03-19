---
name: Python Pip Dependency Confusion
version: 1.0.0
description: Detects Python package configurations vulnerable to dependency confusion attacks through PyPI.
tags: [supply-chain, pip, python, dependency-confusion, owasp-a06]
languages: [python, generic]
severity: high
confidence: medium
cwe: [CWE-494]
owasp: [A06:2025]
---

# Python Pip Dependency Confusion

## Overview
Similar to npm dependency confusion, attackers can publish packages with internal names to PyPI. When `pip` searches for a package, public PyPI is checked and a higher version number wins over private index entries.

## Remediation
- Use `--index-url` pointing to private registry with `--extra-index-url` for PyPI fallback
- Use `--no-index` with `--find-links` for air-gapped installs
- Add package to PyPI placeholder to prevent namespace hijacking
