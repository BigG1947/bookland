package store

import (
	"bookland/internal/db"
	"bookland/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBookRepository_Add(t *testing.T) {
	testCases := []struct {
		name  string
		book  models.Book
		valid bool
	}{
		{
			name: "valid book",
			book: models.Book{
				Name:      "Test book",
				Release:   time.Now().UTC(),
				Coast:     250,
				Pages:     300,
				PosterURL: "img.png",
				AuthorId:  1,
				GenreId:   1,
			},
			valid: true,
		},
		{
			name: "invalid author",
			book: models.Book{
				Name:      "Test book",
				Release:   time.Now().UTC(),
				Coast:     250,
				Pages:     300,
				PosterURL: "img.png",
				AuthorId:  200,
				GenreId:   1,
			},
			valid: false,
		},
		{
			name: "invalid genre",
			book: models.Book{
				Name:      "Test book",
				Release:   time.Now().UTC(),
				Coast:     250,
				Pages:     300,
				PosterURL: "img.png",
				AuthorId:  1,
				GenreId:   900,
			},
			valid: false,
		},
		{
			name:  "invalid book",
			book:  models.Book{},
			valid: false,
		},
	}

	conn := db.NewTestSQLiteDB(t)
	defer conn.Close()
	defer db.DropTestSQLiteDB(t)
	br := newBookRepository(conn)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := br.Add(&tc.book)

			if tc.valid {
				assert.NoError(t, err)
				assert.NotZero(t, tc.book.Id)
			} else {
				assert.Error(t, err)
			}
		})

	}
}

func TestBookRepository_GetById(t *testing.T) {
	testCases := []struct {
		name   string
		idBook int
		valid  bool
	}{
		{
			name:   "valid id",
			idBook: 1,
			valid:  true,
		},
		{
			name:   "invalid id",
			idBook: 99,
			valid:  false,
		},
	}

	conn := db.NewTestSQLiteDB(t)
	defer conn.Close()
	defer db.DropTestSQLiteDB(t)
	br := newBookRepository(conn)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			book, err := br.GetById(tc.idBook)
			if tc.valid {
				assert.NoError(t, err)
				assert.NotNil(t, book)
			} else {
				assert.Error(t, err)
				assert.Nil(t, book)
			}
		})
	}
}

func TestBookRepository_Update(t *testing.T) {
	validBook := &models.Book{
		Id:        1,
		Name:      "updating book",
		Release:   time.Date(2010, 10, 10, 0, 0, 0, 0, time.Local),
		Coast:     888,
		Pages:     999,
		PosterURL: "img.png2",
		AuthorId:  1,
		GenreId:   1,
	}

	incorrectBook := &models.Book{
		Id:        1,
		Name:      "updating book",
		Release:   time.Time{},
		Coast:     888,
		Pages:     999,
		PosterURL: "",
		AuthorId:  256,
		GenreId:   234,
	}

	testCases := []struct {
		name  string
		book  *models.Book
		valid bool
	}{
		{
			name:  "valid book",
			book:  validBook,
			valid: true,
		},
		{
			name:  "incorrect book",
			book:  incorrectBook,
			valid: false,
		},
	}

	conn := db.NewTestSQLiteDB(t)
	defer conn.Close()
	defer db.DropTestSQLiteDB(t)
	br := newBookRepository(conn)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := br.Update(tc.book)
			if tc.valid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}

	return
}

func TestBookRepository_Delete(t *testing.T) {
	testCase := []struct {
		name     string
		idBook   int
		idAuthor int
		deleted  bool
	}{
		{
			name:     "valid id. id_author",
			idBook:   1,
			idAuthor: 1,
			deleted:  true,
		},
		{
			name:     "invalid id",
			idBook:   99,
			idAuthor: 1,
			deleted:  false,
		},
		{
			name:     "invalid id author",
			idBook:   1,
			idAuthor: 99,
			deleted:  false,
		},
		{
			name:     "invalid id, id_author",
			idBook:   99,
			idAuthor: 99,
			deleted:  false,
		},
	}

	conn := db.NewTestSQLiteDB(t)
	defer conn.Close()
	defer db.DropTestSQLiteDB(t)
	br := newBookRepository(conn)

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			countBefore, err := br.Count()
			assert.NoError(t, err)
			err = br.Delete(tc.idBook, tc.idAuthor)
			countAfter, err := br.Count()
			assert.NoError(t, err)
			if tc.deleted {
				assert.NoError(t, err)
				b, err := br.GetById(tc.idBook)
				assert.Nil(t, b)
				assert.Error(t, err)
				assert.Equal(t, countBefore-1, countAfter)
			} else {
				assert.Equal(t, countBefore, countAfter)
			}
		})
	}
}

