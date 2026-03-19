---
name: Insecure Network Socket
version: 1.0.0
description: Detects plain TCP sockets without TLS wrapping used for sensitive network communication.
tags: [network, socket, tls, owasp-a02]
languages: [python, javascript, typescript, go, java, c, cpp]
severity: high
confidence: medium
cwe: [CWE-311, CWE-319]
owasp: [A02:2025]
---

# Insecure Network Socket

## Overview
Plain TCP sockets transmit data unencrypted. Any network observer can intercept credentials, session tokens, and sensitive data. Applications should use TLS-wrapped sockets for all sensitive communication.

## Detection Strategy
- `socket.socket()` in Python without ssl.wrap_socket()
- `net.Socket` in Node.js instead of `tls.connect()`
- `net.Dial("tcp", ...)` in Go instead of `tls.Dial`

## Remediation
Wrap all sockets in TLS using the appropriate library:
- Python: `ssl.create_default_context()` + `context.wrap_socket()`
- Node.js: `tls.connect()` or `https`
- Go: `tls.Dial()` or `crypto/tls`
