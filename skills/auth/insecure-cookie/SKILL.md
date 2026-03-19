---
confidence: high
cwe:
    - CWE-614
    - CWE-1004
description: Detects cookies set without Secure, HttpOnly, or SameSite attributes.
languages:
    - php
    - python
    - javascript
    - java
    - go
    - ruby
    - kotlin
    - dart
    - zig
    - elixir
name: Insecure Cookie Configuration
owasp:
    - A05:2025
severity: medium
tags:
    - cookie
    - auth
    - session
    - owasp-a05
version: 1.0.0
---

# Insecure Cookie Configuration

## Overview
Cookies that store session tokens or authentication data must be configured with security attributes to prevent theft and CSRF attacks:
- **Secure**: Cookie only sent over HTTPS
- **HttpOnly**: Cookie inaccessible to JavaScript (prevents XSS token theft)
- **SameSite**: Prevents cross-site request forgery (Strict or Lax)

Missing any of these attributes expands the attack surface.

## Detection Strategy
Look for `Set-Cookie` headers or cookie-setting function calls that omit one or more of the critical security flags.

## Remediation
Always set session cookies with all three security attributes.

**Vulnerable (Go):**
```go
http.SetCookie(w, &http.Cookie{
    Name:  "session",
    Value: token,
})
```

**Safe (Go):**
```go
http.SetCookie(w, &http.Cookie{
    Name:     "session",
    Value:    token,
    Secure:   true,
    HttpOnly: true,
    SameSite: http.SameSiteStrictMode,
})
```
