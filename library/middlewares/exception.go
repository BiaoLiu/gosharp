package middlewares

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gosharp/library/def"
	"gosharp/library/log"
	"net/http"
)

func ExceptionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				//登录认证异常
				if result, ok := r.(*ExceptionResult); ok {
					c.JSON(result.HttpStatus, result.Data)
					c.Abort()
				} else {
					var err error
					if err, ok = r.(error); !ok {
						err = errors.New(fmt.Sprintf("%s", r))
					}
					//日志记录
					log.Logger.Error(fmt.Sprintf("request url: %s \r %s", c.Request.URL.RequestURI, r))

					c.JSON(http.StatusInternalServerError, gin.H{"rescode": rescode.Error, "data": nil, "msg": err.Error()})
					c.Abort()
				}
			}
		}()
		c.Next()
	}
}
