---
confidence: high
cwe:
    - CWE-732
    - CWE-311
description: Detects security misconfigurations in Terraform, Pulumi, and CloudFormation templates including overly permissive resources and disabled security features.
languages:
    - generic
    - yaml
    - kotlin
    - dart
    - zig
    - elixir
name: Infrastructure as Code Security Issues
owasp:
    - A05:2025
severity: high
tags:
    - iac
    - terraform
    - cloudformation
    - owasp-a05
version: 1.0.0
---

# Infrastructure as Code Security Issues

## Overview
Infrastructure as Code (IaC) templates define cloud resources and their security posture. Misconfigurations in Terraform/CloudFormation/Pulumi templates propagate to all deployed environments.

Common IaC security issues:
- Encryption disabled on storage/databases
- Public-facing databases without VPC restriction
- Logging and monitoring disabled
- Overly permissive IAM and security groups
- SSH access open to all IPs

## Remediation
Use security linting tools (tfsec, checkov, cfn-nag, KICS) in CI/CD pipelines to catch IaC misconfigurations before deployment.
