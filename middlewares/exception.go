package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gosharp/utils/log"
	"gosharp/utils/rescode"
	"net/http"
)

func ExceptionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				log.Logger.Error(fmt.Sprintf("request url: %s \r %s", c.Request.URL.Path, r))
				c.JSON(http.StatusInternalServerError, gin.H{"rescode": rescode.Error, "data": nil, "msg": r})
				c.Abort()
			}
		}()
		c.Next()
	}
}
