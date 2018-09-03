package serializers

import (
	"gosharp/models"
	"time"
)

type UserSerializer struct {
	User *models.AuthUser
}

type UserResponse struct {
	Id          int       `json:"id" example:"1"`
	Username    string    `json:"username" example:"用户名"`
	Mobile      string    `json:"mobile" example:"手机号码"`
	Email       string    `json:"email" example:"邮箱"`
	CreatedTime time.Time `json:"created_time" example:"创建时间"`
}

func (s *UserSerializer) Response() UserResponse {
	return UserResponse{
		Id:          s.User.Id,
		Username:    s.User.Username,
		Mobile:      s.User.Mobile,
		Email:       s.User.Email,
		CreatedTime: s.User.DateJoined,
	}
}
