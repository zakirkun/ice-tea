---
name: Server Timing Information Leak
version: 1.0.0
description: Detects Server-Timing headers that expose internal component names and timing data to clients.
tags: [web, information-disclosure, timing, owasp-a05]
languages: [javascript, typescript, python, go, java, php]
severity: low
confidence: high
cwe: [CWE-208]
owasp: [A05:2025]
---

# Server Timing Information Leak

## Overview
The `Server-Timing` HTTP header exposes timing information about server-side operations to browsers. While useful for performance profiling, it can reveal:
- Internal service names (database, cache, microservices)
- Query execution times (aids SQL injection timing attacks)
- Internal architecture details

## Remediation
- Remove Server-Timing headers in production
- If needed for monitoring, limit to non-sensitive metric names
- Restrict Server-Timing to same-origin via `Timing-Allow-Origin` header
