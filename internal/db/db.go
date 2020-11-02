package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func NewSQLiteDB(dbName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec("PRAGMA foreign_keys = ON;"); err != nil {
		log.Fatal(err)
	}

	return db, nil
}
