---
name: Java Insecure Deserialization (ObjectInputStream)
version: 1.0.0
description: Detects Java ObjectInputStream.readObject() on untrusted data, enabling remote code execution via gadget chains.
tags: [deserialization, java, rce, owasp-a08]
languages: [java]
severity: critical
confidence: high
cwe: [CWE-502]
owasp: [A08:2025]
---

# Java Insecure Deserialization

## Overview
Java's `ObjectInputStream.readObject()` deserializes objects from byte streams. When called with untrusted data, attackers can use "gadget chains" from libraries like Commons Collections, Spring, and Apache to achieve Remote Code Execution.

Famous exploits: Apache Struts, Jenkins, WebLogic, JBoss all suffered RCE from this vulnerability.

## Remediation
- Replace Java serialization with JSON/XML with schema validation
- Use `ObjectInputFilter` (Java 9+) to restrict deserializable classes
- Use serialization filtering frameworks like SerialKiller
