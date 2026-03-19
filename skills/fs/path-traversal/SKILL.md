---
confidence: high
cwe:
    - CWE-22
description: Detects unsafe file access involving user input
languages:
    - go
    - python
    - kotlin
    - dart
    - zig
    - elixir
name: Path Traversal Detection
owasp:
    - A01:2025
severity: high
tags:
    - fs
    - traversal
    - file-access
    - owasp-a01
version: 1.0.0
---

# Path Traversal

## Overview
Path Traversal (Directory Traversal) occurs when user-supplied input is used to construct a file path without proper neutralization of special elements (like `../`).

## Remediation
Always sanitize user input, use `filepath.Clean` in Go, or `os.path.abspath` in Python, and verify that the final resolved path resides within the expected base directory.
