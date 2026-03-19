---
name: Secrets in Dockerfile
version: 1.0.0
description: Detects credentials and secrets hardcoded in Dockerfile ENV instructions or ARG values that end up in image layers.
tags: [devops, docker, secrets, owasp-a07]
languages: [generic]
severity: critical
confidence: high
cwe: [CWE-312, CWE-798]
owasp: [A07:2025]
---

# Secrets in Dockerfile

## Overview
Secrets placed in Dockerfile `ENV` or `ARG` instructions are baked into image layers and visible via `docker inspect` or `docker history`, even if removed in a later layer. Anyone with pull access to the image can read these secrets.

## Remediation
- Use Docker secrets (`--secret`) for sensitive values during build
- Use multi-stage builds to exclude build-time secrets from final image
- Use `.env` files with `--env-file` at runtime instead of baking into image
- Use HashiCorp Vault, AWS Secrets Manager, or Kubernetes Secrets for runtime injection
