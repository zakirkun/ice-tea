---
name: Prototype Pollution
version: 1.0.0
description: Detects JavaScript prototype pollution vulnerabilities where attacker-controlled keys can modify Object.prototype.
tags: [prototype-pollution, javascript, web, owasp-a08]
languages: [javascript, typescript]
severity: high
confidence: medium
cwe: [CWE-1321]
owasp: [A08:2025]
---

# Prototype Pollution

## Overview
Prototype pollution occurs when JavaScript code recursively merges or assigns properties from user-controlled objects without filtering dangerous keys like `__proto__`, `constructor`, or `prototype`. Modifying `Object.prototype` affects all JavaScript objects in the application, enabling:
- Authentication bypass (`isAdmin` injected into base object)
- Denial of Service
- Remote Code Execution (in some Node.js contexts)

## Detection Strategy
- Recursive merge/assign functions without `__proto__` key filtering
- Direct `obj[key] = value` where `key` comes from user input
- Libraries known to be vulnerable (lodash < 4.17.13, merge < 2.1.1)

## Remediation
- Validate and reject `__proto__`, `constructor`, `prototype` as keys
- Use `Object.create(null)` for pure hash maps
- Use `JSON.parse()` with schema validation before merging
- Update vulnerable dependencies

**Vulnerable:**
```js
function merge(target, src) {
    for (const key of Object.keys(src)) {
        target[key] = src[key]; // key could be __proto__
    }
}
```

**Safe:**
```js
const FORBIDDEN_KEYS = new Set(['__proto__', 'constructor', 'prototype']);
function merge(target, src) {
    for (const key of Object.keys(src)) {
        if (!FORBIDDEN_KEYS.has(key)) target[key] = src[key];
    }
}
```
