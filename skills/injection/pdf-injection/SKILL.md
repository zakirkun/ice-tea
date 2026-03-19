---
name: PDF / Document Injection
version: 1.0.0
description: Detects user input embedded in PDF or document generation without sanitization, enabling XSS through PDF rendering and data exfiltration.
tags: [injection, pdf, document, owasp-a03]
languages: [javascript, typescript, python, java, php]
severity: high
confidence: medium
cwe: [CWE-79]
owasp: [A03:2025]
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
