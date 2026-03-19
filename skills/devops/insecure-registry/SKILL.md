---
name: Insecure Container Registry Configuration
version: 1.0.0
description: Detects Docker/container registry configurations using unauthenticated or HTTP (non-HTTPS) registries.
tags: [devops, docker, registry, owasp-a05]
languages: [generic, yaml]
severity: high
confidence: high
cwe: [CWE-327]
owasp: [A05:2025]
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
