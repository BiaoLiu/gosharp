package unit_test

import (
	"github.com/gin-gonic/gin"
	"gosharp/db"
	"gosharp/middlewares"
	"gosharp/routers"
	"gosharp/startup"
)

func NewEngine() *gin.Engine {
	engine := gin.New()
	//配置初始化
	server.Init("../config", "../logs")
	//数据库初始化
	db.Init()
	defer db.Close()

	//注册中间件
	middlewares.Register(engine)
	//注册路由
	routers.Register(engine)

	return engine
}
