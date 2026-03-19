---
confidence: high
cwe:
    - CWE-502
description: Detects JSON deserialization configurations that allow polymorphic type instantiation, enabling object injection attacks.
languages:
    - java
    - javascript
    - typescript
    - python
    - kotlin
    - dart
    - zig
    - elixir
name: Insecure JSON Deserialization with Type Polymorphism
owasp:
    - A08:2025
severity: high
tags:
    - deserialization
    - json
    - owasp-a08
version: 1.0.0
---

# Insecure JSON Deserialization

## Overview
While JSON itself does not execute code, certain JSON library configurations allow attackers to instantiate arbitrary classes through type metadata embedded in JSON. The most notable example is Jackson's `enableDefaultTyping()` which was responsible for multiple critical CVEs.

## Remediation
- Jackson: Never use `enableDefaultTyping()`, use `@JsonTypeInfo` with explicit subtypes
- Newtonsoft.Json: Avoid `TypeNameHandling.All`
- Validate and whitelist expected types before deserialization
