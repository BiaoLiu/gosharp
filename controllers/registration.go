package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"gosharp/forms"
	"gosharp/serializers"
	"gosharp/services"
)

// @Tags 登录注册模块
// @Summary 登录
// @Produce  json
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @Success 200 {object} serializers.UserResponse
// @Security Token
// @Router /login [post]
func PostLogin(c *gin.Context) {
	form := forms.LoginForm{}
	if !bindAndValidateForm(c, &form) {
		return
	}
	token := c.Request.Header["Authorization"]
	fmt.Println(token)
	//登录
	err := services.Login(&form)
	if err != nil {
		setError(c, err)
		return
	}
	serializer := serializers.UserSerializer{User: form.User}

	APIResponse(c, true, serializer.Response(), "")
}

// @Tags 测试模块
// @Summary 测试
// @Produce  json
// @Success 200 {string}  success
// @Security Token
// @Router /test [get]
func Test(c *gin.Context) {
	value, _ := c.Cookie("robo2025")

	var userId int
	err := securecookie.DecodeMulti("robo", value, &userId, securecookie.CodecsFromPairs([]byte("secret"))...)
	if err != nil {
		fmt.Println(userId)
	}

	c.JSON(200, gin.H{"status": "success"})
}
