---
confidence: medium
cwe:
    - CWE-190
    - CWE-191
description: Detects arithmetic operations that can overflow or underflow integer bounds, leading to heap overflows, logic bypasses, or unexpected behavior.
languages:
    - c
    - cpp
    - java
    - go
    - kotlin
    - dart
    - zig
    - elixir
name: Integer Overflow / Underflow
owasp:
    - A06:2025
severity: high
tags:
    - integer-overflow
    - memory
    - c
    - cpp
    - owasp-a06
version: 1.0.0
---

# Integer Overflow / Underflow

## Overview
Integer overflow occurs when arithmetic results exceed the data type's maximum value. In C/C++:
- Signed integer overflow is **undefined behavior** (UB) — compilers can optimize it away
- Unsigned overflow wraps around (e.g., `0 - 1 = UINT_MAX`)

Common exploitable scenarios:
- **Heap overflow**: `malloc(user_len + 4)` — if `user_len = UINT_MAX - 3`, result wraps to small allocation
- **Loop condition bypass**: `for (i = 0; i <= MAX_INT; i++)` — wraps to 0
- **Signed/unsigned mismatch**: `if (len < MAX)` where `len` is unsigned and `MAX` is signed

## Remediation
- Validate input sizes before arithmetic operations
- Use safe integer libraries (`SafeInt`, `checked_arithmetic`)
- Use compiler sanitizers (UBSan, ASan) to detect overflow
