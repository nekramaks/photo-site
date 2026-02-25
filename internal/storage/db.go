package storage

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Init(path string) {
	var err error
	DB, err = sql.Open("sqlite", path)

	if err != nil {
		log.Fatal("no connect to db", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Db`s not answering", err)
	}

	createTables()
}

func createTables() {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		service TEXT NOT NULL UNIQUE,
		token TEXT NOT NULL UNIQUE,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
		CREATE TABLE IF NOT EXISTS photos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER NOT NULL,
		filename TEXT NOT NULL,
		FOREIGN KEY (event_id) REFERENCES events (id)
		);
	`)
	if err != nil {
		log.Fatal("error creating tables", err)
	}
}
