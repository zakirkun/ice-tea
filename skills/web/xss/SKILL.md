---
confidence: medium
cwe:
    - CWE-79
description: "Detects DOM-based XSS vulnerabilities by tracing user input to dangerous sinks like innerHTML, eval, and document.write. Use when reviewing JavaScript or TypeScript for XSS risks, unsafe DOM manipulation, or security audit."
languages:
    - javascript
    - typescript
    - kotlin
    - dart
    - zig
    - elixir
name: cross-site-scripting-xss-detection
owasp:
    - A03:2025
severity: high
tags:
    - xss
    - web
    - injection
    - owasp-a03
version: 1.0.0
---

# Cross-Site Scripting (XSS)

Scans frontend code for dangerous DOM manipulations where user-controlled input reaches execution sinks. See `patterns.yaml` for detection rules.

## Detection Workflow

1. **Identify sources** — URL parameters (`location.search`, `location.hash`), form inputs, `postMessage` data
2. **Trace to sinks** — `innerHTML`, `outerHTML`, `document.write()`, `eval()`, `setTimeout(string)`
3. **Check sanitization** — verify DOMPurify or equivalent sits between source and sink
4. **Report** — file, line, sink type, source, and suggested fix

### Vulnerable

```javascript
element.innerHTML = userInput;
document.write(location.search);
```

### Safe

```javascript
element.textContent = userInput;
element.innerHTML = DOMPurify.sanitize(userInput);
```
