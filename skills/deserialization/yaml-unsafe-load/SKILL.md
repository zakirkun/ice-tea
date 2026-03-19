---
name: Unsafe YAML Deserialization
version: 1.0.0
description: Detects YAML parsing using unsafe loaders that execute arbitrary Python or Ruby code embedded in YAML.
tags: [deserialization, yaml, rce, owasp-a08]
languages: [python, ruby, javascript, typescript]
severity: critical
confidence: high
cwe: [CWE-502]
owasp: [A08:2025]
---

# Unsafe YAML Deserialization

## Overview
YAML loaders that support the full YAML specification (including `!!python/object/apply:`) can execute arbitrary code when parsing malicious YAML documents.

Python's `yaml.load()` without `Loader=yaml.SafeLoader` is the most common instance.

Malicious payload:
```yaml
!!python/object/apply:os.system ["id"]
```

## Remediation
- Python: Use `yaml.safe_load()` or `yaml.load(data, Loader=yaml.SafeLoader)`
- Ruby: Use `YAML.safe_load()` instead of `YAML.load()`
- Node.js: `js-yaml` uses `safeLoad` (default safe, but `load()` with `unsafe=true` is dangerous)
