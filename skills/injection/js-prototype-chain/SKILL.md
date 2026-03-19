---
name: JavaScript Prototype Chain Manipulation
version: 1.0.0
description: Detects patterns that allow direct manipulation of the JavaScript prototype chain, enabling prototype pollution and security bypass.
tags: [injection, prototype-chain, javascript, owasp-a08]
languages: [javascript, typescript]
severity: high
confidence: medium
cwe: [CWE-1321]
owasp: [A08:2025]
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
