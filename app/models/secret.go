package models

import "time"

type Secret struct {
	Id        int64
	Users     []User `gorm:"many2many:user_secret;"`
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
