---
confidence: medium
cwe:
    - CWE-643
description: Detects XPath queries built from user input without escaping, enabling authentication bypass and data disclosure.
languages:
    - java
    - python
    - javascript
    - typescript
    - php
    - kotlin
    - dart
    - zig
    - elixir
name: XPath Injection
owasp:
    - A03:2025
severity: high
tags:
    - xpath
    - injection
    - xml
    - owasp-a03
version: 1.0.0
---

# XPath Injection

## Overview
XPath injection is analogous to SQL injection but targets XML documents. An attacker can manipulate XPath expressions to bypass authentication, access unauthorized data nodes, or cause denial of service.

Classic bypass: `' or '1'='1` in username field when query is `//user[name/text()='{username}']`

## Remediation
Use parameterized XPath (XQuery with bound variables) or escape user input using library-provided escaping functions.

**Vulnerable (Java):**
```java
String query = "//user[name='" + username + "' and password='" + password + "']";
NodeList nodes = (NodeList) xpath.evaluate(query, doc, XPathConstants.NODESET);
```

**Safe (Java):**
```java
// Use parameterized XQuery or properly escape the input
```
