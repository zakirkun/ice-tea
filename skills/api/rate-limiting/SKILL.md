---
confidence: low
cwe:
    - CWE-770
    - CWE-307
description: "Detects sensitive API endpoints (auth, password reset, OTP) without rate limiting middleware. Use when auditing API security, checking for brute force protection, or reviewing throttling configuration."
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
name: missing-api-rate-limiting
owasp:
    - A04:2025
severity: high
tags:
    - rate-limiting
    - api
    - dos
    - owasp-api4
version: 1.0.0
---

# Missing API Rate Limiting

Identifies sensitive endpoints (auth, password reset, OTP) that lack rate limiting middleware. See `patterns.yaml` for framework-specific detection rules.

## Detection Workflow

1. **Enumerate endpoints** — find route definitions for auth, reset, OTP, and data-listing handlers
2. **Check middleware** — verify rate limiter is applied (express-rate-limit, Flask-Limiter, or custom)
3. **Validate limits** — flag excessively high limits (>5000 requests/window) as ineffective
4. **Classify risk** — auth/OTP endpoints are critical; list endpoints are medium

### Vulnerable

```javascript
router.post('/api/reset-password', async (req, res) => { ... });
```

### Safe

```javascript
const limiter = rateLimit({ windowMs: 15 * 60 * 1000, max: 5 });
router.post('/api/reset-password', limiter, async (req, res) => { ... });
```
