---
confidence: high
cwe:
    - CWE-532
description: Detects sensitive data passed to NSLog, print, or os_log that ends up in device logs accessible via Xcode/Console.
languages:
    - generic
    - kotlin
    - dart
    - zig
    - elixir
name: iOS Sensitive Data in NSLog / os_log
owasp:
    - A09:2025
severity: high
tags:
    - ios
    - mobile
    - logging
    - owasp-m4
version: 1.0.0
---

# iOS Sensitive Data in Logs

## Overview
iOS device logs (accessible via Xcode Console, `idevicesyslog`, and Crash Reports) can be read by:
- Developers with physical device access
- Malicious apps with log access entitlements
- Anyone who can read the device logs via iTunes backup

Logging passwords, tokens, PII, or financial data is a compliance violation.

## Remediation
- Remove all NSLog/print calls in production builds using macros
- Use `os_log` with `%{private}` format specifier for sensitive data
- Build with `DEBUG` flag guard: `#if DEBUG ... NSLog(...) ... #endif`
