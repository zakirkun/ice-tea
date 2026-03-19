---
name: Business Logic Privilege Escalation
version: 1.0.0
description: Detects application logic that allows users to escalate their own privileges or roles through manipulated requests.
tags: [business-logic, privilege-escalation, auth, owasp-a01]
languages: [javascript, typescript, python, php, java, go]
severity: critical
confidence: medium
cwe: [CWE-269]
owasp: [A01:2025]
---

# Business Logic Privilege Escalation

## Overview
Privilege escalation through business logic occurs when:
- Role is accepted from client request during registration/update
- Admin endpoints rely on user-submitted role parameter
- User can invite themselves to higher-privilege groups
- Role update logic does not verify the caller's permission to grant that role

## Remediation
- Never accept role/permission from client
- Role assignment must be done by admins only, verified server-side
- Log all privilege change events
