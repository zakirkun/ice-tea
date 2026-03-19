---
confidence: high
cwe:
    - CWE-312
    - CWE-922
description: Detects sensitive data stored in plaintext in SharedPreferences, SQLite databases, or external storage in Android applications.
languages:
    - java
    - kotlin
    - dart
    - zig
    - elixir
name: Android Insecure Data Storage
owasp:
    - A02:2025
severity: high
tags:
    - android
    - mobile
    - data-storage
    - owasp-m2
version: 1.0.0
---

# Android Insecure Data Storage

## Overview
Android applications commonly store sensitive data insecurely:
1. **SharedPreferences plaintext**: Tokens, passwords, PII stored in XML files readable on rooted devices
2. **External storage**: Files on SD card readable by all apps with READ_EXTERNAL_STORAGE permission
3. **SQLite without encryption**: Sensitive data in databases accessible after physical extraction
4. **Logcat logs**: Sensitive data logged and accessible via `adb logcat`
5. **Clipboard**: Sensitive data copied to clipboard accessible by other apps

## Detection Strategy
- `SharedPreferences.edit().putString("password", ...)` — storing sensitive values in plain SharedPreferences
- `Environment.getExternalStorageDirectory()` — writing to external storage
- `Log.d/i/v/w/e` with sensitive parameter values

## Remediation
- Use Android Keystore System for cryptographic keys
- Use EncryptedSharedPreferences (Jetpack Security library)
- Store sensitive files in internal storage with MODE_PRIVATE
- Use SQLCipher for encrypted database storage
- Never log sensitive data
