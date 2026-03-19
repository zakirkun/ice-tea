---
name: Wild Pointer / Uninitialized Pointer
version: 1.0.0
description: Detects uninitialized or dangling pointer usage that can lead to arbitrary memory access.
tags: [memory, wild-pointer, c, cpp, owasp-a06]
languages: [c, cpp]
severity: high
confidence: medium
cwe: [CWE-457]
owasp: [A06:2025]
---

# Wild / Uninitialized Pointer

## Overview
Wild pointers are pointers that have not been initialized to a valid address. Dereferencing them reads from or writes to arbitrary memory, leading to:
- Crashes (SIGSEGV)
- Data corruption
- Security vulnerabilities (if attacker controls the uninitialized memory)

Common causes:
- Pointer declared but not initialized: `int *p;` then `*p = 1;`
- Pointer used after `free()` without zeroing
- Conditional initialization where some paths skip initialization

## Remediation
- Initialize all pointers to `NULL` at declaration
- Use address sanitizers to detect wild pointer access
- Use C++ smart pointers instead of raw pointers
