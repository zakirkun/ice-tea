---
name: Format String Vulnerability
version: 1.0.0
description: Detects format string vulnerabilities where user-controlled input is passed directly as the format argument to printf-family functions.
tags: [format-string, memory, c, cpp, owasp-a06]
languages: [c, cpp]
severity: critical
confidence: high
cwe: [CWE-134]
owasp: [A06:2025]
---

# Format String Vulnerability

## Overview
Format string vulnerabilities occur when user-controlled data is passed as the format argument to `printf()`, `sprintf()`, `fprintf()`, etc. Attackers can:
- **Read stack memory**: `%x %x %x %x` dumps stack values, leaking addresses and secrets
- **Write to arbitrary memory**: `%n` writes the number of bytes printed so far to a pointer on the stack
- **Remote Code Execution**: By writing to the GOT (Global Offset Table) or return address

## Detection Strategy
Any `printf`-family call where the first argument (format string) comes from user input rather than a string literal.

## Remediation
Always use a literal format string with user data as a parameter argument.

**Vulnerable:**
```c
char buf[256];
fgets(buf, sizeof(buf), stdin);
printf(buf);  // Format string vulnerability!
```

**Safe:**
```c
char buf[256];
fgets(buf, sizeof(buf), stdin);
printf("%s", buf);  // User data as argument, not format string
```
