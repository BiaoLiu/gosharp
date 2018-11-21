package services

import (
	"errors"
	"gosharp/db"
	"gosharp/forms"
)

func Login(form *forms.LoginForm) error {
	db.Gorm.Where("(mobile=? or username=?) and is_deleted=0", form.UserName, form.UserName).Find(form.User)
	if form.User == nil {
		return errors.New("用户名或密码错误")
	}
	if form.User.IsActive == 0 {
		return errors.New("用户已被禁用，无法登录")
	}
	return nil
}
