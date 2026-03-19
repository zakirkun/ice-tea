---
confidence: medium
cwe:
    - CWE-79
description: Detects user input embedded in PDF or document generation without sanitization, enabling XSS through PDF rendering and data exfiltration.
languages:
    - javascript
    - typescript
    - python
    - java
    - php
    - kotlin
    - dart
    - zig
    - elixir
name: PDF / Document Injection
owasp:
    - A03:2025
severity: high
tags:
    - injection
    - pdf
    - document
    - owasp-a03
version: 1.0.0
---

# PDF / Document Injection

## Overview
PDF generators that include user-controlled HTML or data can be exploited:
1. **HTML to PDF injection**: Injecting `<script>` or `<link>` tags that read local files via PDF rendering engine (wkhtmltopdf, Puppeteer)
2. **JavaScript in PDF**: PDF actions that execute JavaScript in Adobe Reader
3. **Formula injection in generated Excel**: See CSV Injection

wkhtmltopdf and similar tools can read local files: `<img src="file:///etc/passwd">` in the input HTML.

## Remediation
- HTML-encode all user input before HTML-to-PDF conversion
- Use `--no-local-file-access` flag with wkhtmltopdf
- Use sandboxed PDF generation with no filesystem access
