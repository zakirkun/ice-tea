---
confidence: medium
cwe:
    - CWE-319
description: Detects unencrypted HTTP connections for transmitting sensitive data, credentials, or API calls.
languages:
    - javascript
    - typescript
    - python
    - go
    - java
    - php
    - ruby
    - kotlin
    - dart
    - zig
    - elixir
name: Cleartext Network Traffic
owasp:
    - A02:2025
severity: high
tags:
    - network
    - http
    - cleartext
    - owasp-a02
version: 1.0.0
---

# Cleartext Network Traffic

## Overview
Sending data over HTTP (not HTTPS) exposes it to eavesdropping, especially on:
- Public WiFi networks
- ISP-level interception
- Corporate proxies
- Nation-state adversaries

Sensitive operations that must use HTTPS:
- Login and authentication
- API calls with tokens or session cookies
- Payment and financial transactions
- Personal data transmission

## Detection Strategy
- HTTP URLs hardcoded for API endpoints
- HTTP clients configured without TLS
- Mixed content: HTTPS page loading resources over HTTP

## Remediation
- Use HTTPS for all network communication
- Enable HSTS to prevent downgrade attacks
- Add HTTP→HTTPS redirect at server level
