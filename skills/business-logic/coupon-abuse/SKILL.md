---
confidence: low
cwe:
    - CWE-840
description: Detects discount/coupon code logic without per-user usage limits, allowing unlimited stacking or reuse.
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
name: Coupon and Discount Code Abuse
owasp:
    - A01:2025
severity: medium
tags:
    - business-logic
    - coupon
    - ecommerce
    - owasp-a01
version: 1.0.0
---

# Coupon and Discount Code Abuse

## Overview
Coupon systems without proper constraints allow:
- Reusing a single-use coupon multiple times (no per-user tracking)
- Stacking multiple coupons beyond allowed limits
- Applying percentage-based coupons to items below minimum order value
- Race-condition reuse: applying one coupon from multiple tabs simultaneously

## Remediation
- Track coupon usage per user in the database
- Use database transactions with row-level locking for redemption
- Validate minimum order value requirements server-side
- Rate limit coupon attempts per user
