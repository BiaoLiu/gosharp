package controllers

import (
	"github.com/gin-gonic/gin"
	"gosharp/utils/app"
)

func TestKong(c *gin.Context) {
	app.APIResponse(c, true, nil, "Kong请求成功")
}
