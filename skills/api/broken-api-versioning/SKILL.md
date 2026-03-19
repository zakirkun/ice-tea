---
name: Broken API Versioning Security
version: 1.0.0
description: Detects API versioning patterns where newer security controls do not apply to older API versions still in use.
tags: [api, versioning, owasp-api9]
languages: [javascript, typescript, python, go, java, php]
severity: high
confidence: medium
cwe: [CWE-1059]
owasp: [A09:2025]
---

# Broken API Versioning Security

## Overview
When APIs evolve through versions, older versions often lack security controls added in newer versions:
- Authentication may be optional in v1 but required in v2
- Input validation may differ between versions
- Rate limiting may only apply to newer API versions
- Deprecated endpoints may expose functionality removed for security reasons

## Remediation
- Apply identical security controls to ALL active API versions
- Use shared security middleware that applies to all version routes
- Actively sunset old API versions with deprecation notices
- Monitor and alert on old API version usage
