---
name: Double-Checked Locking Anti-Pattern
version: 1.0.0
description: Detects broken double-checked locking implementations that create race conditions in singleton initialization.
tags: [race-condition, concurrency, java, go, owasp-a06]
languages: [java, go, c, cpp]
severity: high
confidence: medium
cwe: [CWE-609]
owasp: [A06:2025]
---

# Double-Checked Locking Race Condition

## Overview
Double-checked locking is a common but often incorrectly implemented pattern for lazy singleton initialization. Without proper memory ordering or volatile declarations, the second check can observe a partially-initialized object due to instruction reordering.

In Java pre-5.0 (without `volatile`), this is broken. In Go, use `sync.Once` instead.

## Remediation
- Java: Declare singleton field `volatile`
- Go: Use `sync.Once` for lazy initialization
- C++: Use `std::call_once` or `std::atomic`
