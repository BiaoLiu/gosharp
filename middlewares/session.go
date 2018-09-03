package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"gosharp/utils/auth"
)

func SessionMiddleware(store sessions.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		//请求上下文设置session store
		c.Set(auth.SessionStore, store)
		c.Next()
	}
}
