package http

import (
	"github.com/gin-gonic/gin"
	"gosharp/internal/model/user"
	"gosharp/library/app"
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
	arg := new(user.LoginReq)
	if err := app.BindAndValidate(c, arg); err != nil {
		app.APIResponse(c, false, nil, err.Error())
		return
	}

	//登录
	if err := svc.Login(arg); err != nil {
		app.APIResponse(c, false, nil, err.Error())
		return
	}

	//todo 保存登录信息

	serializer := user.UserSerializer{Model: arg.User}

	app.APIResponse(c, true, serializer.SingleResponse(), "")
}
