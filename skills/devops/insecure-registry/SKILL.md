---
confidence: high
cwe:
    - CWE-327
description: Detects Docker/container registry configurations using unauthenticated or HTTP (non-HTTPS) registries.
languages:
    - generic
    - yaml
    - kotlin
    - dart
    - zig
    - elixir
name: Insecure Container Registry Configuration
owasp:
    - A05:2025
severity: high
tags:
    - devops
    - docker
    - registry
    - owasp-a05
version: 1.0.0
---

# Insecure Container Registry Configuration

## Overview
Using insecure Docker registries (HTTP rather than HTTPS, or self-signed certificates without verification) exposes:
- Image pull/push credentials to network interception
- Registry poisoning through MITM — serving malicious images

## Remediation
- Use HTTPS for all container registries
- Configure proper TLS certificates
- Use `--insecure-registry` only in development, never in production
