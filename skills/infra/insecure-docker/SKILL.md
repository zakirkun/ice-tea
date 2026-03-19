---
name: Insecure Docker Configuration
version: 1.0.0
description: Detects dangerous Docker and Docker Compose configurations including privileged containers, secrets in environment variables, and running as root.
tags: [docker, container, infra, owasp-a05]
languages: [yaml, generic]
severity: high
confidence: high
cwe: [CWE-250, CWE-732]
owasp: [A05:2025]
---

# Insecure Docker Configuration

## Overview
Docker misconfiguration can lead to container escapes and privilege escalation:
1. **`--privileged` flag**: Gives container full host capabilities, enabling container escape
2. **Running as root**: Default behavior — prefer non-root user
3. **Secrets in ENV variables**: Visible via `docker inspect` and process listings
4. **`--cap-add=SYS_ADMIN`**: Grants dangerous capabilities
5. **Mounting Docker socket**: Allows container to control the host Docker daemon
6. **`--pid=host`**: Shares host PID namespace

## Remediation
- Add `USER nonroot` in Dockerfile
- Use Docker secrets or env files for sensitive values
- Remove unnecessary capabilities with `--cap-drop ALL --cap-add` only what's needed
- Never mount `/var/run/docker.sock` unless strictly necessary
- Enable Docker Content Trust
