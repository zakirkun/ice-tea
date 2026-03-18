# CI/CD Integration & SARIF Output

## Overview

Ice Tea is designed as a CI/CD-native security scanner. It produces standardized output formats that integrate directly with GitHub Code Scanning and GitLab Security Dashboards.

## Output Formats

### 1. SARIF (Static Analysis Results Interchange Format)

SARIF v2.1.0 is the primary output format. It is a JSON-based standard adopted by GitHub, Microsoft, and many security tools.

#### SARIF Structure

```json
{
  "$schema": "https://raw.githubusercontent.com/oasis-tcs/sarif-spec/main/sarif-2.1/schema/sarif-schema-2.1.0.json",
  "version": "2.1.0",
  "runs": [
    {
      "tool": {
        "driver": {
          "name": "ice-tea",
          "version": "1.0.0",
          "informationUri": "https://github.com/user/ice-tea",
          "rules": [
            {
              "id": "G201",
              "name": "sql-injection-concat",
              "shortDescription": {
                "text": "SQL query built with string concatenation"
              },
              "fullDescription": {
                "text": "Detected SQL query constructed using string concatenation with potentially untrusted input..."
              },
              "help": {
                "text": "Use parameterized queries instead of string concatenation."
              },
              "properties": {
                "tags": ["security", "injection", "sql"],
                "precision": "high",
                "security-severity": "9.0"
              }
            }
          ]
        }
      },
      "results": [
        {
          "ruleId": "G201",
          "level": "error",
          "message": {
            "text": "SQL injection: user input from r.URL.Query() flows into db.Query() without parameterization"
          },
          "locations": [
            {
              "physicalLocation": {
                "artifactLocation": {
                  "uri": "handlers/user.go"
                },
                "region": {
                  "startLine": 42,
                  "endLine": 42,
                  "startColumn": 5,
                  "endColumn": 55
                }
              }
            }
          ],
          "partialFingerprints": {
            "primaryLocationLineHash": "abc123..."
          },
          "codeFlows": [
            {
              "threadFlows": [
                {
                  "locations": [
                    {
                      "location": {
                        "physicalLocation": {
                          "artifactLocation": { "uri": "handlers/user.go" },
                          "region": { "startLine": 38 }
                        },
                        "message": { "text": "Source: user input from URL query" }
                      }
                    },
                    {
                      "location": {
                        "physicalLocation": {
                          "artifactLocation": { "uri": "handlers/user.go" },
                          "region": { "startLine": 42 }
                        },
                        "message": { "text": "Sink: unsanitized input in SQL query" }
                      }
                    }
                  ]
                }
              ]
            }
          ]
        }
      ]
    }
  ]
}
```

#### Key SARIF Features

| Feature | Purpose |
|---------|---------|
| `partialFingerprints` | Prevents duplicate alerts across runs |
| `codeFlows` | Shows data flow path (source → sink) |
| `security-severity` | CVSS-like score for prioritization |
| `help` | Remediation guidance per rule |

### 2. GitLab SAST JSON Report

GitLab requires a specific JSON format (`gl-sast-report.json`):

```json
{
  "version": "15.0.0",
  "vulnerabilities": [
    {
      "id": "unique-finding-id",
      "category": "sast",
      "name": "SQL Injection via string concatenation",
      "message": "User input flows into SQL query...",
      "description": "Detailed description...",
      "severity": "Critical",
      "confidence": "High",
      "scanner": {
        "id": "ice-tea",
        "name": "Ice Tea Security Scanner"
      },
      "location": {
        "file": "handlers/user.go",
        "start_line": 42,
        "end_line": 42
      },
      "identifiers": [
        {
          "type": "cwe",
          "name": "CWE-89",
          "value": "89",
          "url": "https://cwe.mitre.org/data/definitions/89.html"
        }
      ]
    }
  ]
}
```

### 3. Console Output

Human-readable terminal output with severity coloring:

```
[CRITICAL] handlers/user.go:42 - SQL Injection (G201)
  SQL query built with string concatenation using untrusted input
  CWE-89 | OWASP A05:2025
  Fix: Use parameterized queries instead of string concatenation

[HIGH] auth/login.go:15 - Hardcoded Credential (G101)
  Password found hardcoded in source code
  CWE-798 | OWASP A07:2025
  Fix: Use environment variables or a secrets manager

Summary: 2 critical, 1 high, 3 medium, 0 low (6 total findings)
```

### 4. JSON Output

Raw JSON array of findings for programmatic consumption.

## GitHub Actions Integration

### GitHub Action Definition

```yaml
# .github/actions/ice-tea/action.yml
name: 'Ice Tea Security Scanner'
description: 'AI-powered security scanner for code'
inputs:
  target:
    description: 'Directory or file to scan'
    required: false
    default: '.'
  config:
    description: 'Path to configuration file'
    required: false
  severity-threshold:
    description: 'Minimum severity to report (critical, high, medium, low)'
    required: false
    default: 'medium'
  enable-llm:
    description: 'Enable LLM deep reasoning engine'
    required: false
    default: 'false'
  llm-api-key:
    description: 'API key for LLM provider'
    required: false
runs:
  using: 'composite'
  steps:
    - name: Download Ice Tea
      shell: bash
      run: |
        curl -sSfL https://github.com/user/ice-tea/releases/latest/download/ice-tea-linux-amd64 -o /usr/local/bin/ice-tea
        chmod +x /usr/local/bin/ice-tea
    - name: Run Scan
      shell: bash
      run: |
        ice-tea scan ${{ inputs.target }} \
          --format sarif \
          --output results.sarif \
          --severity ${{ inputs.severity-threshold }} \
          ${{ inputs.enable-llm == 'true' && '--enable-llm' || '' }}
    - name: Upload SARIF
      uses: github/codeql-action/upload-sarif@v3
      with:
        sarif_file: results.sarif
```

### Usage in Workflow

```yaml
# .github/workflows/security.yml
name: Security Scan
on: [push, pull_request]

jobs:
  ice-tea-scan:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: ./.github/actions/ice-tea
        with:
          severity-threshold: medium
          enable-llm: true
          llm-api-key: ${{ secrets.LLM_API_KEY }}
```

## GitLab CI Integration

```yaml
# .gitlab-ci.yml
include:
  - template: Security/SAST.gitlab-ci.yml

ice-tea-sast:
  stage: test
  image: golang:1.22
  before_script:
    - curl -sSfL https://github.com/user/ice-tea/releases/latest/download/ice-tea-linux-amd64 -o /usr/local/bin/ice-tea
    - chmod +x /usr/local/bin/ice-tea
  script:
    - ice-tea scan . --format gitlab --output gl-sast-report.json
  artifacts:
    reports:
      sast: gl-sast-report.json
```

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | No findings above threshold |
| 1 | Findings found above severity threshold |
| 2 | Configuration or runtime error |
| 3 | Invalid input (bad target path, etc.) |
