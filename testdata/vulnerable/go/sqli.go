package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, username TEXT, password TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	db.Exec("INSERT INTO users (username, password) VALUES ('admin', 'supersecret')")
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("user")

	// VULNERABLE: Direct string concatenation leading to SQL injection
	query := "SELECT * FROM users WHERE username = '" + username + "'"

	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var u, p string
		rows.Scan(&id, &u, &p)
		fmt.Fprintf(w, "Found user: %s\n", u)
	}
}

func main() {
	initDB()
	http.HandleFunc("/user", GetUserHandler)
	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
