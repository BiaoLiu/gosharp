package main

import (
	"github.com/gin-gonic/gin"
	"gosharp/config"
	"gosharp/db"
	_ "gosharp/docs"
	"gosharp/middlewares"
	"gosharp/routers"
	"gosharp/startup"
	"gosharp/utils/log"
)

// @title gosharp API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8200
// @BasePath /

// @securityDefinitions.apikey Token
// @in header
// @name Authorization
func main() {
	engine := gin.Default()
	//配置文件初始化
	config.Init("config")
	//日志初始化
	log.Init("logs")
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
