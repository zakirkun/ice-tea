---
name: Broken Function Level Authorization
version: 1.0.0
description: Detects API endpoints that perform privileged operations without verifying the caller has the required role or permission.
tags: [bfla, api, auth, rbac, owasp-api5]
languages: [javascript, typescript, python, go, java, php]
severity: critical
confidence: medium
cwe: [CWE-285, CWE-862]
owasp: [A01:2025]
---

# Broken Function Level Authorization

## Overview
BFLA (API Security Top 10 #5) occurs when an API does not properly enforce which users can access which functions. Common manifestations:
- Admin-only endpoints accessible to regular users
- HTTP method confusion: `GET /api/users` is protected, `DELETE /api/users/{id}` is not
- Predictable admin paths: `/api/v1/admin/users` with no role check

## Detection Strategy
- Route handlers for admin operations (delete, ban, role-change) without role middleware
- Missing `isAdmin`, `hasRole()`, or `@PreAuthorize` checks on privileged endpoints

## Remediation
Apply role-based access control at every sensitive endpoint.

**Vulnerable:**
```js
app.delete('/api/admin/users/:id', authenticate, async (req, res) => {
    // No admin role check!
    await User.findByIdAndDelete(req.params.id);
});
```

**Safe:**
```js
app.delete('/api/admin/users/:id', authenticate, requireRole('admin'), async (req, res) => {
    await User.findByIdAndDelete(req.params.id);
});
```
