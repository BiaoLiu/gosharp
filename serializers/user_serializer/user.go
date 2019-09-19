package user_serializer

import (
	"gosharp/models"
	"time"
)

type UserResponse struct {
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
	Model *models.AuthUser
}

func (s *UserSerializer) SingleResponse() UserResponse {
	return UserResponse{
		Id:          s.Model.ID,
		Username:    s.Model.Username,
		Mobile:      s.Model.Mobile,
		Email:       s.Model.Email.String,
		CreatedTime: s.Model.DateJoined,
	}
}
