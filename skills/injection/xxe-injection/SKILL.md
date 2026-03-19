---
name: XML External Entity (XXE) Injection
version: 1.0.0
description: Detects XML parsers configured to process external entities, enabling file disclosure and SSRF.
tags: [xxe, xml, injection, owasp-a05]
languages: [java, python, php, go, javascript, typescript, ruby]
severity: critical
confidence: high
cwe: [CWE-611]
owasp: [A05:2025]
---

# XML External Entity (XXE) Injection

## Overview
XXE occurs when an XML parser processes external entity references within XML input. Attackers can:
- Read arbitrary files from the server filesystem (`/etc/passwd`, private keys)
- Perform Server-Side Request Forgery (SSRF) to internal services
- Execute Denial of Service via "Billion Laughs" attack
- In some configurations, achieve Remote Code Execution

## Detection Strategy
Look for XML parsing operations that do not explicitly disable external entity resolution:
- Java: `DocumentBuilderFactory` without `setFeature("http://xml.org/sax/features/external-general-entities", false)`
- Python: `xml.etree.ElementTree`, `lxml`, `defusedxml` not used
- PHP: `simplexml_load_string()` / `DOMDocument` without `LIBXML_NOENT` disabled
- Go: Standard `encoding/xml` is safe, but `etree` or other libs may not be

## Remediation
Disable DTD processing and external entities entirely, or use a secure XML parsing library.

**Vulnerable (Java):**
```java
DocumentBuilderFactory dbf = DocumentBuilderFactory.newInstance();
DocumentBuilder db = dbf.newDocumentBuilder(); // XXE enabled by default!
```

**Safe (Java):**
```java
DocumentBuilderFactory dbf = DocumentBuilderFactory.newInstance();
dbf.setFeature("http://apache.org/xml/features/disallow-doctype-decl", true);
dbf.setFeature("http://xml.org/sax/features/external-general-entities", false);
dbf.setFeature("http://xml.org/sax/features/external-parameter-entities", false);
```
