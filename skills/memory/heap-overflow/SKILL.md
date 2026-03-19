---
name: Heap Buffer Overflow
version: 1.0.0
description: Detects heap buffer overflows from unsafe memory operations and insufficient size validation.
tags: [memory, heap-overflow, c, cpp, owasp-a06]
languages: [c, cpp]
severity: critical
confidence: medium
cwe: [CWE-122]
owasp: [A06:2025]
---

# Heap Buffer Overflow

## Overview
Heap overflows occur when data is written beyond the bounds of a dynamically allocated buffer. Unlike stack overflows, heap overflows can overwrite:
- Adjacent allocations (other object data)
- Heap metadata (size, free list pointers)
- Function pointers stored in objects

Modern heap exploitation techniques include:
- Tcache poisoning (glibc)
- Unsafe unlink attacks
- House of Force

## Remediation
- Always validate buffer sizes before copying
- Use safe string functions with explicit size limits
- Enable compiler hardening: `-D_FORTIFY_SOURCE=2`, AddressSanitizer
