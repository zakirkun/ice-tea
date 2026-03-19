---
confidence: low
cwe:
    - CWE-362
description: Detects concurrent counter increments and decrements without atomic operations or mutex protection.
languages:
    - go
    - java
    - javascript
    - typescript
    - c
    - cpp
    - kotlin
    - dart
    - zig
    - elixir
name: Non-Atomic Counter Increment Race Condition
owasp:
    - A06:2025
severity: medium
tags:
    - race-condition
    - concurrency
    - counter
    - owasp-a06
version: 1.0.0
---

# Non-Atomic Counter Race Condition

## Overview
Incrementing/decrementing shared counters (`count++`, `balance += amount`) without atomic operations creates race conditions where concurrent updates lose some increments, leading to incorrect values. This affects rate limiters, inventory counters, and financial balances.

## Remediation
- Go: Use `sync/atomic` package (`atomic.AddInt64()`)
- Java: Use `AtomicInteger`/`AtomicLong`
- JavaScript: Use a single event loop (Node.js) or worker locks
- C/C++: Use `std::atomic<int>` or `_Atomic int`
