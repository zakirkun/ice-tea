---
name: Missing Input Size Limits
version: 1.0.0
description: Detects API endpoints without request body size limits, enabling denial of service via oversized payloads.
tags: [api, dos, resource-exhaustion, owasp-a04]
languages: [javascript, typescript, python, go, java, php]
severity: high
confidence: low
cwe: [CWE-770]
owasp: [A04:2025]
---

# Missing Input Size Limits

## Overview
APIs without request body size limits are vulnerable to:
- Memory exhaustion by sending multi-gigabyte request bodies
- CPU exhaustion via large JSON/XML documents requiring expensive parsing
- Storage DoS if large uploads are buffered to disk

## Remediation
- Set explicit body size limits: Express default is 100kb, consider 1MB max for most APIs
- Use streaming parsers for large uploads instead of buffering in memory
- Implement per-user upload quota
