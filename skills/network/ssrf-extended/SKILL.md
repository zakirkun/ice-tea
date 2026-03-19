---
confidence: high
cwe:
    - CWE-918
description: Detects SSRF vulnerabilities targeting cloud metadata services, internal networks, and non-HTTP protocols.
languages:
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
name: Extended SSRF Detection (Cloud Metadata & Protocol Exploits)
owasp:
    - A10:2025
severity: critical
tags:
    - ssrf
    - network
    - cloud
    - owasp-a10
version: 1.0.0
---

# Extended SSRF Detection (Cloud Metadata & Protocol Exploits)

## Overview
Server-Side Request Forgery (SSRF) attacks using specific targets and protocols:

1. **Cloud metadata endpoints**:
   - AWS: `http://169.254.169.254/latest/meta-data/`
   - GCP: `http://metadata.google.internal/`
   - Azure: `http://169.254.169.254/metadata/`

2. **Non-HTTP protocols**:
   - `file://`: Read local files
   - `gopher://`: Binary protocol for internal service exploitation
   - `dict://`: Info leak via Redis/Memcached
   - `ftp://`: Internal FTP access

3. **IPv6 bypass**: `http://[::1]/` to reach localhost
4. **URL encoding bypass**: `http://127.0.0.1%2F/` or `http://0x7f000001/`

## Remediation
- Allowlist permitted URL schemes (only `https://`)
- Allowlist permitted destination IPs/domains
- Use a dedicated HTTP client proxy that enforces policies
- Disable `file://`, `gopher://`, `dict://` in HTTP clients
