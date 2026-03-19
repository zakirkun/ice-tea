---
confidence: low
cwe:
    - CWE-362
description: Detects singleton implementations that are not thread-safe and can result in multiple instances being created under concurrent access.
languages:
    - java
    - go
    - python
    - javascript
    - typescript
    - kotlin
    - dart
    - zig
    - elixir
name: Thread-Unsafe Singleton Pattern
owasp:
    - A06:2025
severity: medium
tags:
    - race-condition
    - concurrency
    - singleton
    - owasp-a06
version: 1.0.0
---

# Thread-Unsafe Singleton

## Overview
Lazy-initialized singletons without thread synchronization can result in multiple instances being created when the class is initialized concurrently by multiple threads. This causes inconsistent state, duplicate resource consumption, and potential security issues if security-related objects (loggers, auth managers) are duplicated.

## Remediation
- Java: Use `enum`, `static initialization block`, or `synchronized` + `volatile`
- Go: Use `sync.Once`
- Python: Use module-level variable (thread-safe by GIL) or `threading.Lock`
