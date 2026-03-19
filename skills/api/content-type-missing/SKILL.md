---
name: Missing or Incorrect Content-Type Validation
version: 1.0.0
description: Detects API endpoints that do not validate or enforce Content-Type headers, enabling MIME-type confusion and CSRF attacks.
tags: [api, content-type, csrf, owasp-a05]
languages: [javascript, typescript, python, go, java, php]
severity: medium
confidence: medium
cwe: [CWE-16]
owasp: [A05:2025]
---

# Missing Content-Type Validation

## Overview
APIs that accept requests without validating Content-Type are vulnerable to:
- **CSRF**: Browser forms submit as `application/x-www-form-urlencoded` which HTML forms can send cross-origin
- **Content confusion**: Unexpected parsing if body is mismatched with Content-Type
- **Polyglot attacks**: Content that is valid as multiple types

Modern CSRF protection relies on the browser's cross-origin restriction on JSON Content-Type.

## Remediation
- Validate Content-Type for all POST/PUT/PATCH endpoints
- Reject requests with unexpected Content-Type
- Use CSRF tokens even for JSON APIs as defense-in-depth
