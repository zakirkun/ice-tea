---
name: Use After Free
version: 1.0.0
description: Detects potential use-after-free vulnerabilities where memory is accessed after being freed.
tags: [use-after-free, memory, c, cpp, owasp-a06]
languages: [c, cpp]
severity: critical
confidence: medium
cwe: [CWE-416]
owasp: [A06:2025]
---

# Use After Free

## Overview
Use-after-free (UAF) occurs when a program continues to use a pointer after the memory it points to has been freed. This can lead to:
- Memory corruption
- Arbitrary code execution
- Privilege escalation
- Information disclosure

UAF vulnerabilities are common in browsers, kernels, and network services.

## Detection Strategy
- `free(ptr)` followed by accessing `ptr` without setting it to `NULL`
- `delete ptr` in C++ followed by further dereference
- Returning a pointer to freed memory

## Remediation
- Set pointers to `NULL` immediately after `free()`
- Use smart pointers in C++ (`unique_ptr`, `shared_ptr`)
- Use memory-safe languages for new projects

**Vulnerable:**
```c
char *buf = malloc(256);
free(buf);
strcpy(buf, input); // Use after free!
```

**Safe:**
```c
char *buf = malloc(256);
free(buf);
buf = NULL; // Prevent use after free
```
