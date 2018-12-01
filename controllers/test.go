package controllers

import (
	"github.com/gin-gonic/gin"
)

func TestKong(c *gin.Context) {
	APIResponse(c, true, nil, "Kong请求成功")
}
