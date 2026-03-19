---
name: SNMP Weak Community String
version: 1.0.0
description: Detects SNMP configurations using default or weak community strings that allow unauthorized network device management.
tags: [network, snmp, credentials, owasp-a07]
languages: [python, go, javascript, generic]
severity: high
confidence: high
cwe: [CWE-1392]
owasp: [A07:2025]
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
