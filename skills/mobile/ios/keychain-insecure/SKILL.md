---
name: iOS Insecure Keychain Usage
version: 1.0.0
description: Detects iOS Keychain items stored with insecure accessibility attributes that allow access when device is locked or without user authentication.
tags: [ios, mobile, keychain, owasp-m2]
languages: [java, generic]
severity: high
confidence: high
cwe: [CWE-312]
owasp: [A02:2025]
---

# iOS Insecure Keychain Usage

## Overview
iOS Keychain stores sensitive data with different accessibility levels:
- `kSecAttrAccessibleAlways` — accessible even when device is locked (insecure)
- `kSecAttrAccessibleWhenUnlocked` — accessible only when device is unlocked (preferred)
- `kSecAttrAccessibleAfterFirstUnlock` — accessible after first unlock (less secure)

Using `kSecAttrAccessibleAlways` means data can be read by malware or when device is physically stolen while locked.

## Remediation
- Use `kSecAttrAccessibleWhenUnlockedThisDeviceOnly` for most sensitive items
- Use `kSecAttrAccessibleWhenPasscodeSetThisDeviceOnly` for highest security
- Never use `kSecAttrAccessibleAlways`
