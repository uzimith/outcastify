package models

import "time"

type User struct {
	Id        int64
	Name      string `sql:"size:255"`
	Room      string `sql:"size:16;not null"`
	Join      bool
	Token     string `sql:"size:16;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
