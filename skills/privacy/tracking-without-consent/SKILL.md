---
name: User Tracking Without Consent
version: 1.0.0
description: Detects user tracking mechanisms deployed without checking for consent, including session recording, heatmaps, and behavioral analytics.
tags: [privacy, tracking, gdpr, owasp-a05]
languages: [javascript, typescript]
severity: medium
confidence: medium
cwe: [CWE-359]
owasp: [A05:2025]
---

# User Tracking Without Consent

## Overview
Session recording tools (Hotjar, FullStory, LogRocket), heatmap tools, and behavioral analytics that collect user interactions may capture sensitive data entered by users. Loading these unconditionally without consent violates GDPR and may capture passwords, payment data, or health information.

## Remediation
- Integrate tracking scripts with consent management
- Configure tools to mask sensitive form fields
- Provide opt-out mechanism for tracking
