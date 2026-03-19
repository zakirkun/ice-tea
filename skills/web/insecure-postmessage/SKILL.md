---
confidence: medium
cwe:
    - CWE-346
    - CWE-942
description: Detects careless use of the HTML5 Web Messaging API (postMessage), such as sending messages to the generic '*' origin.
languages:
    - generic
    - kotlin
    - dart
    - zig
    - elixir
name: Insecure PostMessage Configuration
owasp:
    - A05:2025
severity: medium
tags:
    - web
    - frontend
    - postmessage
    - cors
version: 1.0.0
---

# Insecure PostMessage Configuration

## Overview
Detects careless use of the HTML5 Web Messaging API (postMessage), such as sending messages to the generic '*' origin.

## Remediation
Always specify the exact target origin in window.postMessage(). When receiving messages, rigorously verify event.origin.
