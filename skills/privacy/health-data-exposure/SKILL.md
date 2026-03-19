---
name: Health / Medical Data Exposure
version: 1.0.0
description: Detects health and medical information logged, transmitted without encryption, or insufficiently protected, violating HIPAA and similar regulations.
tags: [privacy, hipaa, health-data, owasp-a02]
languages: [javascript, typescript, python, go, java, php]
severity: critical
confidence: medium
cwe: [CWE-359]
owasp: [A02:2025]
---

# Health / Medical Data Exposure

## Overview
Protected Health Information (PHI) under HIPAA includes diagnoses, medications, treatment history, and health identifiers. Improper handling creates legal liability and severe breach consequences.

## Remediation
- Encrypt PHI at rest and in transit
- Implement audit logging for all PHI access
- Apply minimum necessary access principles
- Never log PHI without encryption and access controls
