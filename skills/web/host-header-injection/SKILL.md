---
confidence: medium
cwe:
    - CWE-113
description: Detects applications that use the HTTP Host header for URL generation, password reset links, or routing without validation.
languages:
    - javascript
    - typescript
    - python
    - php
    - java
    - go
    - ruby
    - kotlin
    - dart
    - zig
    - elixir
name: Host Header Injection
owasp:
    - A01:2025
severity: high
tags:
    - web
    - host-header
    - injection
    - owasp-a01
version: 1.0.0
---

# Host Header Injection

## Overview
Applications that trust the `Host` header for generating links (password reset emails, canonical URLs) are vulnerable to Host Header Injection. Attackers set a malicious `Host:` header, causing the app to generate links pointing to the attacker's domain. The victim clicks a password reset link containing an attacker-controlled URL.

## Remediation
- Configure your web application with an explicit trusted host list
- Use Django's `ALLOWED_HOSTS`, Laravel's `config/app.php` URL, or server-side canonical URL
- Validate the `Host` header against an allowlist before using it
