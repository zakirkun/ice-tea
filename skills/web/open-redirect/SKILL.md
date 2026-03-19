---
confidence: medium
cwe:
    - CWE-601
description: Detects HTTP redirects to user-controlled URLs, enabling phishing and server-side request forgery.
languages:
    - generic
    - kotlin
    - dart
    - zig
    - elixir
name: Unvalidated Redirects and Forwards
owasp:
    - A01:2025
severity: medium
tags:
    - redirect
    - phishing
    - web
    - owasp-a01
version: 1.0.0
---

# Unvalidated Redirects and Forwards

## Overview
Detects HTTP redirects to user-controlled URLs, enabling phishing and server-side request forgery.

## Remediation
Validate redirect URLs. Avoid using user input directly; map input to internal routing enums instead.
