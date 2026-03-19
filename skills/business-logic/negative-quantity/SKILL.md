---
name: Negative Quantity / Amount Manipulation
version: 1.0.0
description: Detects shopping cart and transaction logic that accepts negative quantities or amounts, enabling credit manipulation.
tags: [business-logic, ecommerce, owasp-a01]
languages: [javascript, typescript, python, php, java, go]
severity: high
confidence: medium
cwe: [CWE-840]
owasp: [A01:2025]
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
