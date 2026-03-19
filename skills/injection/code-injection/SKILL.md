---
name: Code Injection (eval / dynamic execution)
version: 1.0.0
description: Detects user-controlled input passed to code evaluation functions like eval, exec, or dynamic code generation.
tags: [injection, code-injection, rce, owasp-a03]
languages: [python, php, javascript, typescript, ruby]
severity: critical
confidence: high
cwe: [CWE-94]
owasp: [A03:2025]
---

# Code Injection

## Overview
Code injection occurs when user-controlled input is executed as code by the application. This differs from Command Injection (OS commands) and SQL Injection — the attacker's input is evaluated by the application's own interpreter.

Common functions to avoid with user input:
- Python: `eval()`, `exec()`, `compile()`
- PHP: `eval()`, `assert()` with string argument, `preg_replace()` with `/e`
- JavaScript: `eval()`, `new Function()`, `setTimeout(string)`
- Ruby: `eval()`, `binding.eval()`, `instance_eval`

## Remediation
- Never pass user input to code evaluation functions
- Use data structures and lookup tables instead of dynamic code evaluation
