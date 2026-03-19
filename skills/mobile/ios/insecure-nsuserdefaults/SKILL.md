---
confidence: high
cwe:
    - CWE-312
description: Detects storage of sensitive information in NSUserDefaults which is unencrypted and accessible in device backups.
languages:
    - generic
    - kotlin
    - dart
    - zig
    - elixir
name: Sensitive Data in NSUserDefaults
owasp:
    - A02:2025
severity: high
tags:
    - ios
    - mobile
    - data-storage
    - owasp-m2
version: 1.0.0
---

# Sensitive Data in NSUserDefaults

## Overview
`NSUserDefaults` stores data in a plaintext plist file in the app's Library directory. This data:
- Is included in unencrypted iTunes/iCloud backups
- Is accessible on jailbroken devices
- May be logged by system diagnostics

Sensitive data such as auth tokens, passwords, or PII must never be stored in NSUserDefaults.

## Remediation
- Use iOS Keychain for credentials and tokens
- Use encrypted Core Data or Realm for sensitive structured data
- Mark sensitive files with `NSURLIsExcludedFromBackupKey` if they must be in the app sandbox
