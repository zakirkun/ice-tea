---
name: MFA / 2FA Bypass Vulnerabilities
version: 1.0.0
description: Detects insecure multi-factor authentication implementations that can be bypassed.
tags: [mfa, 2fa, auth, bypass, owasp-a07]
languages: [javascript, typescript, python, go, java, php]
severity: critical
confidence: medium
cwe: [CWE-287, CWE-303]
owasp: [A07:2025]
---

# MFA / 2FA Bypass Vulnerabilities

## Overview
Poorly implemented MFA can be bypassed through various techniques:
1. **Client-side MFA check**: MFA validation happens in JavaScript that can be disabled
2. **Predictable OTP**: TOTP/OTP generated with Math.random() instead of cryptographically secure RNG
3. **Missing MFA enforcement**: Endpoints that allow direct access skipping the MFA step
4. **Long-lived OTP**: OTP codes valid for too long (> 5 minutes)
5. **No attempt limit on OTP**: Brute-forceable OTP codes

## Detection Strategy
- OTP generation using `Math.random()` or `rand()` instead of CSPRNG
- MFA verification that can be bypassed by manipulating the `mfa_verified` session variable
- Hardcoded backup codes in source

## Remediation
- Use TOTP (RFC 6238) with a cryptographically secure library like `speakeasy` or `pyotp`
- Validate MFA server-side, never client-side
- Limit OTP attempts (max 3-5 before requiring restart)
- Keep OTP validity window to 30-60 seconds

**Vulnerable (Node.js):**
```js
const otp = Math.floor(Math.random() * 1000000);
```

**Safe (Node.js):**
```js
const speakeasy = require('speakeasy');
const token = speakeasy.totp({ secret: user.mfa_secret, encoding: 'base32' });
```
