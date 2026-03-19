---
confidence: medium
cwe:
    - CWE-319
description: Detects network configurations that fail to apply security controls to IPv6, allowing bypass via IPv6 tunneling.
languages:
    - generic
    - python
    - javascript
    - go
    - kotlin
    - dart
    - zig
    - elixir
name: IPv6 Tunneling and Misconfiguration
owasp:
    - A05:2025
severity: medium
tags:
    - network
    - ipv6
    - security-bypass
    - owasp-a05
version: 1.0.0
---

# IPv6 Tunneling and Misconfiguration

## Overview
Security controls applied only to IPv4 can be bypassed via IPv6:
- Firewalls blocking IPv4 but allowing all IPv6
- SSRF filters blocking `127.0.0.1` but not `::1` (IPv6 loopback)
- Rate limiters keyed on IPv4 that miss IPv6 addresses
- DNS rebinding via IPv6 link-local addresses

## Remediation
- Apply identical security rules to IPv4 and IPv6
- Block `::1` and IPv6 link-local ranges in SSRF filters
- Test security controls explicitly with IPv6 addresses
