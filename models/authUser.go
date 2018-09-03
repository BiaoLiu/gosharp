package models

import (
	"time"
)

type AuthUser struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	Username   string    `xorm:"not null VARCHAR(150)"`
	Password   string    `xorm:"not null VARCHAR(128)"`
	Mobile     string    `xorm:"not null VARCHAR(20)"`
	Realname   string    `xorm:"VARCHAR(20)"`
	Gender     int       `xorm:"TINYINT(4)"`
	Email      string    `xorm:"VARCHAR(254)"`
	IsActive   int       `xorm:"not null TINYINT(4)"`
	DateJoined time.Time `xorm:"not null DATETIME"`
	LastLogin  time.Time `xorm:"DATETIME"`
	IsDeleted  int       `xorm:"not null default 0 TINYINT(4)"`
}
