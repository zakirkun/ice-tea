---
confidence: high
cwe:
    - CWE-200
    - CWE-400
description: Detects insecure GraphQL configurations including enabled introspection, unbounded query depth, and missing authentication on resolvers.
languages:
    - javascript
    - typescript
    - python
    - java
    - go
    - kotlin
    - dart
    - zig
    - elixir
name: GraphQL Security Issues
owasp:
    - A08:2025
severity: high
tags:
    - graphql
    - api
    - introspection
    - owasp-api8
version: 1.0.0
---

# GraphQL Security Issues

## Overview
GraphQL APIs have unique security risks:
1. **Introspection enabled in production**: Reveals entire schema to attackers for reconnaissance
2. **Unbounded query depth**: `{ user { friends { friends { friends { ... } } } } }` — DoS via deeply nested queries
3. **Batch query abuse**: Unlimited query batching enables brute force and DoS
4. **Missing resolver-level auth**: Even with authentication, each resolver must check permissions
5. **Verbose error messages**: GraphQL errors often expose internal details

## Remediation
- Disable introspection in production (`introspection: false`)
- Enforce query depth limits (max 5-7 levels)
- Implement query complexity limits
- Use persisted queries to restrict query set
- Add authentication checks in each resolver
