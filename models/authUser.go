package models

import (
	"database/sql"
	"time"
)

type AuthUser struct {
	ID         int            `gorm:"column:id"`
	Mobile     string         `gorm:"column:mobile"`
	Password   sql.NullString `gorm:"column:password"`
	Realname   sql.NullString `gorm:"column:realname"`
	UserType   int            `gorm:"column:user_type"`
	Username   string         `gorm:"column:username"`
	DateJoined time.Time      `gorm:"column:date_joined"`
	Email      sql.NullString `gorm:"column:email"`
	Gender     sql.NullInt64  `gorm:"column:gender"`
	IsActive   int            `gorm:"column:is_active"`
	IsDeleted  int            `gorm:"column:is_deleted"`
	IsSubuser  int            `gorm:"column:is_subuser"`
	LastLogin  time.Time      `gorm:"column:last_login"`
	MainUserID int            `gorm:"column:main_user_id"`
}
