package user_form

import (
	"gosharp/models"
	"gosharp/utils/app"
)

type LoginForm struct {
	app.BaseForm

	// 用户名
	// required: true
	UserName string `form:"username" json:"username" valid:"Required;MaxSize(50)"`
	// 密码
	// required: true
	Password string `form:"password" json:"password" valid:"Required"`
	// swagger:ignore
	User *models.AuthUser
}
