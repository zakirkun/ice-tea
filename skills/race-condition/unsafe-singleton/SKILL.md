---
name: Thread-Unsafe Singleton Pattern
version: 1.0.0
description: Detects singleton implementations that are not thread-safe and can result in multiple instances being created under concurrent access.
tags: [race-condition, concurrency, singleton, owasp-a06]
languages: [java, go, python, javascript, typescript]
severity: medium
confidence: low
cwe: [CWE-362]
owasp: [A06:2025]
---

# Thread-Unsafe Singleton

## Overview
Lazy-initialized singletons without thread synchronization can result in multiple instances being created when the class is initialized concurrently by multiple threads. This causes inconsistent state, duplicate resource consumption, and potential security issues if security-related objects (loggers, auth managers) are duplicated.

## Remediation
- Java: Use `enum`, `static initialization block`, or `synchronized` + `volatile`
- Go: Use `sync.Once`
- Python: Use module-level variable (thread-safe by GIL) or `threading.Lock`
