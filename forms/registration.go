package forms

import (
	"gosharp/models"
)

type LoginForm struct {
	BaseForm
	UserName string `form:"username" json:"username" valid:"Required;MaxSize(50)"`
	Password string `form:"password" json:"password" valid:"Required"`

	User *models.AuthUser
}
