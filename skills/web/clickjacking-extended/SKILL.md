---
confidence: medium
cwe:
    - CWE-1021
description: Detects advanced clickjacking vulnerabilities including UI redressing via transparent overlays and framebusting bypasses.
languages:
    - javascript
    - typescript
    - python
    - go
    - php
    - kotlin
    - dart
    - zig
    - elixir
name: Extended Clickjacking Detection
owasp:
    - A05:2025
severity: high
tags:
    - web
    - clickjacking
    - ui-redressing
    - owasp-a05
version: 1.0.0
---

# Extended Clickjacking Detection

## Overview
Beyond basic iframe embedding, clickjacking variants include:
- **Cursorjacking**: Replacing the browser cursor to mislead click position
- **Drag-and-drop jacking**: Tricking users into dragging content
- **Touchjacking**: Exploiting touch events on mobile
- **Framebusting bypass**: JavaScript framebusting that can be bypassed via `sandbox="allow-scripts"`

## Remediation
- Use CSP `frame-ancestors 'none'` — not bypassable by sandbox attribute
- Do not rely on JavaScript-only framebusting
