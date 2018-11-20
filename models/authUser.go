package models

import (
	"database/sql"
	"time"
)

type AuthUser struct {
	ID         int            `gorm:"column:id"`
	Mobile     string         `gorm:"column:mobile"`
	Username   string         `gorm:"column:username"`
	Password   sql.NullString `gorm:"column:password"`
	DateJoined time.Time      `gorm:"column:date_joined"`
	Email      sql.NullString `gorm:"column:email"`
	Gender     sql.NullInt64  `gorm:"column:gender"`
	IsActive   int            `gorm:"column:is_active"`
	IsDeleted  int            `gorm:"column:is_deleted"`
	LastLogin  time.Time      `gorm:"column:last_login"`
}
