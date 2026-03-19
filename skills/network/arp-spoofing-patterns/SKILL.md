---
confidence: low
cwe:
    - CWE-923
description: Detects network configurations and code patterns that are vulnerable to ARP spoofing attacks.
languages:
    - python
    - generic
    - kotlin
    - dart
    - zig
    - elixir
name: ARP Spoofing Defense Missing
owasp:
    - A02:2025
severity: medium
tags:
    - network
    - arp
    - mitm
    - owasp-a02
version: 1.0.0
---

# ARP Spoofing Defense Missing

## Overview
ARP spoofing allows an attacker on the same network segment to intercept traffic by poisoning ARP caches, causing victims to send traffic to the attacker instead of the gateway. This enables:
- MITM attacks on HTTP traffic
- Credential theft on cleartext protocols
- Session hijacking

Code that sends sensitive data without verifying end-to-end integrity (TLS, signatures) is vulnerable if ARP spoofing occurs.

## Remediation
- Use TLS for all sensitive communication (prevents data theft even with MITM)
- Configure static ARP entries for critical servers
- Enable dynamic ARP inspection on managed switches
- Use VPN for all inter-server communication on shared networks
