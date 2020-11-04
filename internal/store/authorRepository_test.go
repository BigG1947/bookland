package store

import (
	"bookland/internal/db"
	"bookland/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAuthorRepository_Get(t *testing.T) {
	testCases := []struct {
		name  string
		id    int
		valid bool
	}{
		{
			name:  "valid id author",
			id:    1,
			valid: true,
		}, {
			name:  "invalid id author",
			id:    0,
			valid: false,
		},
	}

	conn := db.NewTestSQLiteDB(t)
	defer func() {
		if err := conn.Close(); err != nil {
			t.Fatal()
		}
	}()
	defer db.DropTestSQLiteDB(t)
	ar := newAuthorRepository(conn)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			author, err := ar.Get(tc.id)
			if tc.valid {
				assert.NotNil(t, author)
				assert.Nil(t, err)
			} else {
				assert.Nil(t, author)
				assert.NotNil(t, err)
			}
		})
	}
}

func TestAuthorRepository_Add(t *testing.T) {
	author := &models.Author{
		Id:        0,
		LastName:  "Potter",
		FirstName: "Harry",
		BirthDay:  time.Time{},
		Bio:       "Test Bio",
	}

	conn := db.NewTestSQLiteDB(t)
	defer func() {
		if err := conn.Close(); err != nil {
			t.Fatal()
		}
	}()
	defer db.DropTestSQLiteDB(t)
	ar := newAuthorRepository(conn)

	err := ar.Add(author)

	assert.NotZero(t, author.Id)
	assert.NoError(t, err)
}

func TestAuthorRepository_Update(t *testing.T) {
	updateAuthor := &models.Author{
		Id:        1,
		LastName:  "Test Author",
		FirstName: "Test Author",
		BirthDay:  time.Time{},
		Bio:       "Test Bio",
	}

	conn := db.NewTestSQLiteDB(t)
	defer func() {
		if err := conn.Close(); err != nil {
			t.Fatal()
		}
	}()
	defer db.DropTestSQLiteDB(t)
	ar := newAuthorRepository(conn)

	err := ar.Update(updateAuthor)
	assert.NoError(t, err)
	actualAuthor, err := ar.Get(int(updateAuthor.Id))
	assert.NoError(t, err)
	assert.Equal(t, updateAuthor, actualAuthor)
}

func TestAuthorRepository_Delete(t *testing.T) {
	conn := db.NewTestSQLiteDB(t)
	defer func() {
		if err := conn.Close(); err != nil {
			t.Fatal()
		}
	}()
	defer db.DropTestSQLiteDB(t)
	ar := newAuthorRepository(conn)

	err := ar.Delete(1)
	assert.NoError(t, err)
}

func TestAuthorRepository_Count(t *testing.T) {
	conn := db.NewTestSQLiteDB(t)
	defer func() {
		if err := conn.Close(); err != nil {
			t.Fatal()
		}
	}()
	defer db.DropTestSQLiteDB(t)
	ar := newAuthorRepository(conn)

	count, err := ar.Count()
	assert.NoError(t, err)
	assert.Equal(t, 2, count)
}

func TestAuthorRepository_SearchByName(t *testing.T) {
	valid := "arry"
	invalid := "invalid value for search"

	conn := db.NewTestSQLiteDB(t)
	defer func() {
		if err := conn.Close(); err != nil {
			t.Fatal()
		}
	}()
	defer db.DropTestSQLiteDB(t)
	ar := newAuthorRepository(conn)

	authors, err := ar.SearchByName(valid)
	assert.NotNil(t, authors)
	assert.NoError(t, err)

	authors, err = ar.SearchByName(invalid)
	assert.Nil(t, authors)
	assert.NoError(t, err)
}

func TestAuthorRepository_GetPerPage(t *testing.T) {
	conn := db.NewTestSQLiteDB(t)
	defer func() {
		if err := conn.Close(); err != nil {
			t.Fatal()
		}
	}()
	defer db.DropTestSQLiteDB(t)
	ar := newAuthorRepository(conn)

	authors, err := ar.GetPerPage(10, 1)
	assert.NotNil(t, authors)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(authors))

	authors, err = ar.GetPerPage(10, 2)
	assert.Nil(t, authors)
	assert.NoError(t, err)
	assert.Zero(t, len(authors))
}
