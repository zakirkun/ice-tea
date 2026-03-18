---
name: Cross-Site Scripting (XSS) Detection
version: 1.0.0
description: Detects DOM-based Cross-Site Scripting (XSS) vulnerabilities in JavaScript
tags: [xss, web, injection, owasp-a03]
languages: [javascript, typescript]
severity: high
confidence: medium
cwe: [CWE-79]
owasp: [A03:2025]
---

# Cross-Site Scripting (XSS)

## Overview
XSS vulnerabilities occur when an application includes untrusted data in a web page without proper validation or escaping.

## Detection Strategy
This SKILL specifically looks for dangerous DOM manipulations in frontend code where user-controlled input might be executed as script.

Sinks:
- `innerHTML` assignment
- `document.write()`
- `eval()`
- `setTimeout()` with string evaluation

## Remediation
Use safer alternatives like `textContent` or `innerText` instead of `innerHTML`. Use DOMPurify if HTML insertion is strictly required.
