package forms

import (
	"gosharp/models"
)

// swagger:parameters wxLoginFormWrap
type wxLoginFormWrap struct {
	// in:body
	Body LoginForm
}

type LoginForm struct {
	BaseForm
	// 用户名
	UserName string `form:"username" json:"username" valid:"Required;MaxSize(50)"`
	// 密码
	Password string `form:"password" json:"password" valid:"Required"`
	// swagger:ignore
	User *models.AuthUser
}
