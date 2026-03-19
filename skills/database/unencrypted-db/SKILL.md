---
name: Unencrypted Database
version: 1.0.0
description: Detects database configurations and connections missing encryption at rest and in transit.
tags: [database, encryption, owasp-a02]
languages: [python, javascript, typescript, go, java, php, yaml]
severity: high
confidence: medium
cwe: [CWE-311]
owasp: [A02:2025]
---

# Unencrypted Database

## Overview
Databases storing sensitive data must be encrypted both at rest (disk encryption) and in transit (TLS). Without encryption at rest, physical access or cloud storage access exposes all data. Without TLS, credentials and data can be intercepted on the network.

## Remediation
- Enable TLS/SSL in database connections (`sslmode=require` for PostgreSQL)
- Use encrypted database volumes (AWS RDS encryption, cloud disk encryption)
- Use SQLCipher for SQLite encryption
