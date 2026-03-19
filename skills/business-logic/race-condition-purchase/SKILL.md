---
confidence: low
cwe:
    - CWE-362
description: Detects financial transaction code that can be exploited via concurrent requests to exceed balance limits or purchase limits.
languages:
    - javascript
    - typescript
    - python
    - php
    - java
    - go
    - kotlin
    - dart
    - zig
    - elixir
name: Race Condition in Purchase / Financial Transaction
owasp:
    - A01:2025
severity: high
tags:
    - business-logic
    - race-condition
    - concurrency
    - owasp-a01
version: 1.0.0
---

# Race Condition in Purchase / Financial Transaction

## Overview
A race condition in financial transactions allows attackers to send multiple concurrent requests to:
- Withdraw more funds than available balance
- Buy more items than stock allows
- Redeem the same voucher/coupon multiple times
- Use the same gift card balance repeatedly

This is a time-of-check/time-of-use (TOCTOU) vulnerability at the database level.

## Remediation
- Use database-level row locking (`SELECT ... FOR UPDATE`)
- Use optimistic locking with version counters
- Use atomic operations (`UPDATE balance SET balance = balance - amount WHERE balance >= amount`)
- Use idempotency keys for payment APIs
