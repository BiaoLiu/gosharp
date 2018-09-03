package server

import (
	"github.com/gin-gonic/gin"
	"gosharp/utils/validation"
)

func Init(engine *gin.Engine) {
	//设置beego validation 错误信息
	validation.SetValidationMessage()
}
