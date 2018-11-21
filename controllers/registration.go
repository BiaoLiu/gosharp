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
	token := c.Request.Header["Authorization"]
	fmt.Println(token)
	//登录
	err := services.Login(&form)
	if err != nil {
		CheckError(c, err)
		return
	}
	serializer := serializers.UserSerializer{User: form.User}

	APIResponse(c, true, serializer.Response(), "")
}

// swagger:operation GET /test/swag-operation test 测试swagger-operation
// ---
// summary: 测试swagger
// description: 测试swagger operation
// parameters:
// - name: name
//   in: query
//   description: 名称
//   type: string
//   required: true
// responses:
//   "200":
//     description: 测试结果
func TestSwagOperation(c *gin.Context) {
	name := c.Query("name")
	APIResponse(c, true, gin.H{"name": name}, "")
}

// swagger:route POST /test/swag-route test wxLoginFormWrap
// 测试route
//
// 测试route
// responses:
//   200: userResponseWrap
func TestSwagRoute(c *gin.Context) {
	var form forms.LoginForm
	if !bindAndValidateForm(c, &form) {
		return
	}

	APIResponse(c, true, gin.H{"username": form.UserName, "password": form.Password}, "")
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
