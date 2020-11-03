package store

import (
	"bookland/internal/models"
	"database/sql"
)

type bookRepository struct {
	db *sql.DB
}

func newBookRepository(db *sql.DB) *bookRepository {
	return &bookRepository{db: db}
}

func (br *bookRepository) Add(b *models.Book) error {
	res, err := br.db.Exec(
		"INSERT INTO book(name, released, coast, pages, poster, author_id, genre_id) VALUES (?, ?, ?, ?, ?, ?, ?)",
		b.Name, b.Release, b.Coast, b.Pages, b.PosterURL, b.AuthorId, b.GenreId,
	)
	if err != nil {
		return err
	}
	if b.Id, err = res.LastInsertId(); err != nil {
		return err
	}
	return nil
}

func (br *bookRepository) GetById(id int) (*models.Book, error) {
	b := &models.Book{}

	err := br.db.QueryRow(
		"SELECT b.id, b.name, b.released, b.coast, b.pages, b.poster, b.author_id, a.last_name + ' ' + a.first_name,  b.genre_id, g.name FROM book b INNER JOIN author a ON a.id = b.author_id INNER JOIN genre g on b.genre_id = g.id WHERE b.id = ?", id,
	).Scan(&b.Id, &b.Name, &b.Release, &b.Coast, &b.Pages, &b.PosterURL, &b.AuthorId, &b.AuthorName, &b.GenreId, &b.GenreName)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func (br *bookRepository) Update(b *models.Book) error {
	_, err := br.db.Exec(
		"UPDATE book SET name = ?, poster = ?, coast = ?, pages = ?, released = ?, author_id = ?, genre_id = ? WHERE id = ?",
		b.Name, b.PosterURL, b.Coast, b.Pages, b.Release, b.AuthorId, b.GenreId, b.Id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (br *bookRepository) Delete(id int, idAuthor int) error {
	if _, err := br.db.Exec("DELETE FROM book WHERE id = ? AND author_id = ?", id, idAuthor); err != nil {
		return err
	}
	return nil
}

func (br *bookRepository) Count() (int, error) {
	var count int
	if err := br.db.QueryRow("SELECT COUNT(id) FROM book").Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func (br *bookRepository) GetPerPage(perPage int, page int) ([]models.Book, error) {
	start := (page - 1) * perPage
	rows, err := br.db.Query(
		"SELECT b.id, b.name, b.released, b.coast, b.pages, b.poster, b.author_id, a.last_name + ' ' + a.first_name,  b.genre_id, g.name FROM book b INNER JOIN author a ON a.id = b.author_id INNER JOIN genre g on b.genre_id = g.id ORDER BY b.id DESC LIMIT ?, ?",
		start, perPage,
	)
	if err != nil {
		return nil, err
	}

	var books []models.Book
	for rows.Next() {
		var b models.Book
		err := rows.Scan(&b.Id, &b.Name, &b.Release, &b.Coast, &b.Pages, &b.PosterURL, &b.AuthorId, &b.AuthorName, &b.GenreId, &b.GenreName)
		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

func (br *bookRepository) GetByGenre(idGenre, perPage, page int) ([]models.Book, error) {
	start := (page - 1) * perPage
	rows, err := br.db.Query(
		"SELECT b.id, b.name, b.released, b.coast, b.pages, b.poster, b.author_id, a.last_name + ' ' + a.first_name,  b.genre_id, g.name FROM book b INNER JOIN author a ON a.id = b.author_id INNER JOIN genre g on b.genre_id = g.id WHERE genre_id = ? ORDER BY b.id DESC LIMIT ?, ?",
		idGenre, start, perPage,
	)
	if err != nil {
		return nil, err
	}

	var books []models.Book
	for rows.Next() {
		var b models.Book
		err := rows.Scan(&b.Id, &b.Name, &b.Release, &b.Coast, &b.Pages, &b.PosterURL, &b.AuthorId, &b.AuthorName, &b.GenreId, &b.GenreName)
		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

func (br *bookRepository) GetByAuthor(idAuthor, perPage, page int) ([]models.Book, error) {
	start := (page - 1) * perPage
	rows, err := br.db.Query(
		"SELECT b.id, b.name, b.released, b.coast, b.pages, b.poster, b.author_id, a.last_name + ' ' + a.first_name,  b.genre_id, g.name FROM book b INNER JOIN author a ON a.id = b.author_id INNER JOIN genre g on b.genre_id = g.id WHERE author_id = ? ORDER BY b.id DESC LIMIT ?, ?",
		idAuthor, start, perPage,
	)
	if err != nil {
		return nil, err
	}

	var books []models.Book
	for rows.Next() {
		var b models.Book
		err := rows.Scan(&b.Id, &b.Name, &b.Release, &b.Coast, &b.Pages, &b.PosterURL, &b.AuthorId, &b.AuthorName, &b.GenreId, &b.GenreName)
		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

func (br *bookRepository) Search(value string) ([]models.Book, error) {
	value = "%" + value + "%"
	rows, err := br.db.Query(
		`SELECT b.id, b.name, b.released, b.coast, b.pages, b.poster, b.author_id, 
    	a.last_name + ' ' + a.first_name,  
    	b.genre_id, g.name 
		FROM book b INNER JOIN author a ON a.id = b.author_id INNER JOIN genre g on b.genre_id = g.id 
		WHERE b.name LIKE ? OR g.name LIKE ? OR a.first_name LIKE ? OR a.first_name = ?`,
		value, value, value, value,
	)
	if err != nil {
		return nil, err
	}

	var books []models.Book
	for rows.Next() {
		var b models.Book
		if err := rows.Scan(
			&b.Id, &b.Name, &b.Release, &b.Coast, &b.Pages, &b.PosterURL, &b.AuthorId, &b.AuthorName, &b.GenreId, &b.GenreName,
		); err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}
