---
name: Ruby Marshal Deserialization
version: 1.0.0
description: Detects Marshal.load() or Marshal.restore() on untrusted data, enabling code execution.
tags: [deserialization, ruby, rce, owasp-a08]
languages: [ruby]
severity: critical
confidence: high
cwe: [CWE-502]
owasp: [A08:2025]
---

# Ruby Marshal Deserialization

## Overview
Ruby's `Marshal.load()` deserializes a binary format that can instantiate arbitrary Ruby objects. Attackers can craft malicious payloads that invoke `initialize`, `[]`, and other methods during deserialization, leading to code execution.

This affected Rails versions using Marshal for cookie sessions before 4.x.

## Remediation
- Never use `Marshal.load()` on untrusted data
- Use `JSON.parse()` or `MessagePack` for data exchange
- For sessions, use signed/encrypted JSON-based session stores
