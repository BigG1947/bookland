package store

import "database/sql"

type store struct {
	Books *bookRepository
}

func NewStore(db *sql.DB) *store {
	return &store{
		Books: newBookRepository(db),
	}
}
