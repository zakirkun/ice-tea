---
name: Coupon and Discount Code Abuse
version: 1.0.0
description: Detects discount/coupon code logic without per-user usage limits, allowing unlimited stacking or reuse.
tags: [business-logic, coupon, ecommerce, owasp-a01]
languages: [javascript, typescript, python, php, java, go]
severity: medium
confidence: low
cwe: [CWE-840]
owasp: [A01:2025]
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
