package user_service

import (
	"errors"
	"gosharp/forms/user_form"
	"gosharp/library/db"
	"gosharp/models"
)

func Login(form *user_form.LoginForm) error {
	user := new(models.AuthUser)
	if db.Gorm.Where("username=?", form.UserName).Find(user).RecordNotFound() {
		return errors.New("用户不存在")
	}

	if form.Password != user.Password { //具体判断逻辑，请自行替换
		return errors.New("用户名或密码错误")
	}
	if user.IsActive == 0 {
		return errors.New("用户已被禁用，无法登录")
	}
	return nil
}
