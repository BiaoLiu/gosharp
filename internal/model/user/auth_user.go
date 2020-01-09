package user

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

type LoginReq struct {
	// 用户名
	// required: true
	UserName string `form:"username" json:"username" valid:"Required;MaxSize(50)"`
	// 密码
	// required: true
	Password string `form:"password" json:"password" valid:"Required"`
	// swagger:ignore
	User *AuthUser
}

type UserResp struct {
	Id int `json:"id"`
	//用户名
	Username string `json:"username"`
	//手机号码
	Mobile string `json:"mobile"`
	//邮箱
	Email string `json:"email"`
	//创建时间
	CreatedTime time.Time `json:"created_time"`
}

type UserSerializer struct {
	Model *AuthUser
}

func (s *UserSerializer) SingleResponse() UserResp {
	return UserResp{
		Id:          s.Model.ID,
		Username:    s.Model.Username,
		Mobile:      s.Model.Mobile,
		Email:       s.Model.Email.String,
		CreatedTime: s.Model.DateJoined,
	}
}
