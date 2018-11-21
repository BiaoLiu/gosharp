package routers

import (
	"github.com/gin-gonic/gin"
	"gosharp/controllers"
	"gosharp/utils/rescode"
	"net/http"
)

func errorWrapper(handler gin.HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"rescode": rescode.Error, "data": nil, "msg": r})
			}
		}()
		handler(c)
	}
}

func Register(router *gin.Engine) {
	router.POST("/login", controllers.PostLogin)

	router.GET("/test", controllers.Test)
	router.GET("/test/swag-operation", controllers.TestSwagOperation)
	router.POST("/test/swag-route", controllers.TestSwagRoute)

	router.GET("/testconsul", controllers.TestConsul)

}
