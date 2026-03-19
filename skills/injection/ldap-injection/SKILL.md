---
name: LDAP Injection
version: 1.0.0
description: Detects LDAP queries constructed from user input without proper escaping, enabling authentication bypass and data exfiltration.
tags: [ldap, injection, auth, owasp-a03]
languages: [java, python, javascript, typescript, php, go]
severity: high
confidence: medium
cwe: [CWE-90]
owasp: [A03:2025]
---

# LDAP Injection

## Overview
LDAP injection occurs when user-controlled input is incorporated into LDAP search filters without sanitization. Attackers can:
- Bypass authentication by injecting `*)(uid=*))(|(uid=*`
- Enumerate directory entries
- Access unauthorized data

## Detection Strategy
Look for LDAP filter strings built via string concatenation or format strings containing request parameters.

## Remediation
Use LDAP escaping functions to sanitize input before including in filters.

**Vulnerable (Java):**
```java
String filter = "(&(uid=" + username + ")(userPassword=" + password + "))";
DirContext ctx = new InitialDirContext(env);
ctx.search(baseDN, filter, controls);
```

**Safe (Java):**
```java
String filter = "(&(uid={0})(userPassword={1}))";
// Use parameterized search with MessageFormat or escaping
String safeFilter = "(&(uid=" + LdapEncoder.encode(username) + "))";
```
