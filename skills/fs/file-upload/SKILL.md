---
confidence: medium
cwe:
    - CWE-434
description: Detects file upload handlers that lack extension validation, MIME type checking, or store files in web-accessible directories.
languages:
    - php
    - python
    - javascript
    - typescript
    - go
    - java
    - kotlin
    - dart
    - zig
    - elixir
name: Insecure File Upload
owasp:
    - A04:2025
severity: critical
tags:
    - file-upload
    - web
    - rce
    - owasp-a04
version: 1.0.0
---

# Insecure File Upload

## Overview
Unrestricted file upload allows attackers to upload web shells (PHP, JSP, ASPX) or other malicious files that can then be executed by the server. This often leads to Remote Code Execution.

Attack scenarios:
1. Upload `shell.php` — server executes it if stored in webroot
2. Upload a file with double extension: `shell.php.jpg` — some servers still execute
3. Upload SVG with embedded XSS — client-side attack
4. Upload oversized file — DoS via disk exhaustion

## Detection Strategy
- `move_uploaded_file()` without extension whitelist check
- `multer` storage without file type validation
- Upload directory inside webroot
- Trust of `Content-Type` header alone (user-controlled)

## Remediation
- Whitelist allowed extensions (not blacklist)
- Validate MIME type using file content (magic bytes), not only headers
- Store uploaded files outside webroot
- Rename uploaded files to random UUIDs
- Scan with antivirus for high-risk applications
- Set file size limits
