---
name: GCP Misconfiguration
version: 1.0.0
description: Detects common Google Cloud Platform security misconfigurations including public storage buckets, overprivileged service accounts, and disabled audit logging.
tags: [gcp, cloud, google-cloud, owasp-a05]
languages: [yaml, generic, python, javascript, go]
severity: critical
confidence: high
cwe: [CWE-732, CWE-798]
owasp: [A05:2025]
---

# GCP Misconfiguration

## Overview
Common GCP security misconfigurations:
1. **Public Cloud Storage buckets**: `allUsers` or `allAuthenticatedUsers` access
2. **Service account key file exposure**: SA keys committed to version control
3. **Overprivileged service accounts**: `roles/editor` or `roles/owner` granted
4. **Default service account auto-mount**: Pods automatically get default SA token
5. **Disabled audit logging**: Data access logs disabled
6. **GKE legacy ABAC**: Insecure legacy authorization model

## Remediation
- Enable Uniform Bucket-Level Access to prevent per-object ACLs
- Use Workload Identity instead of service account key files
- Follow principle of least privilege for IAM bindings
- Enable VPC Service Controls
