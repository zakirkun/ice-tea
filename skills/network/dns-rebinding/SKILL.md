---
name: DNS Rebinding Vulnerability
version: 1.0.0
description: Detects server-side host validation logic vulnerable to DNS rebinding attacks, allowing attackers to bypass IP-based restrictions.
tags: [dns-rebinding, network, ssrf, owasp-a10]
languages: [javascript, typescript, python, go, java]
severity: high
confidence: medium
cwe: [CWE-350]
owasp: [A10:2025]
---

# DNS Rebinding Vulnerability

## Overview
DNS rebinding allows an attacker to bypass Same-Origin Policy and IP allowlists. The attack works by:
1. Victim visits attacker's page at `evil.com`
2. Attacker's DNS TTL expires; `evil.com` resolves to `127.0.0.1`
3. JavaScript makes requests to `evil.com` which now reach the victim's localhost services
4. Services that only check the IP (not the Host header) are bypassed

Vulnerable applications that check `Host` header or resolve hostname for validation without proper binding.

## Detection Strategy
- URL validation that checks hostname string rather than resolved IP
- HTTP servers binding to all interfaces without Host header validation
- SSRF protection based on DNS resolution rather than allowlist
