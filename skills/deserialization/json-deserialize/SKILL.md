---
name: Insecure JSON Deserialization with Type Polymorphism
version: 1.0.0
description: Detects JSON deserialization configurations that allow polymorphic type instantiation, enabling object injection attacks.
tags: [deserialization, json, owasp-a08]
languages: [java, javascript, typescript, python]
severity: high
confidence: high
cwe: [CWE-502]
owasp: [A08:2025]
---

# Insecure JSON Deserialization

## Overview
While JSON itself does not execute code, certain JSON library configurations allow attackers to instantiate arbitrary classes through type metadata embedded in JSON. The most notable example is Jackson's `enableDefaultTyping()` which was responsible for multiple critical CVEs.

## Remediation
- Jackson: Never use `enableDefaultTyping()`, use `@JsonTypeInfo` with explicit subtypes
- Newtonsoft.Json: Avoid `TypeNameHandling.All`
- Validate and whitelist expected types before deserialization
