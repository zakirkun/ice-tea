---
confidence: medium
cwe:
    - CWE-400
description: Detects HTTP clients and network connections without timeout configuration, enabling slowloris and resource exhaustion attacks.
languages:
    - python
    - javascript
    - typescript
    - go
    - java
    - php
    - kotlin
    - dart
    - zig
    - elixir
name: Missing Network Timeout Configuration
owasp:
    - A04:2025
severity: medium
tags:
    - network
    - timeout
    - dos
    - owasp-a04
version: 1.0.0
---

# Missing Network Timeout

## Overview
Network connections without timeouts are vulnerable to:
- **Slowloris**: Attacker opens connections and sends headers slowly, exhausting connection pool
- **Resource exhaustion**: Long-running connections hold goroutines/threads indefinitely
- **Hanging requests**: A slow external service stalls all users waiting for responses

## Remediation
- Set connect timeout, read timeout, and write timeout independently
- Use context with deadline for all outbound HTTP requests
- Configure server read/write timeouts

**Safe (Go):**
```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
```
