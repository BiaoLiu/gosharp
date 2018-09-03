package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := "*"
		fmt.Println(c.Request.Header["Origin"])
		if origins := c.Request.Header["Origin"]; len(origins) > 0 {
			origin = origins[0]
		}
		c.Header("Access-Control-Allow-Origin", origin)
		if origin != "*" {
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		c.Header("Access-Control-Allow-Headers", "Accept, Accept-Encoding, Authorization, Content-Type, DNT, Origin, User-Agent, X-CSFRTOKEN, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS, DELETE")
		c.Header("Access-Control-Expose-Headers", "Authorization")
		//c.Set("content-type", "application/json")

		if c.Request.Method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		c.Next()
	}
}
