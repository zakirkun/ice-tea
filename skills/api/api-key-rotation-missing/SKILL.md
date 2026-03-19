---
name: API Key Without Rotation or Expiration
version: 1.0.0
description: Detects API key implementations without expiration dates or rotation mechanisms, creating long-lived credentials.
tags: [api, api-key, rotation, owasp-a07]
languages: [javascript, typescript, python, go, java, php]
severity: medium
confidence: low
cwe: [CWE-522]
owasp: [A07:2025]
---

# API Key Without Rotation

## Overview
Static API keys without expiration or rotation remain valid indefinitely after compromise. Best practices require:
- Short-lived keys with automatic expiration
- Easy rotation mechanism
- Key usage monitoring and anomaly detection
- Immediate revocation capability

## Remediation
- Add expiration dates to all API keys
- Implement key rotation workflows
- Notify users before expiration
- Monitor for unusual usage patterns
