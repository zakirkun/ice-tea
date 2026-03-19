---
confidence: medium
cwe:
    - CWE-494
description: Detects CI/CD pipelines that download, use, or publish build artifacts without cryptographic hash verification.
languages:
    - yaml
    - generic
    - kotlin
    - dart
    - zig
    - elixir
name: Build Artifact Without Integrity Verification
owasp:
    - A06:2025
severity: high
tags:
    - devops
    - supply-chain
    - integrity
    - owasp-a06
version: 1.0.0
---

# Build Artifact Without Integrity Verification

## Overview
CI/CD pipelines that download tools, binaries, or artifacts without verifying their SHA256 checksums are vulnerable to supply chain attacks where a compromised CDN or package server serves malicious payloads.

## Remediation
- Always verify SHA256 checksums after downloading artifacts
- Use package managers with lock file integrity checks
- Pin artifact versions to immutable digests where possible
