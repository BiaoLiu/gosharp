package models

import (
	"time"
)

type Article struct {
	Id        int `gorm:"primary_key;auto_increment"`
	Title     string
	ImageUrl  string
	Content   string
	Sort      int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
