---
name: Path Traversal Detection
version: 1.0.0
description: Detects unsafe file access involving user input
tags: [fs, traversal, file-access, owasp-a01]
languages: [go, python]
severity: high
confidence: high
cwe: [CWE-22]
owasp: [A01:2025]
---

# Path Traversal

## Overview
Path Traversal (Directory Traversal) occurs when user-supplied input is used to construct a file path without proper neutralization of special elements (like `../`).

## Remediation
Always sanitize user input, use `filepath.Clean` in Go, or `os.path.abspath` in Python, and verify that the final resolved path resides within the expected base directory.
