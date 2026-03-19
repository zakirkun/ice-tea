---
confidence: medium
cwe:
    - CWE-656
description: Detects weak jailbreak detection implementations that can be easily bypassed.
languages:
    - generic
    - kotlin
    - dart
    - zig
    - elixir
name: iOS Jailbreak Detection Bypass Vulnerability
owasp:
    - A08:2025
severity: medium
tags:
    - ios
    - mobile
    - jailbreak
    - owasp-m8
version: 1.0.0
---

# iOS Jailbreak Detection Bypass

## Overview
Weak jailbreak detection that only checks for the existence of files like `/Applications/Cydia.app` can be bypassed by jailbreak tools that hide these files. Applications relying on security features that assume a trusted device environment need robust jailbreak detection.

## Remediation
- Use multiple detection methods (file system checks, process checks, environment checks)
- Consider using an approved mobile security library (IOSSecuritySuite)
- Defense-in-depth: assume some users will bypass detection and protect sensitive operations at the server
