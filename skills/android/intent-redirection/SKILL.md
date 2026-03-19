---
confidence: medium
cwe:
    - CWE-940
description: Detects Android intent redirection vulnerabilities where attacker-controlled intents are re-sent or forwarded, enabling privilege escalation.
languages:
    - java
    - kotlin
    - dart
    - zig
    - elixir
name: Android Intent Redirection
owasp:
    - A01:2025
severity: high
tags:
    - android
    - mobile
    - intent
    - owasp-m1
version: 1.0.0
---

# Android Intent Redirection

## Overview
Intent redirection occurs when an exported component receives an Intent from an untrusted source and forwards it to another component (potentially a private one). This allows attackers to:
- Access private Activities, Services, or ContentProviders
- Bypass permission requirements
- Escalate privileges

Common vulnerable pattern:
```java
// Exported component receives intent with embedded redirect
Intent redirectIntent = (Intent) getIntent().getParcelableExtra("extra_intent");
startActivity(redirectIntent); // Attacker controls redirectIntent!
```

## Remediation
- Validate Intent targets before forwarding
- Reject Intents with `FLAG_GRANT_READ_URI_PERMISSION` or `FLAG_GRANT_WRITE_URI_PERMISSION` if not expected
- Do not forward attacker-controlled Intents to sensitive components
