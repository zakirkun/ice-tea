---
confidence: medium
cwe:
    - CWE-209
description: Detects stack traces, exception details, and internal error information sent in HTTP responses to end users.
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
name: Verbose Error Messages Exposed to Clients
owasp:
    - A05:2025
severity: medium
tags:
    - error-handling
    - logging
    - information-disclosure
    - owasp-a05
version: 1.0.0
---

# Verbose Error Messages Exposed to Clients

## Overview
Exposing detailed error information to end users reveals:
- **Stack traces**: Internal file paths, function names, code structure
- **Database errors**: Table names, column names, SQL queries
- **Framework errors**: Version information, configuration details
- **Exception messages**: Business logic and data structure hints

This information directly aids attackers in crafting more targeted attacks.

## Remediation
- Return generic error messages to clients (`Internal Server Error`, `Something went wrong`)
- Log detailed error information server-side for debugging
- Use custom error handlers in your framework
- Set appropriate HTTP status codes without leaking implementation details

**Vulnerable:**
```js
app.use((err, req, res, next) => {
    res.status(500).json({ error: err.stack }); // Exposes stack trace!
});
```

**Safe:**
```js
app.use((err, req, res, next) => {
    logger.error(err.stack); // Log internally
    res.status(500).json({ error: 'Internal Server Error' });
});
```
