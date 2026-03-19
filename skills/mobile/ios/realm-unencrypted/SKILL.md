---
name: iOS Realm Database Without Encryption
version: 1.0.0
description: Detects iOS Realm database instances configured without encryption key.
tags: [ios, mobile, database, encryption, owasp-m2]
languages: [generic]
severity: high
confidence: high
cwe: [CWE-311]
owasp: [A02:2025]
---

# iOS Realm Database Without Encryption

## Overview
Realm database files are stored unencrypted by default. On jailbroken devices or after physical extraction, these files can be opened directly. Sensitive user data (including PII, financial data, health records) stored in unencrypted Realm databases is exposed.

## Remediation
- Configure Realm with an encryption key stored in the iOS Keychain
- Generate a 64-byte key using SecRandomCopyBytes
- Store the key with `kSecAttrAccessibleWhenUnlockedThisDeviceOnly`
