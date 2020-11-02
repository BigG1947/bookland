package main

import (
	"bookland/internal/db"
	"bookland/internal/store"
	"log"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	conn, err := db.NewSQLiteDB("book.db")
	if err != nil {
		log.Fatalf("%s\n", err)
	}

	_ = store.NewStore(conn)
}
