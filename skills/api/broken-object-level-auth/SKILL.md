---
name: Broken Object Level Authorization (BOLA / IDOR)
version: 1.0.0
description: Detects API endpoints that use user-supplied IDs to access objects without verifying the requesting user owns or has permission to access that object.
tags: [bola, idor, api, auth, owasp-api1]
languages: [javascript, typescript, python, go, java, php]
severity: critical
confidence: medium
cwe: [CWE-284, CWE-639]
owasp: [A01:2025]
---

# Broken Object Level Authorization (BOLA / IDOR)

## Overview
BOLA (also called IDOR — Insecure Direct Object Reference) is the #1 API vulnerability. It occurs when an API endpoint uses a user-supplied ID (from URL, query parameter, or request body) to access an object without verifying that the current user is authorized to access that specific object.

Example: `GET /api/v1/users/{userId}/orders` — if userId from the path is trusted without checking that the authenticated user matches userId, any user can access any other user's orders.

## Detection Strategy
- Route handlers that use URL path parameters (`:id`, `:userId`) to query the database directly without an ownership check
- No comparison between the authenticated user's ID and the requested resource's owner ID
- Use of `findById()` or `getById()` without a tenant/user scope filter

## Remediation
Always scope database queries to the authenticated user's context.

**Vulnerable (Express.js):**
```js
app.get('/api/orders/:orderId', auth, async (req, res) => {
    const order = await Order.findById(req.params.orderId); // No ownership check!
    res.json(order);
});
```

**Safe:**
```js
app.get('/api/orders/:orderId', auth, async (req, res) => {
    const order = await Order.findOne({ _id: req.params.orderId, userId: req.user.id });
    if (!order) return res.status(404).json({ error: 'Not found' });
    res.json(order);
});
```
