# Vulnerable Node.js (Express) Application — Ice Tea Test Target

> **WARNING: This application is intentionally vulnerable. DO NOT deploy in production.**

## Vulnerabilities Included

| ID | Vulnerability | CWE | Endpoint |
|----|--------------|-----|----------|
| 1 | Prototype Pollution | CWE-1321 | `POST /merge` |
| 2 | NoSQL Injection (MongoDB) | CWE-943 | `POST /login` |
| 3 | SSRF | CWE-918 | `POST /fetch` |
| 4 | JWT Weak Secret | CWE-798 | All auth endpoints |
| 5 | ReDoS | CWE-1333 | `GET /validate` |
| 6 | XSS via innerHTML | CWE-79 | `public/app.js` |
| 7 | Command Injection | CWE-78 | `POST /run` |
| 8 | Path Traversal | CWE-22 | `GET /file` |
| 9 | Insecure Cookie | CWE-614 | `POST /login` |
| 10 | Missing Rate Limiting | CWE-307 | `POST /login` |

## Running
```bash
npm install
node app.js
# Server starts on :3000
```
