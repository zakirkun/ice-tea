---
confidence: medium
cwe:
    - CWE-367
description: Detects file operations that check file properties before using them, creating a race window exploitable via symlink attacks.
languages:
    - c
    - cpp
    - python
    - go
    - java
    - kotlin
    - dart
    - zig
    - elixir
name: File Time-of-Check / Time-of-Use (TOCTOU)
owasp:
    - A01:2025
severity: high
tags:
    - race-condition
    - toctou
    - filesystem
    - owasp-a01
version: 1.0.0
---

# File TOCTOU Race Condition

## Overview
TOCTOU (Time-of-Check/Time-of-Use) vulnerabilities occur when there is a window between checking a file's state and using it. An attacker can replace the file with a symlink between the check and the use, potentially reading/writing arbitrary files.

Classic pattern: `if access(path, R_OK) == 0: open(path)` — between access() and open(), attacker creates symlink.

## Remediation
- Use `O_NOFOLLOW` flag to prevent symlink following
- Use `openat()` with AT_FDCWD to operate atomically
- In Python, use `os.open()` with `os.O_NOFOLLOW`
- Validate path within a trusted directory after opening
