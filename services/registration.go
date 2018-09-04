package services

import (
	"github.com/astaxie/beego/validation"
	"gosharp/db"
	"gosharp/forms"
)

func Login(form *forms.LoginForm) []*validation.Error {
	v := validation.Validation{}

	db.Gorm.Where("(mobile=? or username=?) and is_deleted=0", form.UserName, form.UserName).Find(form.User)
	if form.User == nil {
		v.SetError("username", "用户名或密码错误")
		return v.Errors
	}

	if form.User.IsActive == 0 {
		v.SetError("username", "用户已被禁用，无法登录")
	}
	return v.Errors
}
