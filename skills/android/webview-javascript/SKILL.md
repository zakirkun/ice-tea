---
confidence: high
cwe:
    - CWE-749
    - CWE-79
description: Detects insecure Android WebView configurations enabling XSS, JavaScript bridge abuse, and remote code execution.
languages:
    - java
    - kotlin
    - dart
    - zig
    - elixir
name: Android WebView JavaScript Risks
owasp:
    - A03:2025
severity: critical
tags:
    - android
    - webview
    - javascript
    - owasp-m1
version: 1.0.0
---

# Android WebView JavaScript Risks

## Overview
Android WebView is a powerful component that can be misconfigured to allow serious attacks:
1. **`setJavaScriptEnabled(true)` + `addJavascriptInterface()`**: Creates a JavaScript bridge allowing web pages to call Java methods — RCE risk
2. **`setAllowFileAccessFromFileURLs(true)`**: JavaScript in file:// URLs can read other files
3. **Loading untrusted URLs**: Loading attacker-controlled URLs with full bridge access
4. **`setAllowUniversalAccessFromFileURLs(true)`**: JavaScript in file:// can make cross-origin requests

## Remediation
- Only enable JavaScript for trusted content
- Remove `addJavascriptInterface()` or add `@JavascriptInterface` annotation carefully
- Validate URLs before loading them in WebView
- Use `setWebContentsDebuggingEnabled(false)` in production
