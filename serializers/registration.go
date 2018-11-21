package serializers

import (
	"gosharp/models"
	"time"
)

// 用户信息
// swagger:response userResponseWrap
type userResponseWrap struct {
	// in:body
	Body UserResponse
}

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
	User *models.AuthUser
}

func (s *UserSerializer) Response() UserResponse {
	return UserResponse{
		Id:          s.User.ID,
		Username:    s.User.Username,
		Mobile:      s.User.Mobile,
		Email:       s.User.Email.String,
		CreatedTime: s.User.DateJoined,
	}
}
