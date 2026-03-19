---
name: Session Fixation
version: 1.0.0
description: Detects missing session regeneration after login, allowing session fixation attacks.
tags: [session, auth, fixation, owasp-a07]
languages: [php, python, javascript, java, go, ruby]
severity: high
confidence: medium
cwe: [CWE-384]
owasp: [A07:2025]
---

# Session Fixation

## Overview
Session fixation occurs when an application does not regenerate the session identifier after a successful login. An attacker can set a known session ID before authentication, and after the victim logs in, the attacker reuses the same session ID to gain authenticated access.

## Detection Strategy
Look for authentication flows (login functions) that do not call session regeneration functions before or after setting the authenticated user context.

Key patterns:
- Login handlers that set session user data without regenerating the session token
- Use of `session_start()` without a subsequent `session_regenerate_id(true)` in PHP
- Express.js `req.session.regenerate()` not called after login
- Django `cycle_key()` or `flush()` not called after `authenticate()`

## Remediation
Always regenerate the session ID after a successful authentication event.

**Vulnerable (PHP):**
```php
session_start();
if ($valid_login) {
    $_SESSION['user'] = $username; // no session_regenerate_id!
}
```

**Safe (PHP):**
```php
session_start();
if ($valid_login) {
    session_regenerate_id(true);
    $_SESSION['user'] = $username;
}
```

**Safe (Express.js):**
```js
req.session.regenerate((err) => {
    req.session.user = user;
    res.redirect('/dashboard');
});
```
