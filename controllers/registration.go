package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"gosharp/forms"
	"gosharp/serializers"
	"gosharp/services"
)

// swagger:route POST /login auth wxLoginFormWrap
// 登录
//
// 登录
// responses:
//   200: userResponseWrap
func PostLogin(c *gin.Context) {
	form := forms.LoginForm{}
	if err := bindAndValidateForm(c, &form); err != nil {
		APIResponse(c, false, nil, err.Error())
		return
	}
	token := c.Request.Header["Authorization"]
	fmt.Println(token)

	//登录
	if err := services.Login(&form); err != nil {
		APIResponse(c, false, nil, err.Error())
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

// swagger:route POST /test/swag-route test swagRouteFormWrap
// 测试route
//
// 测试route
// responses:
//   200: swagRouteResponseWrap
func TestSwagRoute(c *gin.Context) {
	var form forms.SwagRouteForm
	if err := bindAndValidateForm(c, &form); err != nil {
		APIResponse(c, false, nil, err.Error())
		return
	}
	result := serializers.SwagRouteResponse{Url: form.Url}

	APIResponse(c, true, result, "")
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
