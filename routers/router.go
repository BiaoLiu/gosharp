package routers

import (
	"github.com/gin-gonic/gin"
	"gosharp/controllers"
	rescode "gosharp/utils/def"
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
	router.POST("/login", controllers.Login)

	router.POST("/articles", controllers.ArticleCreate)
	router.PUT("/articles/:id", controllers.ArticleUpdate)
	router.GET("/articles", controllers.ArticleList)
	router.DELETE("/articles", controllers.ArticleDelete)
	router.GET("/articles/:id", func(c *gin.Context) {
		path1 := c.Param("id")
		if path1 == "all" {
			controllers.ArticleAllList(c)
		} else {
			controllers.ArticleRetrieve(c)
		}
	})
}
