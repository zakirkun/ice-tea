# Vulnerable Go Application — Ice Tea Test Target

> **WARNING: This application is intentionally vulnerable. DO NOT deploy in production.**

A vulnerable REST API written in Go (net/http + database/sql) designed to demonstrate and test ice-tea scanner detection capabilities.

## Vulnerabilities Included

| ID | Vulnerability | CWE | Endpoint |
|----|--------------|-----|----------|
| 1 | SQL Injection | CWE-89 | `GET /users?id=` |
| 2 | Command Injection | CWE-78 | `POST /ping` |
| 3 | Path Traversal | CWE-22 | `GET /file?name=` |
| 4 | SSRF | CWE-918 | `POST /fetch` |
| 5 | Hardcoded JWT Secret | CWE-798 | All auth endpoints |
| 6 | Insecure Cookie | CWE-614 | `POST /login` |
| 7 | Weak Crypto (MD5) | CWE-327 | `POST /register` |
| 8 | InsecureSkipVerify TLS | CWE-295 | `POST /webhook` |
| 9 | Open Redirect | CWE-601 | `GET /redirect` |
| 10 | Verbose Error | CWE-209 | All error paths |

## Running
```bash
go mod tidy
go run main.go
# Server starts on :8080
```
