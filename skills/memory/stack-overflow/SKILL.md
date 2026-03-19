---
confidence: medium
cwe:
    - CWE-121
description: Detects code patterns that can cause stack overflows via unbounded recursion or excessive stack allocation.
languages:
    - c
    - cpp
    - java
    - go
    - python
    - kotlin
    - dart
    - zig
    - elixir
name: Stack Overflow Vulnerability
owasp:
    - A06:2025
severity: high
tags:
    - memory
    - stack-overflow
    - c
    - cpp
    - owasp-a06
version: 1.0.0
---

# Stack Overflow

## Overview
Stack overflows occur when the call stack exceeds available memory:
- **Unbounded recursion**: Recursive functions without base case or depth limit
- **Large stack allocations**: Variable-length arrays (VLAs) with user-controlled size
- **Deep call chains**: Complex XML/JSON parsing with deeply nested structures

In C/C++, stack overflows corrupt adjacent memory and can lead to code execution.

## Remediation
- Add recursion depth limits
- Use iterative algorithms instead of recursive ones for user-controlled depth
- Validate nesting depth of parsed structures
- Use `ulimit` or OS-level stack size limits
