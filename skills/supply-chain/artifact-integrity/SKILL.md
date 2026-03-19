---
confidence: medium
cwe:
    - CWE-494
description: Detects download and execution of artifacts without cryptographic hash or signature verification.
languages:
    - generic
    - yaml
    - javascript
    - python
    - kotlin
    - dart
    - zig
    - elixir
name: Artifact Integrity Verification Missing
owasp:
    - A06:2025
severity: high
tags:
    - supply-chain
    - integrity
    - owasp-a06
version: 1.0.0
---

# Artifact Integrity Verification Missing

## Overview
Downloading binary artifacts, scripts, or packages from URLs without verifying their cryptographic signatures or hashes allows man-in-the-middle attacks or compromised CDN attacks to deliver malicious code.

## Remediation
- Always verify SHA256 hash of downloaded artifacts
- Use `--checksum` flags where available
- Use Subresource Integrity (SRI) hashes for CDN resources in HTML
- Verify GPG signatures for critical software
