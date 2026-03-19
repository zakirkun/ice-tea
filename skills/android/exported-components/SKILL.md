---
name: Android Exported Components Without Permission
version: 1.0.0
description: Detects Android Activities, Services, BroadcastReceivers, and ContentProviders exported without proper permission checks.
tags: [android, mobile, exported-components, owasp-m1]
languages: [java, generic]
severity: high
confidence: high
cwe: [CWE-926, CWE-284]
owasp: [A01:2025]
---

# Android Exported Components Without Permission

## Overview
Android components (Activity, Service, BroadcastReceiver, ContentProvider) can be declared `exported=true` in AndroidManifest.xml, making them accessible to other applications. Without permission requirements, any installed app can:
- Start sensitive Activities (bypassing login screens)
- Bind to sensitive Services (access internal functionality)
- Trigger BroadcastReceivers (inject events)
- Query ContentProviders (access application data)

## Detection Strategy
- `android:exported="true"` without `android:permission` attribute
- Intent filters implicitly make components exported on older API levels
- ContentProvider with `android:readPermission` or `android:writePermission` missing

## Remediation
- Set `android:exported="false"` for components not intended for external use
- Add `android:permission` with a custom signature-level permission
- Validate caller identity with `checkCallingPermission()` or `checkCallingOrSelfPermission()`
