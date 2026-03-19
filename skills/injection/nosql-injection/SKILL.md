---
name: NoSQL Injection
version: 1.0.0
description: Detects NoSQL query injection vulnerabilities in MongoDB, Redis, and other NoSQL databases.
tags: [nosql, injection, mongodb, owasp-a03]
languages: [javascript, typescript, python, go, java, php]
severity: high
confidence: medium
cwe: [CWE-943]
owasp: [A03:2025]
---

# NoSQL Injection

## Overview
NoSQL databases like MongoDB are vulnerable to injection attacks that exploit the query language structure. Unlike SQL injection, NoSQL injection often exploits JSON/object operators:
- `$where` clause with JavaScript execution
- Operator injection: `{ "username": { "$gt": "" } }` bypasses authentication
- `$regex` injection for data enumeration

## Detection Strategy
- MongoDB `$where` operator with user input (arbitrary JS execution)
- Direct use of request body as a query object without validation
- `findOne()` / `find()` with unvalidated user objects

## Remediation
- Never use `$where` with user input
- Sanitize input using `mongo-sanitize` or similar
- Use schema validation (Mongoose)
- Explicitly define expected query fields

**Vulnerable (Node.js):**
```js
const user = await User.findOne({ username: req.body.username, password: req.body.password });
// Attacker sends: {"username": {"$gt": ""}, "password": {"$gt": ""}}
```

**Safe (Node.js):**
```js
const mongoSanitize = require('express-mongo-sanitize');
app.use(mongoSanitize());
// OR manually: ensure fields are strings
const user = await User.findOne({ username: String(req.body.username) });
```
