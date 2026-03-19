---
confidence: high
cwe:
    - CWE-502
description: Detects use of BinaryFormatter, SoapFormatter, and NetDataContractSerializer which are insecure for untrusted data.
languages:
    - generic
    - kotlin
    - dart
    - zig
    - elixir
name: .NET BinaryFormatter Deserialization
owasp:
    - A08:2025
severity: critical
tags:
    - deserialization
    - dotnet
    - csharp
    - rce
    - owasp-a08
version: 1.0.0
---

# .NET BinaryFormatter Deserialization

## Overview
`BinaryFormatter` and related .NET serializers are inherently insecure when used with untrusted data. Microsoft has officially deprecated `BinaryFormatter` in .NET 5+ due to its inability to be secured.

Gadget chains from ysoserial.net can achieve RCE through ViewState, remoting, and other attack surfaces.

## Remediation
- Use `System.Text.Json` or `Newtonsoft.Json` with schema validation
- Microsoft recommends `System.Text.Json` as the secure replacement
- If migration is not possible, validate/whitelist types before deserialization
