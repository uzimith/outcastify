package models

import "time"

type User struct {
	Id        int64
	Name      string `sql:"size:255"`
	Join      bool
	Room      string `sql:"size:16;not null"`
	ShareId   int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
