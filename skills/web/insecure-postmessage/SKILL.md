---
name: Insecure PostMessage Configuration
version: 1.0.0
description: Detects careless use of the HTML5 Web Messaging API (postMessage), such as sending messages to the generic '*' origin.
tags: [web, frontend, postmessage, cors]
languages: [generic]
severity: medium
confidence: medium
cwe: [CWE-346, CWE-942]
owasp: [A05:2025]
---

# Insecure PostMessage Configuration

## Overview
Detects careless use of the HTML5 Web Messaging API (postMessage), such as sending messages to the generic '*' origin.

## Remediation
Always specify the exact target origin in window.postMessage(). When receiving messages, rigorously verify event.origin.
