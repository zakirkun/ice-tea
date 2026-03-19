---
confidence: high
cwe:
    - CWE-1022
description: Detects anchor tags with target="_blank" without rel="noopener noreferrer", allowing opened pages to manipulate the opener.
languages:
    - javascript
    - typescript
    - php
    - python
    - generic
    - kotlin
    - dart
    - zig
    - elixir
name: Reverse Tabnapping
owasp:
    - A05:2025
severity: medium
tags:
    - web
    - tabnapping
    - xss
    - owasp-a05
version: 1.0.0
---

# Reverse Tabnapping

## Overview
When an `<a>` tag uses `target="_blank"` without `rel="noopener noreferrer"`, the opened page gets a reference to the opener window via `window.opener`. A malicious opened page can redirect the original tab to a phishing page (`window.opener.location = 'https://phishing.com'`).

## Remediation
Always add `rel="noopener noreferrer"` to links with `target="_blank"`:

**Vulnerable:**
```html
<a href="https://external.com" target="_blank">Visit</a>
```

**Safe:**
```html
<a href="https://external.com" target="_blank" rel="noopener noreferrer">Visit</a>
```
