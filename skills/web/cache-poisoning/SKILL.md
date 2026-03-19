---
confidence: medium
cwe:
    - CWE-524
description: Detects applications that include unvalidated request headers or parameters in cached responses, enabling cache poisoning attacks.
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
name: Web Cache Poisoning
owasp:
    - A05:2025
severity: high
tags:
    - web
    - cache-poisoning
    - owasp-a05
version: 1.0.0
---

# Web Cache Poisoning

## Overview
Cache poisoning occurs when an attacker can inject malicious content into a shared cache (CDN, reverse proxy, application cache) that is then served to other users. Unkeyed headers (headers that affect the response but are not included in the cache key) are the primary attack vector.

## Remediation
- Configure cache key to include all headers that affect the response
- Validate and sanitize all header values used in responses
- Use `Vary` header to include relevant headers in the cache key
- Disable caching for responses that include user-controlled content
