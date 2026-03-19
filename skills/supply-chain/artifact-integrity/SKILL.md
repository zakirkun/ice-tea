---
name: Artifact Integrity Verification Missing
version: 1.0.0
description: Detects download and execution of artifacts without cryptographic hash or signature verification.
tags: [supply-chain, integrity, owasp-a06]
languages: [generic, yaml, javascript, python]
severity: high
confidence: medium
cwe: [CWE-494]
owasp: [A06:2025]
---

# Artifact Integrity Verification Missing

## Overview
Downloading binary artifacts, scripts, or packages from URLs without verifying their cryptographic signatures or hashes allows man-in-the-middle attacks or compromised CDN attacks to deliver malicious code.

## Remediation
- Always verify SHA256 hash of downloaded artifacts
- Use `--checksum` flags where available
- Use Subresource Integrity (SRI) hashes for CDN resources in HTML
- Verify GPG signatures for critical software
