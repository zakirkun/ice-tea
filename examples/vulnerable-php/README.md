# Vulnerable PHP Application — Ice Tea Test Target

> **WARNING: This application is intentionally vulnerable. DO NOT deploy in production.**

A classic PHP web application with multiple vulnerabilities for testing ice-tea scanner.

## Vulnerabilities Included

| ID | Vulnerability | CWE | File |
|----|--------------|-----|------|
| 1 | SQL Injection | CWE-89 | `login.php`, `search.php` |
| 2 | XSS (Reflected) | CWE-79 | `search.php`, `profile.php` |
| 3 | Local File Inclusion | CWE-98 | `index.php` |
| 4 | File Upload Bypass | CWE-434 | `upload.php` |
| 5 | CSRF | CWE-352 | `transfer.php` |
| 6 | XXE Injection | CWE-611 | `import.php` |
| 7 | Insecure Cookie | CWE-614 | `login.php` |
| 8 | Command Injection | CWE-78 | `tools.php` |
| 9 | Path Traversal | CWE-22 | `download.php` |
| 10 | Plaintext Password Storage | CWE-916 | `register.php` |

## Running
```bash
php -S localhost:8081 -t .
```
