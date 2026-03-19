---
name: Kubernetes Misconfiguration
version: 1.0.0
description: Detects insecure Kubernetes pod and container configurations including privileged containers, hostPID, hostNetwork, and missing resource limits.
tags: [kubernetes, k8s, container, infra, owasp-a05]
languages: [yaml]
severity: high
confidence: high
cwe: [CWE-732, CWE-250]
owasp: [A05:2025]
---

# Kubernetes Misconfiguration

## Overview
Kubernetes misconfigurations can lead to cluster takeover and container escapes:
1. **`privileged: true`**: Pod container with root-level host access
2. **`hostPID: true`**: Pod shares host PID namespace — can see/kill host processes
3. **`hostNetwork: true`**: Pod uses host network, bypassing network policies
4. **`allowPrivilegeEscalation: true`**: Allows container processes to gain more privileges
5. **`runAsRoot: true` or no `runAsNonRoot`**: Container runs as root
6. **No resource limits**: Resource exhaustion DoS possible
7. **Wildcard RBAC permissions**: `resources: ["*"]` with `verbs: ["*"]`

## Remediation
- Use `securityContext` with `runAsNonRoot: true`, `allowPrivilegeEscalation: false`
- Apply `readOnlyRootFilesystem: true`
- Set resource requests and limits
- Follow principle of least privilege for RBAC
