---
name: Server-Side Template Injection (SSTI)
version: 1.0.0
description: Detects user-controlled input passed directly into template rendering engines, enabling code execution.
tags: [ssti, template-injection, injection, owasp-a03]
languages: [python, javascript, typescript, java, php, ruby, go]
severity: critical
confidence: medium
cwe: [CWE-94]
owasp: [A03:2025]
---

# Server-Side Template Injection (SSTI)

## Overview
SSTI occurs when user input is embedded directly into a template string before rendering. Template engines like Jinja2, Twig, Pebble, Freemarker, and Mustache can interpret injected expressions as code, leading to Remote Code Execution (RCE).

Classic payload: `{{7*7}}` → if output is `49`, the template is evaluating user input.

## Detection Strategy
Look for template `render()` calls that concatenate user input directly into the template string rather than passing it as a context variable.

Key patterns:
- `render_template_string(user_input)` in Flask/Jinja2
- `new Template(user_input).render()` in various engines
- String format operations feeding user data into template strings

## Remediation
Never put user input in the template string. Always pass user data as template variables.

**Vulnerable (Python/Jinja2):**
```python
from flask import render_template_string, request
@app.route('/greet')
def greet():
    name = request.args.get('name')
    return render_template_string(f"Hello {name}!")  # SSTI!
```

**Safe (Python/Jinja2):**
```python
@app.route('/greet')
def greet():
    name = request.args.get('name')
    return render_template_string("Hello {{ name }}!", name=name)
```
