---
confidence: high
cwe:
    - CWE-319
description: Detects use of FTP protocol which transmits credentials and data in cleartext.
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
name: FTP Cleartext Communication
owasp:
    - A02:2025
severity: high
tags:
    - network
    - ftp
    - cleartext
    - owasp-a02
version: 1.0.0
---

# FTP Cleartext Communication

## Overview
FTP (File Transfer Protocol) transmits all data including credentials in plaintext. Network monitoring captures:
- Login credentials (USER/PASS commands)
- All file transfers
- Directory listings

FTP is deprecated and should be replaced with SFTP (SSH File Transfer Protocol) or FTPS (FTP over TLS).

## Remediation
- Use SFTP instead of FTP
- If FTP must be used, require explicit TLS (`AUTH TLS`) — FTPS
- Use `Paramiko` (Python), `ssh2` (Node.js), or native SSH libraries
