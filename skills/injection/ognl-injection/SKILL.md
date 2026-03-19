---
confidence: high
cwe:
    - CWE-917
description: Detects Apache Struts and other OGNL-using frameworks vulnerable to expression injection through request parameters.
languages:
    - java
    - kotlin
    - dart
    - zig
    - elixir
name: OGNL Injection (Apache Struts)
owasp:
    - A03:2025
severity: critical
tags:
    - injection
    - ognl
    - struts
    - rce
    - owasp-a03
version: 1.0.0
---

# OGNL Injection

## Overview
Object-Graph Navigation Language (OGNL) is the expression language used by Apache Struts. OGNL injection has led to some of the most severe vulnerabilities in enterprise Java applications, including the Equifax breach (CVE-2017-5638).

Attackers inject OGNL expressions via HTTP parameters, content-type headers, or form fields to execute arbitrary Java code.

## Remediation
- Keep Apache Struts updated to latest version
- Use strict validation for all HTTP parameters
- Consider migrating from Struts to a safer framework
- Enable Struts security workarounds as documented in advisories