func TestBookRepository_Count(t *testing.T) {
	conn := db.NewTestSQLiteDB(t)
	defer conn.Close()
	defer db.DropTestSQLiteDB(t)
	br := newBookRepository(conn)

	count, err := br.Count()

	assert.NoError(t, err)
	assert.NotZero(t, count)
}

func TestBookRepository_GetPerPage(t *testing.T) {
	testCases := []struct {
		name      string
		perPage   int
		page      int
		countBook int
	}{
		{
			name:      "first page",
			perPage:   10,
			page:      1,
			countBook: 10,
		},
		{
			name:      "second page",
			perPage:   10,
			page:      2,
			countBook: 6,
		},
		{
			name:      "empty page",
			perPage:   10,
			page:      3,
			countBook: 0,
		},
	}

	conn := db.NewTestSQLiteDB(t)
	defer conn.Close()
	defer db.DropTestSQLiteDB(t)
	br := newBookRepository(conn)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			books, err := br.GetPerPage(tc.perPage, tc.page)
			assert.NoError(t, err)
			assert.Equal(t, tc.countBook, len(books))
		})
	}
}

func TestBookRepository_GetByGenre(t *testing.T) {
	testCases := []struct {
		name      string
		idGenre   int
		perPage   int
		page      int
		countBook int
	}{
		{
			name:      "first page",
			idGenre:   2,
			perPage:   10,
			page:      1,
			countBook: 6,
		},
		{
			name:      "empty page",
			idGenre:   2,
			perPage:   10,
			page:      2,
			countBook: 0,
		},
	}

	conn := db.NewTestSQLiteDB(t)
	defer conn.Close()
	defer db.DropTestSQLiteDB(t)
	br := newBookRepository(conn)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			books, err := br.GetByGenre(tc.idGenre, tc.perPage, tc.page)
			assert.NoError(t, err)
			assert.Equal(t, tc.countBook, len(books))
		})
	}
}

func TestBookRepository_GetByAuthor(t *testing.T) {
	testCases := []struct {
		name      string
		idAuthor  int
		perPage   int
		page      int
		countBook int
	}{
		{
			name:      "first page",
			idAuthor:  2,
			perPage:   10,
			page:      1,
			countBook: 6,
		},
		{
			name:      "empty page",
			idAuthor:  2,
			perPage:   10,
			page:      2,
			countBook: 0,
		},
	}

	conn := db.NewTestSQLiteDB(t)
	defer conn.Close()
	defer db.DropTestSQLiteDB(t)
	br := newBookRepository(conn)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			books, err := br.GetByAuthor(tc.idAuthor, tc.perPage, tc.page)
			assert.NoError(t, err)
			assert.Equal(t, tc.countBook, len(books))
		})
	}
}

func TestBookRepository_SearchByName(t *testing.T) {
	testCases := []struct {
		name      string
		searchVal string
		found     bool
	}{
		{
			name:      "empty value",
			searchVal: "",
			found:     true,
		},
		{
			name:      "valid value",
			searchVal: "book",
			found:     true,
		},
		{
			name:      "not found value",
			searchVal: "not found book",
			found:     false,
		},
		{
			name:      "search by author name",
			searchVal: "Harry",
			found:     true,
		},
		{
			name:      "search by genre name",
			searchVal: "test_genre 2",
			found:     true,
		},
	}

	conn := db.NewTestSQLiteDB(t)
	defer conn.Close()
	defer db.DropTestSQLiteDB(t)
	br := newBookRepository(conn)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			books, err := br.Search(tc.searchVal)
			assert.NoError(t, err)

			if tc.found {
				assert.NotEmpty(t, books)
			} else {
				assert.Empty(t, books)
			}
		})
	}
}
