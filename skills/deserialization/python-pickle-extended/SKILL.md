---
name: Python Pickle/Shelve Extended Deserialization
version: 1.0.0
description: Detects extended use of pickle, shelve, and marshal for deserializing untrusted data beyond basic patterns.
tags: [deserialization, python, rce, owasp-a08]
languages: [python]
severity: critical
confidence: high
cwe: [CWE-502]
owasp: [A08:2025]
---

# Python Pickle Extended Deserialization

## Overview
Beyond basic `pickle.loads()`, Python has several other insecure deserialization mechanisms:
- `shelve` (backed by pickle)
- `marshal.loads()` (bytecode deserialization)
- `pickletools`
- HTTP APIs passing pickled data in cookies or request bodies

## Remediation
- Replace with `json`, `msgpack`, or `protobuf`
- If using shelve/pickle for caching, ensure data source is trusted and server-controlled
- Never accept serialized Python objects from HTTP clients
