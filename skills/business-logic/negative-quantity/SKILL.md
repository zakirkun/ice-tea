---
confidence: medium
cwe:
    - CWE-840
description: Detects shopping cart and transaction logic that accepts negative quantities or amounts, enabling credit manipulation.
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
name: Negative Quantity / Amount Manipulation
owasp:
    - A01:2025
severity: high
tags:
    - business-logic
    - ecommerce
    - owasp-a01
version: 1.0.0
---

# Negative Quantity / Amount Manipulation

## Overview
Without server-side validation that quantities and amounts must be positive, attackers can submit negative quantities to:
- Receive a refund without returning items
- Add negative-priced items to reduce total to zero or below
- Earn loyalty points by creating and cancelling negative transactions

## Remediation
- Validate all quantity and amount fields are strictly positive (> 0)
- Validate total order amount is positive before processing payment
- Use unsigned integer types where possible
