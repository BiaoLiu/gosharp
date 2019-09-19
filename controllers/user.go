package controllers

import (
	"github.com/gin-gonic/gin"
	"gosharp/forms/user_form"
	"gosharp/library/app"
	"gosharp/serializers/user_serializer"
	"gosharp/services/user_service"
)

// swagger:operation POST /login User 用户
// ---
// summary: 登录
// description: 登录
// parameters:
// - name: Body
//   in: body
//   schema:
//       "$ref": "#/definitions/LoginForm"
// responses:
//     "200":
//       "$ref": "#/responses/userResponseWrap"
func Login(c *gin.Context) {
	form := user_form.LoginForm{}
	if err := app.BindAndValidate(c, &form); err != nil {
		app.APIResponse(c, false, nil, err.Error())
		return
	}

	//登录
	if err := user_service.Login(&form); err != nil {
		app.APIResponse(c, false, nil, err.Error())
		return
	}

	//todo 保存登录信息

	serializer := user_serializer.UserSerializer{Model: form.User}

	app.APIResponse(c, true, serializer.SingleResponse(), "")
}
