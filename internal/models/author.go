package models

import "time"

type Author struct {
	Id        int64     `json:"id"`
	LastName  string    `json:"last_name"`
	FirstName string    `json:"first_name"`
	BirthDay  time.Time `json:"birth_day"`
	Bio       string    `json:"bio"`
}
