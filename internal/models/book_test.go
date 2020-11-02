package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBook_IsValid(t *testing.T) {
	testCases := []struct {
		name  string
		book  *Book
		valid bool
	}{
		{
			name: "valid book",
			book: &Book{
				Id:         0,
				Name:       "Book",
				Release:    time.Date(2010, 10, 10, 0, 0, 0, 0, time.Local),
				Coast:      250,
				Pages:      200,
				PosterURL:  "",
				AuthorId:   1,
				AuthorName: "",
				GenreId:    1,
				GenreName:  "",
			},
			valid: true,
		},
		{
			name: "invalid book name",
			book: &Book{
				Id:         0,
				Name:       "",
				Release:    time.Date(2010, 10, 10, 0, 0, 0, 0, time.Local),
				Coast:      250,
				Pages:      200,
				PosterURL:  "",
				AuthorId:   1,
				AuthorName: "",
				GenreId:    1,
				GenreName:  "",
			},
			valid: false,
		},
		{
			name: "invalid book release date",
			book: &Book{
				Id:         0,
				Name:       "Book",
				Release:    time.Now().Add(time.Hour * 48),
				Coast:      250,
				Pages:      200,
				PosterURL:  "",
				AuthorId:   1,
				AuthorName: "",
				GenreId:    1,
				GenreName:  "",
			},
			valid: false,
		},
		{
			name: "invalid book coast",
			book: &Book{
				Id:         0,
				Name:       "Book",
				Release:    time.Date(2010, 10, 10, 0, 0, 0, 0, time.Local),
				Coast:      0,
				Pages:      200,
				PosterURL:  "",
				AuthorId:   1,
				AuthorName: "",
				GenreId:    1,
				GenreName:  "",
			},
			valid: false,
		},
		{
			name: "invalid book count pages",
			book: &Book{
				Id:         0,
				Name:       "Book",
				Release:    time.Date(2010, 10, 10, 0, 0, 0, 0, time.Local),
				Coast:      250,
				Pages:      0,
				PosterURL:  "",
				AuthorId:   1,
				AuthorName: "",
				GenreId:    1,
				GenreName:  "",
			},
			valid: false,
		},
		{
			name: "invalid book author",
			book: &Book{
				Id:         0,
				Name:       "Book",
				Release:    time.Date(2010, 10, 10, 0, 0, 0, 0, time.Local),
				Coast:      250,
				Pages:      200,
				PosterURL:  "",
				AuthorId:   0,
				AuthorName: "",
				GenreId:    1,
				GenreName:  "",
			},
			valid: false,
		},
		{
			name: "invalid book genre",
			book: &Book{
				Id:         0,
				Name:       "Book",
				Release:    time.Date(2010, 10, 10, 0, 0, 0, 0, time.Local),
				Coast:      250,
				Pages:      200,
				PosterURL:  "",
				AuthorId:   1,
				AuthorName: "",
				GenreId:    0,
				GenreName:  "",
			},
			valid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ok, message := tc.book.IsValid()
			if tc.valid {
				assert.Empty(t, message)
				assert.True(t, ok)
			} else {
				assert.NotEmpty(t, message)
				assert.False(t, ok)
			}
		})
	}

}
