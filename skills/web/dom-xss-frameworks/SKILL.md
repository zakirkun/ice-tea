---
name: Framework-Specific DOM XSS
version: 1.0.0
description: Detects the use of dangerous API methods in modern frontend frameworks (React, Vue, Angular) that can lead to Cross-Site Scripting (XSS).
tags: [xss, frontend, react, vue, angular, owasp-a03]
languages: [generic]
severity: high
confidence: medium
cwe: [CWE-79]
owasp: [A03:2025]
---

# Framework-Specific DOM XSS

## Overview
Detects the use of dangerous API methods in modern frontend frameworks (React, Vue, Angular) that can lead to Cross-Site Scripting (XSS).

## Remediation
Avoid using raw HTML injection APIs. If inevitable, sanitize input using a library like DOMPurify.
