---
confidence: medium
cwe:
    - CWE-912
description: Detects internal IP addresses, localhost references, and hardcoded hostnames that indicate misconfiguration or information disclosure.
languages:
    - generic
    - javascript
    - typescript
    - python
    - go
    - java
    - php
    - kotlin
    - dart
    - zig
    - elixir
name: Hardcoded IP Addresses and Hostnames
owasp:
    - A05:2025
severity: medium
tags:
    - hardcoded-ips
    - infra
    - configuration
    - owasp-a05
version: 1.0.0
---

# Hardcoded IP Addresses and Hostnames

## Overview
Hardcoded IP addresses and internal hostnames in source code can:
- Expose internal network topology to attackers who gain code access
- Prevent deployment flexibility (tied to a specific environment)
- Point to services with weaker security than production (dev/staging servers)
- Enable SSRF if an attacker can influence which IP is connected to

## Detection Strategy
- Internal IP ranges: `10.x.x.x`, `172.16-31.x.x`, `192.168.x.x`
- Localhost with non-standard ports: `127.0.0.1:8080`
- Cloud provider metadata endpoints: `169.254.169.254`
- Hardcoded database hostnames in non-config files

## Remediation
Move all hostnames and IPs to environment variables or configuration files that are not committed to source control.
