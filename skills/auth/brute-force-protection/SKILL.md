---
name: Missing Brute Force Protection
version: 1.0.0
description: Detects login and authentication endpoints that lack rate limiting, account lockout, or CAPTCHA protection.
tags: [brute-force, auth, rate-limiting, owasp-a07]
languages: [javascript, typescript, python, go, java, php, ruby]
severity: high
confidence: low
cwe: [CWE-307]
owasp: [A07:2025]
---

# Missing Brute Force Protection

## Overview
Authentication endpoints without rate limiting allow attackers to make unlimited login attempts, enabling:
- **Credential stuffing**: Testing breached username/password combinations
- **Password spraying**: Testing common passwords across many accounts
- **Dictionary attacks**: Exhaustively guessing passwords

## Detection Strategy
Identify login/authentication route handlers that do not implement:
- Rate limiting middleware (express-rate-limit, Flask-Limiter, etc.)
- Account lockout after N failed attempts
- CAPTCHA verification

## Remediation
- Add rate limiting to all authentication endpoints (e.g., 5 attempts per 15 minutes per IP)
- Implement progressive delays or account lockout after repeated failures
- Consider CAPTCHA for high-value applications
- Use fail2ban or similar at the infrastructure level

**Safe (Express.js):**
```js
const rateLimit = require('express-rate-limit');
const loginLimiter = rateLimit({
    windowMs: 15 * 60 * 1000,
    max: 5,
    message: 'Too many login attempts'
});
app.post('/login', loginLimiter, loginHandler);
```
