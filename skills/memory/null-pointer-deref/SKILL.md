---
name: Null Pointer Dereference
version: 1.0.0
description: Detects code that dereferences pointers or return values that could be NULL without validation, causing crashes or privilege escalation.
tags: [null-pointer, memory, c, cpp, owasp-a06]
languages: [c, cpp, java, go]
severity: high
confidence: medium
cwe: [CWE-476]
owasp: [A06:2025]
---

# Null Pointer Dereference

## Overview
Null pointer dereference occurs when code uses a pointer without checking whether it's NULL. This causes:
- **Crash/DoS**: SIGSEGV on Unix, access violation on Windows
- **Kernel privilege escalation**: NULL pointer dereference in kernel context can map page 0 and execute attacker code
- **Logic bypass**: Skipping NULL checks allows unexpected code paths

## Detection Strategy
- Return values of `malloc()`, `calloc()`, `realloc()` used without NULL check
- Results of `fopen()`, `popen()` dereferenced without NULL check
- Java objects returned from `getById()` or map lookups used without null check

## Remediation
Always check pointers for NULL before dereferencing.

**Vulnerable (C):**
```c
char *buf = malloc(256);
strcpy(buf, input); // buf might be NULL if malloc failed!
```

**Safe (C):**
```c
char *buf = malloc(256);
if (buf == NULL) { perror("malloc"); exit(1); }
strcpy(buf, input);
```
