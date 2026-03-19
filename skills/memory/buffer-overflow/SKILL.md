---
confidence: high
cwe:
    - CWE-120
    - CWE-119
    - CWE-787
description: Detects unsafe C/C++ functions that copy data without bounds checking, enabling stack and heap buffer overflows.
languages:
    - c
    - cpp
    - kotlin
    - dart
    - zig
    - elixir
name: Buffer Overflow
owasp:
    - A06:2025
severity: critical
tags:
    - buffer-overflow
    - memory
    - c
    - cpp
    - owasp-a06
version: 1.0.0
---

# Buffer Overflow

## Overview
Buffer overflows occur when data is written beyond the boundaries of allocated memory. In C/C++, unsafe library functions have no bounds checking:
- `gets()`: Reads unlimited input into fixed buffer — deprecated and removed from C11
- `strcpy()`: Copies until null terminator with no size limit
- `strcat()`: Concatenates with no size limit
- `sprintf()`: Formats with no output buffer size check

Consequences: Stack corruption, return address overwrite, arbitrary code execution, privilege escalation.

## Remediation
Replace unsafe functions with safe alternatives:
- `gets()` → `fgets(buf, sizeof(buf), stdin)` or `getline()`
- `strcpy()` → `strncpy()` + manual null termination, or `strlcpy()`
- `strcat()` → `strncat()` or `strlcat()`
- `sprintf()` → `snprintf(buf, sizeof(buf), ...)`
