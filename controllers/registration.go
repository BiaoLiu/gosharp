package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"gosharp/forms"
	"gosharp/serializers"
	"gosharp/services"
	"gosharp/utils/app"
)

// swagger:route POST /login auth wxLoginFormWrap
// 登录
//
// 登录
// responses:
//   200: userResponseWrap
func PostLogin(c *gin.Context) {
	form := forms.LoginForm{}
	if err := app.BindAndValidate(c, &form); err != nil {
		app.APIResponse(c, false, nil, err.Error())
		return
	}
	token := c.Request.Header["Authorization"]
	fmt.Println(token)

	//登录
	if err := services.Login(&form); err != nil {
		app.APIResponse(c, false, nil, err.Error())
		return
	}

	serializer := serializers.UserSerializer{User: form.User}

	app.APIResponse(c, true, serializer.Response(), "")
}

func Test(c *gin.Context) {
	value, _ := c.Cookie("gosharp")

	var userId int
	err := securecookie.DecodeMulti("gosharp", value, &userId, securecookie.CodecsFromPairs([]byte("secret"))...)
	if err != nil {
		fmt.Println(userId)
	}

	c.JSON(200, gin.H{"status": "success"})
}
