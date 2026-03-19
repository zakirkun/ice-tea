---
name: AWS Misconfiguration
version: 1.0.0
description: Detects common AWS security misconfigurations including public S3 buckets, IAM wildcard policies, hardcoded AWS credentials, and unencrypted resources.
tags: [aws, cloud, s3, iam, owasp-a05]
languages: [python, javascript, typescript, go, java, yaml, generic]
severity: critical
confidence: high
cwe: [CWE-732, CWE-798, CWE-312]
owasp: [A05:2025]
---

# AWS Misconfiguration

## Overview
AWS misconfigurations are responsible for major data breaches. Common issues:
1. **Public S3 buckets**: Accidental public-read or public-write ACL
2. **IAM wildcard policies**: `Action: "*"` grants all permissions
3. **Hardcoded AWS credentials**: Access keys in source code
4. **Unencrypted RDS/S3**: Data at rest not encrypted
5. **Overly permissive security groups**: Port 22 or 3306 open to `0.0.0.0/0`
6. **CloudTrail disabled**: No audit trail

## Remediation
- Enable S3 Block Public Access at the account level
- Apply principle of least privilege to all IAM policies
- Use IAM roles instead of hardcoded credentials
- Enable encryption at rest for all services
- Use VPC security groups with minimal required access
