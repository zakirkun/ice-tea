---
confidence: high
cwe:
    - CWE-22
description: Detects insecure ZIP/TAR extraction that does not validate entry paths, allowing directory traversal outside the extraction target.
languages:
    - java
    - python
    - javascript
    - typescript
    - go
    - php
    - kotlin
    - dart
    - zig
    - elixir
name: Zip Slip (Archive Path Traversal)
owasp:
    - A01:2025
severity: high
tags:
    - zip-slip
    - path-traversal
    - filesystem
    - owasp-a01
version: 1.0.0
---

# Zip Slip (Archive Path Traversal)

## Overview
Zip Slip is a directory traversal vulnerability in archive extraction. Malicious archives can contain entry names like `../../etc/cron.d/evil` or `../webroot/shell.php`. When extracted without path validation, files are written outside the intended directory, potentially achieving:
- Remote Code Execution (writing to cron, web shell in webroot)
- Configuration overwrite
- Sensitive file replacement

This affects ZIP, TAR, JAR, WAR, and other archive formats.

## Detection Strategy
- `ZipFile.extractall()` in Python without path validation
- `unzip.extract()` in Java without checking entry name for `..`
- `archive.ExtractAll()` in Go without normalized path check

## Remediation
Always normalize and validate the target path before writing each archive entry.

**Vulnerable (Java):**
```java
ZipEntry entry = zipFile.getEntry(name);
File dest = new File(targetDir, entry.getName());
// entry.getName() could be "../../etc/cron.d/evil"
```

**Safe (Java):**
```java
File dest = new File(targetDir, entry.getName()).getCanonicalFile();
if (!dest.toPath().startsWith(targetDir.toPath())) {
    throw new IOException("Zip Slip detected: " + entry.getName());
}
```
