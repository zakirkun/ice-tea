# INTENTIONALLY VULNERABLE — FOR SECURITY TESTING ONLY
# DO NOT USE IN PRODUCTION

import hashlib
import os
import pickle
import sqlite3
import subprocess
import xml.etree.ElementTree as ET  # VULN: XXE vulnerable parser (CWE-611)

from flask import Flask, jsonify, make_response, redirect, render_template_string, request

app = Flask(__name__)
# VULN: Hardcoded secret key (CWE-798)
app.secret_key = "supersecret123"

DB_PATH = "users.db"


def get_db():
    conn = sqlite3.connect(DB_PATH)
    conn.execute(
        "CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, username TEXT, password TEXT, email TEXT, role TEXT DEFAULT 'user')"
    )
    return conn


# VULN: CORS misconfiguration — allows all origins with credentials (CWE-942)
@app.after_request
def add_cors_headers(response):
    response.headers["Access-Control-Allow-Origin"] = "*"
    response.headers["Access-Control-Allow-Credentials"] = "true"
    return response


# VULN: Server-Side Template Injection via render_template_string (CWE-94)
# Attacker: GET /greet?name={{config.SECRET_KEY}}
# Attacker: GET /greet?name={{''.__class__.mro()[1].__subclasses__()}}
@app.route("/greet")
def greet():
    name = request.args.get("name", "World")
    # VULN: f-string in render_template_string — SSTI
    template = f"<h1>Hello, {name}!</h1>"
    return render_template_string(template)


# VULN: SQL Injection (CWE-89)
@app.route("/users")
def get_users():
    username = request.args.get("username", "")
    conn = get_db()

    # VULN: String concatenation in SQL query
    query = "SELECT id, username, email, role FROM users WHERE username LIKE '%" + username + "%'"

    try:
        rows = conn.execute(query).fetchall()
        return jsonify([{"id": r[0], "username": r[1], "email": r[2], "role": r[3]} for r in rows])
    except Exception as e:
        # VULN: DB error exposed in response (CWE-209)
        return jsonify({"error": str(e), "query": query}), 500
    finally:
        conn.close()


# VULN: Insecure Deserialization with pickle (CWE-502)
# Attacker can send crafted pickle payload that executes arbitrary code
@app.route("/restore", methods=["POST"])
def restore_session():
    data = request.get_data()
    # VULN: Deserializing untrusted pickle data — RCE
    obj = pickle.loads(data)
    return jsonify({"restored": str(obj)})


# VULN: Open Redirect (CWE-601)
@app.route("/redirect")
def open_redirect():
    url = request.args.get("url", "/")
    # VULN: No validation of redirect URL
    return redirect(url)


# VULN: Path Traversal (CWE-22)
@app.route("/read")
def read_file():
    filename = request.args.get("file", "")
    # VULN: No path sanitization
    path = "/var/www/static/" + filename
    try:
        with open(path) as f:
            return f.read()
    except Exception as e:
        # VULN: Exposes filesystem path in error
        return jsonify({"error": str(e), "path": path}), 404


# VULN: Command Injection (CWE-78)
@app.route("/exec", methods=["POST"])
def execute_command():
    host = request.form.get("host", "")
    # VULN: shell=True with user input
    result = subprocess.check_output("ping -c 3 " + host, shell=True, stderr=subprocess.STDOUT)
    return result.decode()


# VULN: Weak hash (MD5) for passwords (CWE-916)
@app.route("/register", methods=["POST"])
def register():
    username = request.form.get("username")
    password = request.form.get("password")
    email = request.form.get("email")

    # VULN: MD5 password hash without salt
    hashed = hashlib.md5(password.encode()).hexdigest()

    conn = get_db()
    try:
        conn.execute(
            "INSERT INTO users (username, password, email) VALUES (?, ?, ?)",
            (username, hashed, email),
        )
        conn.commit()
        return jsonify({"message": f"User {username} registered"})
    except Exception as e:
        return jsonify({"error": str(e)}), 500
    finally:
        conn.close()


# VULN: XXE Injection (CWE-611) — xml.etree.ElementTree is vulnerable
@app.route("/parse-xml", methods=["POST"])
def parse_xml():
    xml_data = request.get_data(as_text=True)
    # VULN: No DTD/entity protection
    try:
        root = ET.fromstring(xml_data)
        names = [child.text for child in root.iter("name")]
        return jsonify({"names": names})
    except Exception as e:
        return jsonify({"error": str(e)}), 400


# VULN: Sensitive data logging (CWE-532)
@app.route("/login", methods=["POST"])
def login():
    username = request.form.get("username")
    password = request.form.get("password")

    # VULN: Password logged in cleartext
    app.logger.info(f"Login attempt: username={username} password={password}")

    conn = get_db()
    hashed = hashlib.md5(password.encode()).hexdigest()
    row = conn.execute(
        "SELECT id, role FROM users WHERE username=? AND password=?", (username, hashed)
    ).fetchone()
    conn.close()

    if row:
        resp = make_response(jsonify({"message": "Login successful", "role": row[1]}))
        # VULN: Session cookie without Secure or HttpOnly (CWE-614)
        resp.set_cookie("session_id", str(row[0]))
        return resp
    return jsonify({"error": "Invalid credentials"}), 401


if __name__ == "__main__":
    # VULN: Debug mode enabled — exposes interactive debugger (CWE-215)
    app.run(debug=True, host="0.0.0.0", port=5000)
