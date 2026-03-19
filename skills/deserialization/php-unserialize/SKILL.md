---
confidence: high
cwe:
    - CWE-502
description: Detects PHP unserialize() called on user-controlled input, enabling object injection and remote code execution.
languages:
    - php
    - kotlin
    - dart
    - zig
    - elixir
name: PHP Insecure Deserialization (unserialize)
owasp:
    - A08:2025
severity: critical
tags:
    - deserialization
    - php
    - rce
    - owasp-a08
version: 1.0.0
---

# PHP Insecure Deserialization

## Overview
PHP's `unserialize()` function reconstructs PHP objects from a string representation. When called with user-controlled input, attackers can craft malicious serialized strings that:
- Invoke `__wakeup()` and `__destruct()` magic methods
- Chain gadgets from existing classes to achieve RCE
- Read/write arbitrary files

This is the basis of PHP Object Injection attacks.

## Remediation
- Never call `unserialize()` on user input
- Use `json_decode()` for data exchange
- If deserialization is required, validate with `allowed_classes` option

**Safe:**
```php
$data = json_decode($_POST['data'], true); // Safe alternative
// OR if PHP object needed, restrict classes:
$obj = unserialize($data, ['allowed_classes' => ['SafeClass']]);
```
