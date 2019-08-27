// Package classification gosharp API.
//
// gosharp API
//
//      Host: localhost:9000 | api.xxx.com
//      Version: 1.0.0
//
//      Security:
//      - Token:
//
//      SecurityDefinitions:
//      Token:
//           type: apiKey
//           name: Authorization
//           in: header
// swagger:meta
package main

import (
	"github.com/gin-gonic/gin"
	"gosharp/library/config"
	"gosharp/middlewares"
	"gosharp/routers"
	"gosharp/startup"
)

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
