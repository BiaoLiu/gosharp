package http

import (
	"github.com/gin-gonic/gin"
	"gosharp/library/app"
)

func TestKong(c *gin.Context) {
	app.APIResponse(c, true, nil, "Kong请求成功")
}
