---
name: HTTP Response Splitting
version: 1.0.0
description: Detects HTTP response splitting via CRLF injection in response headers, enabling cache poisoning and XSS.
tags: [crlf, http-response-splitting, web, owasp-a03]
languages: [java, python, javascript, typescript, php, go, ruby]
severity: high
confidence: medium
cwe: [CWE-113]
owasp: [A03:2025]
---

# HTTP Response Splitting

## Overview
HTTP Response Splitting is a more severe form of header injection. By injecting `\r\n` (CRLF) sequences into response headers, an attacker can:
- Inject arbitrary HTTP headers
- Add a second HTTP response body (response splitting)
- Poison shared caches (CDNs, proxies) with malicious content
- Execute XSS by injecting a fake HTML body

## Detection Strategy
Look for unvalidated user input placed in any HTTP response header, particularly `Location`, `Set-Cookie`, and `Content-Type`.

## Remediation
- Reject input containing `\r` or `\n` characters
- Use framework APIs that automatically sanitize header values
- Apply output encoding before placing user data in headers
