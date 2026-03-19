---
name: Memory Race Condition
version: 1.0.0
description: Detects concurrent memory access without proper synchronization in C/C++ and Go programs.
tags: [memory, race-condition, concurrency, c, cpp, owasp-a06]
languages: [c, cpp, go]
severity: high
confidence: low
cwe: [CWE-362]
owasp: [A06:2025]
---

# Memory Race Condition

## Overview
Memory race conditions occur when two threads access shared memory simultaneously without synchronization, and at least one access is a write. This leads to:
- Data corruption (non-deterministic results)
- Security bypass (auth state race)
- Potential exploitation via crafted timing

## Remediation
- Use mutex/lock primitives for all shared data access
- Use atomic operations for simple counters and flags
- Design for lock-free or copy-on-write patterns where performance matters
- Use Go race detector: `go test -race`
