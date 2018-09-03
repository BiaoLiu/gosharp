package main

import (
	"github.com/gin-gonic/gin"
	"gosharp/config"
	"gosharp/db"
	"gosharp/middlewares"
	"gosharp/routers"
	"gosharp/startup"
	"gosharp/utils/log"
)

func main() {
	engine := gin.Default()
	//配置文件初始化

	config.Init()
	//日志初始化
	log.Init()
	//数据库初始化
	db.Init()
	defer db.Close()
	//gin配置初始化
	server.Init(engine)

	//注册中间件
	middlewares.Register(engine)
	//注册路由
	routers.Register(engine)
	//engine.Static("/static", "./static")
	//engine.StaticFile("/protocol.html", "./views/protocol.html")
	//engine.LoadHTMLGlob("views/*")
	engine.Run(":" + config.Viper.GetString("PORT"))
}
