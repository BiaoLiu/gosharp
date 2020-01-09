package service

import (
	"errors"
	"gosharp/internal/model/user"
)

func (s *Service) Login(arg *user.LoginReq) error {
	user := new(user.AuthUser)
	if s.dao.DB().Where("username=?", arg.UserName).Find(user).RecordNotFound() {
		return errors.New("用户不存在")
	}

	if arg.Password != user.Password { //具体判断逻辑，请自行替换
		return errors.New("用户名或密码错误")
	}
	if user.IsActive == 0 {
		return errors.New("用户已被禁用，无法登录")
	}
	return nil
}
