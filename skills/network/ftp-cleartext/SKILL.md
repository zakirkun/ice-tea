---
name: FTP Cleartext Communication
version: 1.0.0
description: Detects use of FTP protocol which transmits credentials and data in cleartext.
tags: [network, ftp, cleartext, owasp-a02]
languages: [python, javascript, typescript, go, java, php]
severity: high
confidence: high
cwe: [CWE-319]
owasp: [A02:2025]
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
