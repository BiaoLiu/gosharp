package unit_test

import (
	"github.com/gin-gonic/gin"
	"gosharp/library"
	"gosharp/routers"
	"gosharp/startup"
)

func NewEngine() *gin.Engine {
	engine := gin.New()
	//初始化
	server.Init("../config", "../logs")

	//注册中间件
	library.Register(engine)
	//注册路由
	routers.Register(engine)

	return engine
}
