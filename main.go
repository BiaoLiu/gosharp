package main

import (
	"github.com/gin-gonic/gin"
	"gosharp/config"
	_ "gosharp/docs"
	"gosharp/middlewares"
	"gosharp/routers"
	"gosharp/startup"
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
	//初始化
	server.Init("config", "logs")

	//注册中间件
	middlewares.Register(engine)
	//注册路由
	routers.Register(engine)
	//engine.Static("/static", "./static")
	//engine.StaticFile("/protocol.html", "./views/protocol.html")
	//engine.LoadHTMLGlob("views/*")
	engine.Run(":" + config.Viper.GetString("PORT"))
}
