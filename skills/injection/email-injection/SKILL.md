---
confidence: high
cwe:
    - CWE-93
description: Detects email sending functions that include user input in email headers without CRLF stripping, enabling spam relay and header manipulation.
languages:
    - python
    - php
    - javascript
    - typescript
    - java
    - ruby
    - kotlin
    - dart
    - zig
    - elixir
name: Email Header Injection
owasp:
    - A03:2025
severity: high
tags:
    - injection
    - email
    - owasp-a03
version: 1.0.0
---

# Email Header Injection

## Overview
Email header injection (also known as SMTP header injection) allows attackers to inject additional `To:`, `CC:`, `BCC:` headers or modify the message by inserting CRLF sequences (`\r\n`) into email fields. This turns the application into a spam relay.

Attack: Setting the "From name" to `victim@example.com\r\nBCC: spam@list.com` adds a blind carbon copy to all sent emails.

## Remediation
- Strip `\r`, `\n` from all user-supplied email header values
- Use email library functions that automatically prevent injection
- Never concatenate user input directly into SMTP headers
