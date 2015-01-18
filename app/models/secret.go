package models

import "time"

type Secret struct {
	Id        int64
	Users     []User `gorm:"many2many:user_secret;"`
	Room      string `sql:"size:16;not null"`
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
