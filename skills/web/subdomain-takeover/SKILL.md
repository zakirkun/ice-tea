---
confidence: medium
cwe:
    - CWE-350
description: Detects CNAME records pointing to cloud services that may be dangling, and configurations referencing external services vulnerable to takeover.
languages:
    - javascript
    - typescript
    - yaml
    - generic
    - kotlin
    - dart
    - zig
    - elixir
name: Subdomain Takeover Risk
owasp:
    - A05:2025
severity: high
tags:
    - web
    - subdomain-takeover
    - dns
    - owasp-a05
version: 1.0.0
---

# Subdomain Takeover Risk

## Overview
Subdomain takeover occurs when a DNS CNAME points to a third-party service (GitHub Pages, Heroku, S3, Netlify) whose content has been removed. An attacker can claim the now-unclaimed resource on the third-party platform and host malicious content that browsers serve under the original company's subdomain.

## Remediation
- Regularly audit DNS records and remove CNAMEs for deprovisioned services
- Before deprovisioning cloud resources, remove DNS records first
- Monitor subdomain availability with tools like `can-i-take-over-xyz`
