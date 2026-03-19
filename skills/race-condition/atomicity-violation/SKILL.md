---
name: Atomicity Violation
version: 1.0.0
description: Detects code sequences that must be atomic but are not protected by synchronization, allowing interleaving that violates invariants.
tags: [race-condition, concurrency, atomicity, owasp-a06]
languages: [go, java, javascript, typescript]
severity: high
confidence: low
cwe: [CWE-362]
owasp: [A06:2025]
---

# Atomicity Violation

## Overview
Atomicity violations occur when a sequence of operations that must execute as a unit can be interleaved by other threads. Unlike simple counter races, these involve multi-step operations where the intermediate state is invalid.

Example: Check-then-act on user balance, file existence, or session state.

## Remediation
- Wrap multi-step operations in mutex/lock
- Use database transactions for multi-step data operations
- Use compare-and-swap (CAS) operations for lock-free atomic sequences
