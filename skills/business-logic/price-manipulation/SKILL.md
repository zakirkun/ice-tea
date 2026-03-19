---
name: Price Manipulation Vulnerability
version: 1.0.0
description: Detects e-commerce logic that trusts client-submitted prices instead of server-side calculation.
tags: [business-logic, ecommerce, owasp-a01]
languages: [javascript, typescript, python, php, java, go]
severity: high
confidence: medium
cwe: [CWE-840]
owasp: [A01:2025]
---

# Price Manipulation Vulnerability

## Overview
Price manipulation occurs when an application trusts the price submitted by the client (in request body or form fields) instead of calculating it server-side from the product catalog. Attackers can submit negative prices, zero prices, or arbitrarily low prices.

## Remediation
- NEVER trust client-submitted prices
- Always calculate price server-side from the product ID and current catalog
- Validate total = sum of (server_price × quantity) before processing payment
