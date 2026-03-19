// INTENTIONALLY VULNERABLE — FOR SECURITY TESTING ONLY
// DO NOT USE IN PRODUCTION
package main

import (
	"crypto/md5"    // VULN: Weak hash algorithm
	"crypto/tls"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"

	jwt "github.com/golang-jwt/jwt/v4"
	_ "github.com/mattn/go-sqlite3"
)

// VULN: Hardcoded JWT secret key (CWE-798)
const jwtSecret = "supersecret123"

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, _ = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY,
		username TEXT UNIQUE,
		password TEXT,
		email TEXT,
		role TEXT DEFAULT 'user'
	)`)

	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/users", getUserHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/file", fileHandler)
	http.HandleFunc("/fetch", fetchHandler)
	http.HandleFunc("/redirect", redirectHandler)
	http.HandleFunc("/profile", profileHandler)

	log.Println("Starting vulnerable server on :8080")
	// VULN: HTTP (no TLS) server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// loginHandler - VULN: Insecure cookie, no CSRF, no rate limiting
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	// VULN: SQL injection (CWE-89) — string concatenation
	query := "SELECT id, role FROM users WHERE username = '" + username + "' AND password = '" + password + "'"
	row := db.QueryRow(query)

	var userID int
	var role string
	if err := row.Scan(&userID, &role); err != nil {
		// VULN: Verbose error — exposes SQL query structure (CWE-209)
		http.Error(w, fmt.Sprintf("Login failed: %s (query: %s)", err.Error(), query), http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"role":    role,
	})
	tokenStr, _ := token.SignedString([]byte(jwtSecret))

	// VULN: Cookie without Secure, HttpOnly, or SameSite (CWE-614)
	http.SetCookie(w, &http.Cookie{
		Name:  "session",
		Value: tokenStr,
	})

	fmt.Fprintf(w, `{"token": "%s"}`, tokenStr)
}

// registerHandler - VULN: Weak hash (MD5) for passwords (CWE-916)
func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")

	// VULN: MD5 password hashing (CWE-327, CWE-916)
	hash := md5.Sum([]byte(password))
	hashedPassword := fmt.Sprintf("%x", hash)

	_, err := db.Exec("INSERT INTO users (username, password, email) VALUES (?, ?, ?)",
		username, hashedPassword, email)
	if err != nil {
		// VULN: Exposes DB error to client (CWE-209)
		http.Error(w, "DB error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, `{"message": "User %s registered"}`, username)
}

// getUserHandler - VULN: SQL injection via query parameter (CWE-89)
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	username := r.URL.Query().Get("username")

	var query string
	if userID != "" {
		// VULN: Direct interpolation into SQL
		query = fmt.Sprintf("SELECT id, username, email, role FROM users WHERE id = %s", userID)
	} else if username != "" {
		// VULN: String concatenation SQL injection
		query = "SELECT id, username, email, role FROM users WHERE username = '" + username + "'"
	} else {
		query = "SELECT id, username, email, role FROM users"
	}

	rows, err := db.Query(query)
	if err != nil {
		// VULN: SQL error sent to client
		http.Error(w, "Query error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	result := "["
	for rows.Next() {
		var id int
		var uname, email, role string
		_ = rows.Scan(&id, &uname, &email, &role)
		result += fmt.Sprintf(`{"id":%d,"username":"%s","email":"%s","role":"%s"},`, id, uname, email, role)
	}
	result += "]"
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, result)
}

// pingHandler - VULN: Command injection (CWE-78)
func pingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	host := r.FormValue("host")

	// VULN: User input directly passed to exec.Command shell (CWE-78)
	cmd := exec.Command("sh", "-c", "ping -c 4 "+host)
	output, err := cmd.CombinedOutput()
	if err != nil {
		// VULN: Error with full command output exposed
		http.Error(w, fmt.Sprintf("Ping failed: %s\nOutput: %s", err.Error(), string(output)), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "<pre>%s</pre>", string(output))
}

// fileHandler - VULN: Path traversal (CWE-22)
func fileHandler(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("name")

	// VULN: No path sanitization — allows ../../etc/passwd
	path := "/var/www/files/" + filename

	data, err := os.ReadFile(path)
	if err != nil {
		// VULN: Exposes filesystem paths in error
		http.Error(w, fmt.Sprintf("Cannot read file %s: %v", path, err), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write(data)
}

// fetchHandler - VULN: Server-Side Request Forgery (CWE-918)
func fetchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	targetURL := r.FormValue("url")

	// VULN: TLS verification disabled + SSRF with no URL validation
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // VULN: CWE-295
	}
	client := &http.Client{Transport: tr}

	// VULN: Any URL accepted including internal/cloud metadata endpoints
	resp, err := client.Get(targetURL)
	if err != nil {
		http.Error(w, "Fetch error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	w.Write(body)
}

// redirectHandler - VULN: Open redirect (CWE-601)
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	redirectURL := r.URL.Query().Get("url")

	// VULN: No validation of redirect target URL
	http.Redirect(w, r, redirectURL, http.StatusFound)
}

// profileHandler - VULN: IDOR / Broken Object Level Auth (CWE-284)
func profileHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")

	// VULN: No authentication check, no ownership verification
	query := fmt.Sprintf("SELECT id, username, email, role FROM users WHERE id = %s", userID)
	row := db.QueryRow(query)

	var id int
	var username, email, role string
	if err := row.Scan(&id, &username, &email, &role); err != nil {
		http.Error(w, "User not found: "+err.Error(), http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, `{"id":%d,"username":"%s","email":"%s","role":"%s"}`, id, username, email, role)
}
