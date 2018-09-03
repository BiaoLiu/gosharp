package services

import (
	"github.com/astaxie/beego/validation"
	"gosharp/db"
	"gosharp/forms"
	"gosharp/utils/log"
)

func Login(form *forms.LoginForm) []*validation.Error {
	v := validation.Validation{}

	if has, err := db.Xorm.Where("(mobile=? or username=?) and is_deleted=0", form.UserName, form.UserName).Get(form.User); !has {
		if err != nil {
			log.Logger.Error("login error. ", err.Error())
		}
		v.SetError("username", "用户名或密码错误")
		return v.Errors
	}

	if form.User.IsActive == 0 {
		v.SetError("username", "用户已被禁用，无法登录")
	}

	return v.Errors
}
