package serializers

import (
	"gosharp/models"
	"time"
)

type UserSerializer struct {
	User *models.AuthUser
}

type UserResponse struct {
	Id          int       `json:"id"`
	Username    string    `json:"username"`
	Mobile      string    `json:"mobile"`
	Email       string    `json:"email"`
	CreatedTime time.Time `json:"created_time"`
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
