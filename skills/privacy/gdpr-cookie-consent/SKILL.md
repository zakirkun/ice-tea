---
name: Missing GDPR Cookie Consent
version: 1.0.0
description: Detects analytics and tracking scripts loaded without cookie consent mechanisms.
tags: [privacy, gdpr, cookie-consent, owasp-a05]
languages: [javascript, typescript, generic]
severity: medium
confidence: medium
cwe: [CWE-311]
owasp: [A05:2025]
---

# Missing GDPR Cookie Consent

## Overview
Under GDPR and ePrivacy Directive, non-essential cookies (analytics, marketing, tracking) require explicit user consent before being set. Loading analytics scripts unconditionally violates privacy law.

## Remediation
- Implement a Consent Management Platform (CMP)
- Only load analytics scripts after explicit consent
- Provide granular consent options (necessary, analytics, marketing)
- Store consent records with timestamps
