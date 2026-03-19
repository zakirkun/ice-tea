---
confidence: high
cwe:
    - CWE-732
    - CWE-798
description: Detects common Azure security misconfigurations including public blob storage, overprivileged managed identities, and insecure ARM templates.
languages:
    - yaml
    - generic
    - python
    - javascript
    - kotlin
    - dart
    - zig
    - elixir
name: Azure Misconfiguration
owasp:
    - A05:2025
severity: critical
tags:
    - azure
    - cloud
    - arm-template
    - owasp-a05
version: 1.0.0
---

# Azure Misconfiguration

## Overview
Common Azure security misconfigurations:
1. **Public blob containers**: Storage containers with anonymous read access
2. **Hardcoded connection strings**: Storage/database connection strings with keys in code
3. **Overprivileged managed identity**: Subscription-level contributor role
4. **No MFA for admin accounts**: Admin accounts without MFA requirement
5. **NSG allowing all inbound**: Network Security Groups open to internet
6. **Key Vault access policies too broad**: Full access granted to many principals

## Remediation
- Set Storage Account containers to private
- Use Managed Identity instead of hardcoded connection strings
- Apply Azure Policy to enforce security baselines
- Enable Microsoft Defender for Cloud
