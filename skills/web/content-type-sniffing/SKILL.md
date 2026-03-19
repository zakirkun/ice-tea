---
name: Content Type Sniffing Vulnerability
version: 1.0.0
description: Detects missing X-Content-Type-Options header and incorrect Content-Type that allows MIME sniffing attacks.
tags: [web, content-type, mime-sniffing, owasp-a05]
languages: [javascript, typescript, python, go, java, php]
severity: medium
confidence: medium
cwe: [CWE-79, CWE-16]
owasp: [A05:2025]
---

# Content Type Sniffing

## Overview
Browsers that MIME-sniff responses can interpret uploaded files differently from the declared content type. An attacker uploads an HTML file disguised as an image, and the browser sniffs it as HTML and executes the embedded JavaScript.

`X-Content-Type-Options: nosniff` prevents this sniffing behavior.

## Remediation
- Set `X-Content-Type-Options: nosniff` on all responses
- Always declare the correct Content-Type for served files
- Do not serve user-uploaded content from the same origin as the application
