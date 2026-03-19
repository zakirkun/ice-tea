---
name: Redis Without Authentication
version: 1.0.0
description: Detects Redis connections and configurations without authentication, allowing unauthenticated access to the cache.
tags: [database, redis, auth, owasp-a05]
languages: [javascript, typescript, python, go, java, php, ruby]
severity: critical
confidence: high
cwe: [CWE-306]
owasp: [A05:2025]
---

# Redis Without Authentication

## Overview
Redis instances running without authentication (`requirepass`) are exposed to unauthenticated access. Any client that can reach the Redis port can read, write, or delete all cached data — including session tokens, sensitive user data, and application secrets.

## Detection Strategy
Look for Redis client connections that do not supply a password, or Redis configuration files with `requirepass` commented out or absent.

## Remediation
- Set `requirepass <strong-password>` in `redis.conf`
- Use Redis ACL (Redis 6+) for fine-grained access control
- Bind Redis to localhost or a private network interface
- Use TLS for Redis connections in production

**Vulnerable (Node.js):**
```js
const client = redis.createClient({ host: 'redis-host', port: 6379 });
```

**Safe (Node.js):**
```js
const client = redis.createClient({ host: 'redis-host', port: 6379, password: process.env.REDIS_PASSWORD });
```
