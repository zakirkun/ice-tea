---
name: Memory Leak
version: 1.0.0
description: Detects common memory leak patterns including unreleased heap allocations, unclosed file handles, and missing deallocation in error paths.
tags: [memory, memory-leak, c, cpp, owasp-a06]
languages: [c, cpp, go, java]
severity: medium
confidence: medium
cwe: [CWE-401]
owasp: [A06:2025]
---

# Memory Leak

## Overview
Memory leaks cause gradual memory consumption growth until the process exhausts available memory (OOM). In long-running servers, even small leaks cause eventual crashes or degraded performance.

Common patterns:
- `malloc()` without corresponding `free()`
- `new` without `delete` in C++
- Missing `free()` in error exit paths
- Goroutine leaks in Go

## Remediation
- Use RAII (C++) or smart pointers to manage memory automatically
- Use ASan/Valgrind to detect leaks
- Go: Use context cancellation to stop goroutines
- Ensure all code paths `free()` allocated resources
