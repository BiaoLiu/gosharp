package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Accept, Accept-Encoding, Authorization, Content-Type, DNT, Origin, User-Agent, X-CSFRTOKEN, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS, DELETE")
		c.Header("Access-Control-Expose-Headers", "Authorization,X-Content-Range,X-Content-Total")
		//c.Set("content-type", "application/json")

		if c.Request.Method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		c.Next()
	}
}
