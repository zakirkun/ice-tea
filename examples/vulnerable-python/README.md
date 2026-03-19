# Vulnerable Python (Flask) Application — Ice Tea Test Target

> **WARNING: This application is intentionally vulnerable. DO NOT deploy in production.**

## Vulnerabilities Included

| ID | Vulnerability | CWE | Endpoint |
|----|--------------|-----|----------|
| 1 | Server-Side Template Injection (SSTI) | CWE-94 | `GET /greet` |
| 2 | SQL Injection | CWE-89 | `GET /users` |
| 3 | Insecure Deserialization (pickle) | CWE-502 | `POST /restore` |
| 4 | Open Redirect | CWE-601 | `GET /redirect` |
| 5 | CORS Misconfiguration | CWE-942 | All endpoints |
| 6 | Path Traversal | CWE-22 | `GET /read` |
| 7 | Command Injection | CWE-78 | `POST /exec` |
| 8 | Debug Mode Enabled | CWE-215 | `app.run(debug=True)` |
| 9 | Weak Hash (MD5) Password | CWE-916 | `POST /register` |
| 10 | XML External Entity | CWE-611 | `POST /parse-xml` |

## Running
```bash
pip install flask
python app.py
# Server starts on :5000
```
