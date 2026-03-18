---
name: Insecure Deserialization
version: 1.0.0
description: Detects deserialization of untrusted data which can lead to Remote Code Execution.
tags: [deserialization, rce, owasp-a08]
languages: [generic]
severity: critical
confidence: medium
cwe: [CWE-502]
owasp: [A08:2025]
---

# Insecure Deserialization

## Overview
Detects deserialization of untrusted data which can lead to Remote Code Execution.

## Remediation
Use safe data formats like JSON instead of native serialization. If native serialization is required, sign the objects.
