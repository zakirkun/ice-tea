---
name: Unvalidated Redirects and Forwards
version: 1.0.0
description: Detects HTTP redirects to user-controlled URLs, enabling phishing and server-side request forgery.
tags: [redirect, phishing, web, owasp-a01]
languages: [generic]
severity: medium
confidence: medium
cwe: [CWE-601]
owasp: [A01:2025]
---

# Unvalidated Redirects and Forwards

## Overview
Detects HTTP redirects to user-controlled URLs, enabling phishing and server-side request forgery.

## Remediation
Validate redirect URLs. Avoid using user input directly; map input to internal routing enums instead.
