---
name: Internal API Documentation Exposure
version: 1.0.0
description: Detects internal API documentation, admin endpoints, and developer tools accidentally exposed in production.
tags: [api, information-disclosure, owasp-a05]
languages: [javascript, typescript, python, go, java, php]
severity: medium
confidence: high
cwe: [CWE-200]
owasp: [A05:2025]
---

# Internal API Documentation Exposure

## Overview
Development and debugging endpoints left enabled in production:
- `/actuator` (Spring Boot) — exposes heap dumps, metrics, environment
- `/debug` — application debugging interface
- `/admin/docs` — internal API documentation
- `/phpinfo.php` — PHP configuration disclosure
- `/__debug__` — Django debug toolbar
- `/robots.txt` referencing sensitive paths

## Remediation
- Disable or require authentication for all admin and debug endpoints
- Use environment guards: `if (process.env.NODE_ENV !== 'production')`
- Configure framework-specific settings to disable development tools in production
