---
name: GraphQL Field-Level Injection
version: 1.0.0
description: Detects GraphQL resolvers vulnerable to injection through unsanitized field arguments and dynamic query construction.
tags: [injection, graphql, sql-injection, owasp-a03]
languages: [javascript, typescript, python, java, go]
severity: high
confidence: medium
cwe: [CWE-89]
owasp: [A03:2025]
---

# GraphQL Field-Level Injection

## Overview
GraphQL resolvers that pass field arguments directly to database queries, OS commands, or other dangerous functions are vulnerable to injection. Unlike REST, GraphQL injection can be harder to detect because:
- Arguments arrive as parsed JavaScript objects (not raw strings)
- Multiple injection points per query
- Batch operations multiply impact

## Remediation
- Use parameterized queries in all database operations within resolvers
- Validate and sanitize all resolver arguments
- Use schema-level validation with type coercion
