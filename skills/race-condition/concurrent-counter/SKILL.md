---
name: Non-Atomic Counter Increment Race Condition
version: 1.0.0
description: Detects concurrent counter increments and decrements without atomic operations or mutex protection.
tags: [race-condition, concurrency, counter, owasp-a06]
languages: [go, java, javascript, typescript, c, cpp]
severity: medium
confidence: low
cwe: [CWE-362]
owasp: [A06:2025]
---

# Non-Atomic Counter Race Condition

## Overview
Incrementing/decrementing shared counters (`count++`, `balance += amount`) without atomic operations creates race conditions where concurrent updates lose some increments, leading to incorrect values. This affects rate limiters, inventory counters, and financial balances.

## Remediation
- Go: Use `sync/atomic` package (`atomic.AddInt64()`)
- Java: Use `AtomicInteger`/`AtomicLong`
- JavaScript: Use a single event loop (Node.js) or worker locks
- C/C++: Use `std::atomic<int>` or `_Atomic int`
