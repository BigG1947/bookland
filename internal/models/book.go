package models

import "time"

type Book struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	Release    time.Time `json:"release"`
	Coast      uint      `json:"coast"`
	Pages      uint      `json:"pages"`
	PosterURL  string    `json:"poster_url"`
	AuthorId   int64     `json:"author_id"`
	AuthorName string    `json:"author_name"`
	GenreId    int64     `json:"genre_id"`
	GenreName  string    `json:"genre_name"`
}

func (b *Book) IsValid() (bool, string) {
	if b.Name == "" {
		return false, "Book name is require field"
	}

	if !b.Release.Before(time.Now()) {
		return false, "Release date must be in past"
	}

	if b.Coast == 0 {
		return false, "Coast is require field"
	}

	if b.Pages == 0 {
		return false, "Pages is require field"
	}

	if b.AuthorId <= 0 {
		return false, "author_id is require field"
	}

	if b.GenreId <= 0 {
		return false, "genre_id is require field"
	}

	return true, ""
}
