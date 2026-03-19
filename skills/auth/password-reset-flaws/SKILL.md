---
name: Password Reset Vulnerabilities
version: 1.0.0
description: Detects insecure password reset implementations including predictable tokens, missing expiration, and host header injection in reset links.
tags: [auth, password-reset, owasp-a07]
languages: [javascript, typescript, python, go, java, php]
severity: high
confidence: medium
cwe: [CWE-640, CWE-838]
owasp: [A07:2025]
---

# Password Reset Vulnerabilities

## Overview
Password reset flows contain numerous potential vulnerabilities:
1. **Predictable tokens**: Reset tokens generated with Math.random() instead of CSPRNG
2. **No expiration**: Tokens valid indefinitely, allowing long-lived account takeover
3. **Host header injection**: Reset URL generated from Host header (phishing)
4. **Token reuse**: Token remains valid after use
5. **User enumeration**: Different responses for valid vs invalid email

## Remediation
- Use cryptographically random tokens (32+ bytes)
- Expire tokens after 15-60 minutes and after first use
- Generate reset URL from server configuration, not Host header
- Return the same response regardless of email existence
