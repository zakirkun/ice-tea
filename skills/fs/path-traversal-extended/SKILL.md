---
name: Path Traversal (Extended)
version: 1.0.0
description: Detects unsafe file writing or dynamic inclusions leading to LFI/path traversal.
tags: [fs, traversal, lfi, owasp-a01]
languages: [generic]
severity: high
confidence: medium
cwe: [CWE-22, CWE-98]
owasp: [A01:2025]
---

# Path Traversal (Extended)

## Overview
Detects unsafe file writing or dynamic inclusions leading to LFI/path traversal.

## Remediation
Sanitize paths and restrict inclusions to allowed directories.
