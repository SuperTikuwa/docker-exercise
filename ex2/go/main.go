package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "docker:docker@tcp(db:3306)/app_db")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var admin struct {
		ID       int
		Name     string
		Email    string
		Password string
	}

	if err := db.QueryRow("SELECT id, name, email, password FROM users limit 1;").Scan(&admin.ID, &admin.Name, &admin.Email, &admin.Password); err != nil {
		panic(err)
	}

	j, err := json.Marshal(admin)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(j))
}
