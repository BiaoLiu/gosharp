package app

import (
	"github.com/gin-gonic/gin"
	"gosharp/library/auth"
	"gosharp/library/type"
	"gosharp/models"
)

type Paginate struct {
	Offset int
	Limit  int
}

func ParamInt(c *gin.Context, name string) int {
	return utils.SafeInt(c.Param(name), 0)
}

// 从请求上下文中获取App登录用户
func AppUser(c *gin.Context) *models.AuthUser {
	user, _ := c.MustGet(auth.AuthUser).(*models.AuthUser)
	return user
}

// 从请求上下文中获取App登录用户id
func AppUserId(c *gin.Context) int {
	user, _ := c.MustGet(auth.AuthUser).(*models.AuthUser)
	return user.ID
}

func Pager(c *gin.Context) *Paginate {
	offset := utils.SafeInt(c.Query("offset"), 0)
	limit := utils.SafeInt(c.Query("limit"), 20)

	paginate := &Paginate{
		offset,
		limit,
	}
	return paginate
}
