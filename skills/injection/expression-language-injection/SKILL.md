---
name: Expression Language Injection (EL Injection)
version: 1.0.0
description: Detects user-controlled input evaluated by Java EL, Spring SpEL, Thymeleaf, or JSP Expression Language.
tags: [injection, el-injection, spel, owasp-a03]
languages: [java]
severity: critical
confidence: high
cwe: [CWE-917]
owasp: [A03:2025]
---

# Expression Language Injection

## Overview
Expression Language (EL) injection allows attackers to execute arbitrary expressions in server-side template engines. Famous CVEs include Spring4Shell (CVE-2022-22965) and Struts RCE vulnerabilities.

Affected technologies:
- Spring SpEL: `#{T(java.lang.Runtime).getRuntime().exec('id')}`
- Thymeleaf: `__${T(java.lang.Runtime).getRuntime().exec('id')}__`
- JSP EL: `${Runtime.exec('id')}`

## Remediation
- Never pass user input to `ExpressionParser.parseExpression()`
- Use `SimpleEvaluationContext` instead of `StandardEvaluationContext`
- Disable SpEL evaluation in Thymeleaf templates when not needed
