---
confidence: high
cwe:
    - CWE-1392
description: Detects SNMP configurations using default or weak community strings that allow unauthorized network device management.
languages:
    - python
    - go
    - javascript
    - generic
    - kotlin
    - dart
    - zig
    - elixir
name: SNMP Weak Community String
owasp:
    - A07:2025
severity: high
tags:
    - network
    - snmp
    - credentials
    - owasp-a07
version: 1.0.0
---

# SNMP Weak Community String

## Overview
SNMP (Simple Network Management Protocol) uses community strings as authentication for v1 and v2c. Default strings `public` (read) and `private` (write) are universally known and allow:
- Reading entire MIB (network topology, device info, traffic stats)
- Writing to network configuration (v2c write access)
- Gaining network intelligence for lateral movement

## Remediation
- Use SNMPv3 with authentication and encryption
- If v2c required, use strong random community strings
- Block SNMP from external interfaces
- Apply ACLs to restrict SNMP access to management stations
