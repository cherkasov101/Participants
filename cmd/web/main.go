package main

import (
	"log"
	"net/http"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "db/participants.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

	createTable()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/answers", answers)
	mux.HandleFunc("/participants", participants)
	mux.HandleFunc("/send", send)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func createTable() {
    createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		surname TEXT,
		patronymic TEXT,
		phone TEXT,
		email TEXT,
		section TEXT,
		birthdate TEXT,
		presentation TEXT,
		topic TEXT
	);
`
    _, err := db.Exec(createTableSQL)
    if err != nil {
        log.Fatal(err)
    }
}