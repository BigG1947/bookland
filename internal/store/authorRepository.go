package store

import (
	"bookland/internal/models"
	"database/sql"
)

type AuthorRepository struct {
	db *sql.DB
}

func newAuthorRepository(db *sql.DB) *AuthorRepository {
	return &AuthorRepository{db: db}
}

func (ar *AuthorRepository) Get(id int) (*models.Author, error) {
	author := &models.Author{}
	row := ar.db.QueryRow("SELECT id, last_name, first_name, birthday, bio FROM author WHERE id = ?", id)
	err := row.Scan(&author.Id, &author.LastName, &author.FirstName, &author.BirthDay, &author.Bio)
	if err != nil {
		return nil, err
	}
	return author, nil
}

func (ar *AuthorRepository) Add(author *models.Author) error {
	res, err := ar.db.Exec(
		"INSERT INTO author(last_name, first_name, birthday, bio) VALUES (?, ?, ?, ?)",
		author.LastName, author.FirstName, author.BirthDay, author.Bio,
	)
	if err != nil {
		return err
	}

	if author.Id, err = res.LastInsertId(); err != nil {
		return err
	}

	return nil
}

func (ar *AuthorRepository) Update(author *models.Author) error {
	if _, err := ar.db.Exec(
		"UPDATE author SET last_name = ?, first_name = ?, birthday = ?, bio = ?",
		author.LastName, author.FirstName, author.BirthDay, author.Bio,
	); err != nil {
		return err
	}
	return nil
}

func (ar *AuthorRepository) Delete(id int) error {
	_, err := ar.db.Exec("DELETE FROM author WHERE id = ?", id)
	if err != nil {
		return err
	}
	return err
}

func (ar *AuthorRepository) Count() (int, error) {
	var count int
	if err := ar.db.QueryRow("SELECT COUNT(id) FROM author").Scan(&count); err != nil {
		return 0, nil
	}
	return count, nil
}

func (ar *AuthorRepository) SearchByName(value string) ([]models.Author, error) {
	value = "%" + value + "%"
	rows, err := ar.db.Query(
		"SELECT id, last_name, first_name, birthday, bio FROM author WHERE last_name LIKE ? OR first_name LIKE ?",
		value, value,
	)
	if err != nil {
		return nil, err
	}

	var authors []models.Author
	for rows.Next() {
		var a models.Author
		if err := rows.Scan(&a.Id, &a.LastName, &a.FirstName, &a.BirthDay, &a.Bio); err != nil {
			return nil, err
		}
		authors = append(authors, a)
	}

	return authors, nil
}

func (ar *AuthorRepository) GetPerPage(perPage int, page int) ([]models.Author, error) {
	start := (page - 1) * perPage
	rows, err := ar.db.Query(
		"SELECT id, last_name, first_name, birthday, bio FROM author ORDER BY last_name LIMIT ?, ?",
		start, perPage,
	)
	if err != nil {
		return nil, err
	}

	var authors []models.Author
	for rows.Next() {
		var a models.Author
		if err := rows.Scan(&a.Id, &a.LastName, &a.FirstName, &a.BirthDay, &a.Bio); err != nil {
			return nil, err
		}
		authors = append(authors, a)
	}

	return authors, nil
}
