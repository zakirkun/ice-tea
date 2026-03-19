---
confidence: medium
cwe:
    - CWE-287
    - CWE-601
    - CWE-352
description: Detects common OAuth 2.0 implementation mistakes including missing state parameter, open redirect in redirect_uri, and token exposure.
languages:
    - javascript
    - typescript
    - python
    - go
    - java
    - php
    - ruby
    - kotlin
    - dart
    - zig
    - elixir
name: OAuth Misconfiguration
owasp:
    - A07:2025
severity: high
tags:
    - oauth
    - auth
    - csrf
    - owasp-a07
version: 1.0.0
---

# OAuth Misconfiguration

## Overview
OAuth 2.0 implementations are prone to several security issues:
1. **Missing `state` parameter**: Makes the OAuth flow vulnerable to CSRF attacks
2. **Open redirect in `redirect_uri`**: Allows token theft by redirecting to attacker-controlled URL
3. **Implicit flow usage**: The OAuth implicit flow exposes tokens in URL fragments
4. **Token in URL**: Access tokens logged in server logs or browser history
5. **Client secret exposure**: OAuth client secrets hardcoded or leaked

## Detection Strategy
- Check authorization URL construction for missing `state` parameter
- Detect use of deprecated implicit grant type (`response_type=token`)
- Find hardcoded client secrets
- Detect redirect_uri validation that allows wildcards or subdomains

## Remediation
- Always include a cryptographically random `state` parameter
- Use PKCE for public clients instead of implicit flow
- Validate `redirect_uri` against an exact allowlist
- Store client secrets in environment variables, never in code

**Vulnerable:**
```python
auth_url = f"https://auth.example.com/oauth/authorize?client_id={CLIENT_ID}&redirect_uri={uri}&response_type=token"
```

**Safe:**
```python
import secrets
state = secrets.token_urlsafe(32)
session['oauth_state'] = state
auth_url = f"https://auth.example.com/oauth/authorize?client_id={CLIENT_ID}&redirect_uri={FIXED_URI}&response_type=code&state={state}"
```
