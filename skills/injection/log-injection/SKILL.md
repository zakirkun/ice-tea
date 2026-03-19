---
name: Log Injection
version: 1.0.0
description: Detects user-controlled input written to log files without sanitization, enabling log forging and potential log viewer attacks.
tags: [log-injection, logging, injection, owasp-a09]
languages: [javascript, typescript, python, go, java, php, ruby]
severity: medium
confidence: medium
cwe: [CWE-117]
owasp: [A09:2025]
---

# Log Injection

## Overview
Log injection occurs when user-supplied data is written to log files without sanitizing newline characters. Attackers can:
- Forge log entries to cover their tracks or frame innocent users
- Inject fake log lines to mislead incident responders
- Exploit log viewer vulnerabilities (XSS in web-based log viewers)
- In some cases, exploit log4j-style lookup injection (Log4Shell)

## Detection Strategy
Look for logging calls that directly include user-supplied request parameters.

## Remediation
- Strip or encode `\r`, `\n`, and ANSI escape sequences from values before logging
- Use structured logging (JSON) which prevents newline-based injection
- Log the value in a structured field, not interpolated into a message string

**Vulnerable (Node.js):**
```js
const username = req.body.username;
logger.info(`User login attempt: ${username}`); // injection via \n
```

**Safe (Node.js):**
```js
logger.info('User login attempt', { username: username.replace(/[\r\n]/g, '') });
```
