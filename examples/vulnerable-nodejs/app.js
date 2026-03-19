// INTENTIONALLY VULNERABLE — FOR SECURITY TESTING ONLY
// DO NOT USE IN PRODUCTION

const express = require('express');
const jwt = require('jsonwebtoken');
const axios = require('axios');
const { exec } = require('child_process');
const fs = require('fs');
const path = require('path');

const app = express();
app.use(express.json());
app.use(express.urlencoded({ extended: true }));
app.use(express.static('public'));

// VULN: Hardcoded JWT secret (CWE-798)
const JWT_SECRET = 'secret123';

// In-memory "database"
const users = [
  { id: 1, username: 'admin', password: 'admin123', role: 'admin' },
  { id: 2, username: 'user', password: 'password', role: 'user' },
];

// ─── 1. NoSQL-style injection via object comparison ───────────────────────────
// VULN: No mongo-sanitize, direct object comparison (CWE-943)
// Attacker sends: {"username": {"$gt": ""}, "password": {"$gt": ""}}
app.post('/login', (req, res) => {
  const { username, password } = req.body;

  // VULN: Password logged in plaintext (CWE-532)
  console.log(`Login attempt: username=${username} password=${password}`);

  // VULN: Comparing objects allows operator injection
  const user = users.find(u =>
    u.username == username && u.password == password
  );

  if (!user) {
    return res.status(401).json({ error: 'Invalid credentials' });
  }

  const token = jwt.sign({ id: user.id, role: user.role }, JWT_SECRET);

  // VULN: Cookie without httpOnly, secure, sameSite (CWE-614)
  res.cookie('session', token);
  res.json({ token });
});

// ─── 2. Prototype Pollution ───────────────────────────────────────────────────
// VULN: Recursive merge without __proto__ protection (CWE-1321)
function merge(target, source) {
  for (const key of Object.keys(source)) {
    // VULN: No check for __proto__, constructor, prototype
    if (typeof source[key] === 'object' && source[key] !== null) {
      target[key] = target[key] || {};
      merge(target[key], source[key]);
    } else {
      target[key] = source[key];
    }
  }
  return target;
}

app.post('/merge', (req, res) => {
  const base = {};
  // VULN: Merging user-controlled JSON into object
  const result = merge(base, req.body);
  res.json({ result });
});

// ─── 3. SSRF ──────────────────────────────────────────────────────────────────
// VULN: Fetches any user-supplied URL including internal/cloud metadata (CWE-918)
app.post('/fetch', async (req, res) => {
  const { url } = req.body;
  try {
    // VULN: No URL allowlist, follows redirects
    const response = await axios.get(url, { maxRedirects: 10 });
    res.json({ data: response.data });
  } catch (err) {
    // VULN: Exposes internal error details (CWE-209)
    res.status(500).json({ error: err.message, stack: err.stack });
  }
});

// ─── 4. Command Injection ─────────────────────────────────────────────────────
// VULN: User input passed to shell command (CWE-78)
app.post('/run', (req, res) => {
  const { host } = req.body;
  // VULN: Shell injection via exec with string concat
  exec('ping -c 3 ' + host, (error, stdout, stderr) => {
    if (error) {
      // VULN: Stack trace sent to client (CWE-209)
      return res.status(500).json({ error: error.message, stderr });
    }
    res.send(`<pre>${stdout}</pre>`);
  });
});

// ─── 5. Path Traversal ────────────────────────────────────────────────────────
// VULN: No path sanitization — allows ../../etc/passwd (CWE-22)
app.get('/file', (req, res) => {
  const filename = req.query.name;
  const filePath = path.join(__dirname, 'files', filename);
  
  try {
    const content = fs.readFileSync(filePath, 'utf8');
    res.send(content);
  } catch (err) {
    // VULN: Full path exposed in error
    res.status(404).json({ error: `File not found: ${filePath}` });
  }
});

// ─── 6. ReDoS ─────────────────────────────────────────────────────────────────
// VULN: Vulnerable regex with nested quantifiers applied to user input (CWE-1333)
const EMAIL_REGEX = /^([a-zA-Z0-9])(([a-zA-Z0-9])*([._-])?)*([a-zA-Z0-9])*@([a-zA-Z0-9])+\.([a-zA-Z]){2,6}$/;

app.get('/validate', (req, res) => {
  const email = req.query.email;
  // VULN: Catastrophic backtracking on crafted input
  const valid = EMAIL_REGEX.test(email);
  res.json({ valid });
});

// ─── 7. JWT None Algorithm Attack ─────────────────────────────────────────────
// VULN: Accepts JWT with any algorithm including 'none' (CWE-347)
app.get('/admin', (req, res) => {
  const token = req.headers.authorization?.replace('Bearer ', '');
  try {
    // VULN: algorithms not restricted — 'none' attack possible
    const decoded = jwt.verify(token, JWT_SECRET, { algorithms: ['HS256', 'none'] });
    if (decoded.role !== 'admin') {
      return res.status(403).json({ error: 'Forbidden' });
    }
    res.json({ message: 'Admin area', users });
  } catch (err) {
    res.status(401).json({ error: err.message });
  }
});

// ─── 8. Open Redirect ─────────────────────────────────────────────────────────
// VULN: Redirect to user-controlled URL (CWE-601)
app.get('/redirect', (req, res) => {
  const url = req.query.url;
  // VULN: No validation of redirect target
  res.redirect(url);
});

// ─── 9. IDOR / Broken Object Level Auth ───────────────────────────────────────
// VULN: Returns any user profile by ID without auth (CWE-639)
app.get('/users/:id', (req, res) => {
  const user = users.find(u => u.id == req.params.id);
  if (!user) return res.status(404).json({ error: 'Not found' });
  // VULN: Returns password hash too (information disclosure)
  res.json(user);
});

// ─── 10. XSS in public/app.js ─────────────────────────────────────────────────
// (See public/app.js)

app.listen(3000, () => {
  console.log('Vulnerable server running on http://localhost:3000');
});
