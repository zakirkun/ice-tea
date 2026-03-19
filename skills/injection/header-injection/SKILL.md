---
confidence: medium
cwe:
    - CWE-113
description: Detects user-controlled data written into HTTP response headers without CRLF stripping, enabling header injection and response splitting.
languages:
    - java
    - python
    - javascript
    - typescript
    - php
    - go
    - ruby
    - kotlin
    - dart
    - zig
    - elixir
name: HTTP Header Injection / CRLF Injection
owasp:
    - A03:2025
severity: high
tags:
    - header-injection
    - crlf
    - http
    - owasp-a03
version: 1.0.0
---

# HTTP Header Injection / CRLF Injection

## Overview
HTTP response header injection occurs when user-controlled data is placed in HTTP response headers without stripping carriage return (`\r`, `%0d`) and newline (`\n`, `%0a`) characters. This enables:
- **Response splitting**: Injecting a fake second HTTP response
- **XSS via header**: Injecting `Set-Cookie` headers with malicious cookies
- **Cache poisoning**: Poisoning shared proxies and CDN caches
- **Open redirect**: Via `Location` header manipulation

## Detection Strategy
Look for response headers that include user input from request parameters or paths.

## Remediation
- Validate and sanitize all user input before placing in headers
- Strip CRLF characters (`\r\n`) from any value placed in a header
- Use framework's built-in header sanitization

**Vulnerable (Python):**
```python
redirect_url = request.args.get('url')
response = make_response('', 302)
response.headers['Location'] = redirect_url  # CRLF injection!
```

**Safe (Python):**
```python
import re
redirect_url = re.sub(r'[\r\n]', '', request.args.get('url', ''))
```
