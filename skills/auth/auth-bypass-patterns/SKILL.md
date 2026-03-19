---
name: Authentication Bypass Patterns
version: 1.0.0
description: Detects common authentication bypass patterns including type juggling, SQL truncation, and logic flaws.
tags: [auth, bypass, owasp-a07]
languages: [javascript, typescript, python, php, java, go]
severity: critical
confidence: medium
cwe: [CWE-287]
owasp: [A07:2025]
---

# Authentication Bypass Patterns

## Overview
Authentication bypasses exploit logical flaws:
1. **PHP type juggling**: `"0e1234..." == "0e5678..."` (magic hash)
2. **SQL truncation**: Long username overflows DB field to match existing user
3. **NULL password bypass**: Some DBs accept NULL as matching any hash
4. **Mass assignment bypass**: Posting `isAdmin: true` with registration data

## Remediation
- Use strict comparison (`===` not `==` in PHP/JS)
- Validate input length against database field limits
- Use parameterized queries everywhere
- Never trust user-submitted role or privilege fields
