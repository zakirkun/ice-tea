---
name: Debug Mode Enabled in Production
version: 1.0.0
description: Detects debug mode enabled in production configurations, exposing stack traces, internal paths, and sensitive configuration details.
tags: [debug, infra, configuration, owasp-a05]
languages: [python, javascript, typescript, go, java, php, ruby]
severity: high
confidence: high
cwe: [CWE-215, CWE-209]
owasp: [A05:2025]
---

# Debug Mode Enabled in Production

## Overview
Debug mode in web frameworks exposes critical information to attackers:
- **Flask DEBUG=True**: Interactive debugger in browser, allows arbitrary code execution
- **Django DEBUG=True**: Full stack traces with local variable values shown to users
- **Node.js** with verbose error logging: Internal file paths, stack traces
- **Spring Boot actuator**: `/actuator/env`, `/actuator/heapdump` exposed

## Remediation
- Set `DEBUG=False` in all production configurations
- Use environment variables to control debug settings
- Implement a custom error handler that returns generic error messages
- Disable development-only actuator endpoints in production

**Vulnerable (Flask):**
```python
app.run(debug=True, host='0.0.0.0')  # Never in production!
```
