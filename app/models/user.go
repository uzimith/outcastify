package models

import "time"

type User struct {
	Id        int64
	Secrets   []Secret `gorm:"many2many:user_secret;"`
	Room      string   `sql:"size:16;not null"`
	Name      string   `sql:"size:255"`
	Join      bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
