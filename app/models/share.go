package models

import "time"

type Share struct {
	Id        int64
	Users     []User
	Room      string
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
