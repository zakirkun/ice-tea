---
name: Server-Side Include (SSI) Injection
version: 1.0.0
description: Detects user-controlled input reflected in pages processed by SSI, enabling file disclosure and command execution.
tags: [injection, ssi, rce, owasp-a03]
languages: [php, generic]
severity: critical
confidence: medium
cwe: [CWE-97]
owasp: [A03:2025]
---

# Server-Side Include Injection

## Overview
Server-Side Includes (SSI) are directives processed by web servers (Apache, Nginx) before serving files. When user input is reflected in SSI-processed files, attackers can inject directives:
- `<!--#exec cmd="id"-->` — executes arbitrary commands
- `<!--#include file="/etc/passwd"-->` — reads arbitrary files

## Remediation
- Disable SSI processing for user-controlled content
- HTML-encode all user input before including in server-processed pages
- Use Content Security Policy to prevent SSI directive execution
