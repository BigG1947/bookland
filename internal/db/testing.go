package db

import (
	"database/sql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/sqlite3"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"testing"
)

func NewTestSQLiteDB(t *testing.T) *sql.DB {
	t.Helper()

	log.SetFlags(log.Lshortfile)

	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		t.Fatal()
	}

	if err = db.Ping(); err != nil {
		t.Fatal()
	}

	if _, err := db.Exec("PRAGMA foreign_keys = ON;"); err != nil {
		DropTestSQLiteDB(t)
		t.Fatal()
	}

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		t.Fatal()
	}

	m, err := migrate.NewWithDatabaseInstance("file://../db/migrations", "test.db", driver)
	if err != nil {
		log.Printf("%s\n", err)
		t.Fatal()
	}

	if err = m.Up(); err != nil {
		t.Fatal()
	}

	if _, err = db.Exec(testData); err != nil {
		t.Fatal()
	}

	return db
}

func DropTestSQLiteDB(t *testing.T) {
	t.Helper()

	if err := os.Remove("test.db"); err != nil {
		t.Fatal()
	}
}

var testData = `
	INSERT INTO genre(id, name) VALUES (1, 'test_genre');
	INSERT INTO genre(id, name) VALUES (2, 'test_genre 2');
	INSERT INTO author(id, last_name, first_name, birthday, bio) VALUES (1, 'Potter', 'Harry', '03.12.1968', 'bio');
	INSERT INTO author(id, last_name, first_name, birthday, bio) VALUES (2, 'Laurence', 'Freddy', '03.12.1982', 'bio');
	INSERT INTO book(id, name, released, coast, pages, poster, author_id, genre_id) VALUES (1, 'test book 1', '03.12.2019', 300, 150, 'img.png', 1, 1);
	INSERT INTO book(id, name, released, coast, pages, poster, author_id, genre_id) VALUES (2, 'test book 2', '03.12.2019', 300, 150, 'img.png', 1, 1);
	INSERT INTO book(id, name, released, coast, pages, poster, author_id, genre_id) VALUES (3, 'test book 3', '03.12.2019', 300, 150, 'img.png', 1, 1);
	INSERT INTO book(id, name, released, coast, pages, poster, author_id, genre_id) VALUES (4, 'test book 4', '03.12.2019', 300, 150, 'img.png', 1, 1);
	INSERT INTO book(id, name, released, coast, pages, poster, author_id, genre_id) VALUES (5, 'test book 5', '03.12.2019', 300, 150, 'img.png', 1, 1);
	INSERT INTO book(id, name, released, coast, pages, poster, author_id, genre_id) VALUES (6, 'test book 6', '03.12.2019', 300, 150, 'img.png', 1, 1);
	INSERT INTO book(id, name, released, coast, pages, poster, author_id, genre_id) VALUES (7, 'test book 7', '03.12.2019', 300, 150, 'img.png', 1, 1);
	INSERT INTO book(id, name, released, coast, pages, poster, author_id, genre_id) VALUES (8, 'test book 8', '03.12.2019', 300, 150, 'img.png', 1, 1);
	INSERT INTO book(id, name, released, coast, pages, poster, author_id, genre_id) VALUES (9, 'test book 9', '03.12.2019', 300, 150, 'img.png', 1, 1);
	INSERT INTO book(id, name, released, coast, pages, poster, author_id, genre_id) VALUES (10, 'test book 10', '03.12.2019', 300, 150, 'img.png', 1, 1);
	INSERT INTO book(id, name, released, coast, pages, poster, author_id, genre_id) VALUES (11, 'test book 11', '03.12.2019', 300, 150, 'img.png', 2, 2);
	INSERT INTO book(id, name, released, coast, pages, poster, author_id, genre_id) VALUES (12, 'test book 12', '03.12.2019', 300, 150, 'img.png', 2, 2);
	INSERT INTO book(id, name, released, coast, pages, poster, author_id, genre_id) VALUES (13, 'test book 13', '03.12.2019', 300, 150, 'img.png', 2, 2);
	INSERT INTO book(id, name, released, coast, pages, poster, author_id, genre_id) VALUES (14, 'test book 14', '03.12.2019', 300, 150, 'img.png', 2, 2);
	INSERT INTO book(id, name, released, coast, pages, poster, author_id, genre_id) VALUES (15, 'test book 15', '03.12.2019', 300, 150, 'img.png', 2, 2);
	INSERT INTO book(id, name, released, coast, pages, poster, author_id, genre_id) VALUES (16, 'test book 16', '03.12.2019', 300, 150, 'img.png', 2, 2);
`
