---
confidence: high
cwe:
    - CWE-377
    - CWE-379
description: Detects creation of temporary files with predictable names or in world-writable locations, enabling symlink attacks and race conditions.
languages:
    - python
    - javascript
    - typescript
    - go
    - java
    - php
    - c
    - cpp
    - kotlin
    - dart
    - zig
    - elixir
name: Unsafe Temporary File Creation
owasp:
    - A01:2025
severity: medium
tags:
    - tempfile
    - filesystem
    - race-condition
    - owasp-a01
version: 1.0.0
---

# Unsafe Temporary File Creation

## Overview
Creating temporary files insecurely can lead to:
- **Symlink attacks**: Attacker pre-creates a symlink at the predicted temp path, redirecting writes to sensitive files
- **TOCTOU race conditions**: Time-of-check vs time-of-use between name generation and file open
- **Information disclosure**: Predictable temp file names allow attackers to read sensitive data

## Detection Strategy
- Using `mktemp` (insecure, returns name only) instead of `mkstemp` (securely opens file)
- Building temp file paths from predictable strings: `/tmp/app_` + PID
- Using `/tmp` directly with hardcoded suffixes

## Remediation
Use language-provided secure temp file APIs that atomically create and open the file.

**Vulnerable (Python):**
```python
import tempfile, os
tmp_path = '/tmp/myapp_' + str(os.getpid())
with open(tmp_path, 'w') as f:  # Race condition!
    f.write(data)
```

**Safe (Python):**
```python
import tempfile
with tempfile.NamedTemporaryFile(mode='w', delete=True) as f:
    f.write(data)
```
