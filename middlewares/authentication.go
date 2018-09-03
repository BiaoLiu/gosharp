package middlewares

import (
	"github.com/gin-gonic/gin"
	"gosharp/utils/auth"
	"gosharp/utils/rescode"
)

type ExceptionResult struct {
	HttpStatus int
	Data       map[string]interface{}
}

//普通用户授权
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, status, res, msg := auth.GetUser(c)
		if user == nil {
			panic(&ExceptionResult{HttpStatus: status, Data: gin.H{"rescode": res, "data": nil, "msg": msg}})
		}
		if user.UserType != 1 {
			panic(&ExceptionResult{HttpStatus: 403, Data: gin.H{"rescode": rescode.Access_Denied, "data": nil, "msg": "非普通用户，无权限访问"}})
		}
		//请求上下文设置user
		c.Set(auth.AuthUser, user)
		c.Next()
	}
}

//供应商授权
func SupplierAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, status, res, msg := auth.GetUser(c)
		if user == nil {
			panic(&ExceptionResult{HttpStatus: status, Data: gin.H{"rescode": res, "data": nil, "msg": msg}})
		}
		if user.UserType != 2 {
			panic(&ExceptionResult{HttpStatus: 403, Data: gin.H{"rescode": rescode.Access_Denied, "data": nil, "msg": "非供应商用户，无权限访问"}})
		}
		//请求上下文设置user
		c.Set(auth.AuthUser, user)
		c.Next()
	}
}

//运营授权
func AdminAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, status, res, msg := auth.GetUser(c)
		if user == nil {
			panic(&ExceptionResult{HttpStatus: status, Data: gin.H{"rescode": res, "data": nil, "msg": msg}})
		}
		if user.UserType != 3 && user.UserType != 4 {
			panic(&ExceptionResult{HttpStatus: 403, Data: gin.H{"rescode": rescode.Access_Denied, "data": nil, "msg": "非管理员用户，无权限访问"}})
		}
		//请求上下文设置user
		c.Set(auth.AuthUser, user)
		c.Next()
	}
}

//超级管理员授权
func SuperAdminAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, status, res, msg := auth.GetUser(c)
		if user == nil {
			panic(&ExceptionResult{HttpStatus: status, Data: gin.H{"rescode": res, "data": nil, "msg": msg}})
		}
		if user.UserType != 4 {
			panic(&ExceptionResult{HttpStatus: 403, Data: gin.H{"rescode": rescode.Access_Denied, "data": nil, "msg": "非超级管理员用户，无权限访问"}})
		}
		//请求上下文设置user
		c.Set(auth.AuthUser, user)
		c.Next()
	}
}
