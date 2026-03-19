---
confidence: medium
cwe:
    - CWE-22
    - CWE-98
description: Detects unsafe file writing or dynamic inclusions leading to LFI/path traversal.
languages:
    - generic
    - kotlin
    - dart
    - zig
    - elixir
name: Path Traversal (Extended)
owasp:
    - A01:2025
severity: high
tags:
    - fs
    - traversal
    - lfi
    - owasp-a01
version: 1.0.0
---

# Path Traversal (Extended)

## Overview
Detects unsafe file writing or dynamic inclusions leading to LFI/path traversal.

## Remediation
Sanitize paths and restrict inclusions to allowed directories.
