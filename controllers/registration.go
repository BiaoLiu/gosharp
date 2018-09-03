package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"gosharp/forms"
	"gosharp/serializers"
	"gosharp/services"
)

func PostLogin(c *gin.Context) {
	form := forms.LoginForm{}
	if !bindAndValidateForm(c, &form) {
		return
	}
	//登录
	err := services.Login(&form)
	if err != nil {
		setError(c, err)
		return
	}
	serializer := serializers.UserSerializer{form.User}

	APIResponse(c, true, serializer.Response(), "")
}

func Test(c *gin.Context) {
	value, _ := c.Cookie("robo2025")

	var userId int
	err := securecookie.DecodeMulti("robo", value, &userId, securecookie.CodecsFromPairs([]byte("secret"))...)
	if err != nil {
		fmt.Println(userId)
	}

	c.JSON(200, gin.H{"status": "success"})
}
