---
confidence: medium
cwe:
    - CWE-1321
description: Detects patterns that allow direct manipulation of the JavaScript prototype chain, enabling prototype pollution and security bypass.
languages:
    - javascript
    - typescript
    - kotlin
    - dart
    - zig
    - elixir
name: JavaScript Prototype Chain Manipulation
owasp:
    - A08:2025
severity: high
tags:
    - injection
    - prototype-chain
    - javascript
    - owasp-a08
version: 1.0.0
---

# JavaScript Prototype Chain Manipulation

## Overview
Beyond basic `__proto__` assignment, prototype chain attacks include:
- Setting properties on `Object.prototype` via indirect paths
- Using `constructor.prototype` to poison objects
- `Object.setPrototypeOf()` with attacker-controlled objects
- Vulnerable deserialization that constructs objects on attacker-defined chains

## Remediation
- Freeze base object prototypes in security-sensitive code: `Object.freeze(Object.prototype)`
- Use `Object.create(null)` for pure hash maps
- Validate all keys before property assignment
