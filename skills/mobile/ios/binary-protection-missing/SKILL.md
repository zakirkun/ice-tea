---
confidence: medium
cwe:
    - CWE-693
description: Detects iOS app configurations missing binary hardening features like ASLR, stack canaries, and ARC.
languages:
    - generic
    - kotlin
    - dart
    - zig
    - elixir
name: iOS Binary Protection Missing
owasp:
    - A08:2025
severity: medium
tags:
    - ios
    - mobile
    - binary-protection
    - owasp-m8
version: 1.0.0
---

# iOS Binary Protection Missing

## Overview
iOS binaries should be compiled with hardening flags:
- **PIE (ASLR)**: Randomizes memory addresses, making exploits harder
- **Stack Canaries**: Detects stack smashing before return address overwrite
- **ARC (Automatic Reference Counting)**: Reduces use-after-free vulnerabilities
- **Bitcode**: Allows Apple to re-optimize binary with future security improvements

Missing these protections makes reverse engineering and exploitation easier.

## Remediation
- Enable PIE: `-pie` linker flag
- Enable Stack Canaries: `-fstack-protector-all`
- Enable ARC in Objective-C projects
- Build with latest Xcode toolchain
