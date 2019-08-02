package models

import (
	"database/sql"
	"time"
)

type AuthUser struct {
	ID         int            `gorm:"column:id"`
	Mobile     string         `gorm:"column:mobile"`
	Username   string         `gorm:"column:username"`
	Password   string         `gorm:"column:password"`
	DateJoined time.Time      `gorm:"column:date_joined"`
	Email      sql.NullString `gorm:"column:email"`
	Gender     sql.NullInt64  `gorm:"column:gender"`
	IsActive   int            `gorm:"column:is_active"`
	DeletedAt  *time.Time     `gorm:"column:deleted_at"`
	LastLogin  *time.Time     `gorm:"column:last_login"`
}
