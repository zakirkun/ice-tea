---
confidence: medium
cwe:
    - CWE-502
description: Detects MessagePack deserialization configured to allow arbitrary object construction.
languages:
    - python
    - javascript
    - typescript
    - java
    - go
    - kotlin
    - dart
    - zig
    - elixir
name: Unsafe MessagePack Deserialization
owasp:
    - A08:2025
severity: high
tags:
    - deserialization
    - msgpack
    - rce
    - owasp-a08
version: 1.0.0
---

# Unsafe MessagePack Deserialization

## Overview
MessagePack is generally safer than pickle/Java serialization, but certain configurations and extensions can allow type coercion or code execution when combined with object extensibility features.

## Remediation
- Use `raw=True` in Python msgpack to avoid string coercion
- Validate deserialized data against a schema before use
- Do not pass deserialized objects directly to code execution paths
