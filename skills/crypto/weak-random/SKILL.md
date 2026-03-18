---
name: Weak Random Number Generation
version: 1.0.0
description: Detects the use of PRNGs (Pseudo-Random Number Generators) that are not cryptographically secure.
tags: [crypto, random, prng, owasp-a02]
languages: [generic]
severity: medium
confidence: medium
cwe: [CWE-338]
owasp: [A02:2025]
---

# Weak Random Number Generation

## Overview
Detects the use of PRNGs (Pseudo-Random Number Generators) that are not cryptographically secure.

## Remediation
Use cryptographically secure PRNGs (CSPRNG), like crypto/rand in Go, secrets in Python, or Crypto.getRandomValues in JS.
