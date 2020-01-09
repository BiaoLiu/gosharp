package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gosharp/library/ecode"
	"net/http"
	"strconv"
)

func SetPagerHeader(c *gin.Context, offset int, len int, total int) {
	contentRange := fmt.Sprintf("%d-%d", offset, offset+len)
	c.Header("X-Content-Range", contentRange)
	c.Header("X-Content-Total", strconv.Itoa(total))
}

/**
全局成功\失败
**/
func APIResponse(c *gin.Context, success bool, data interface{}, msg string) {
	var resCode string
	if success {
		resCode = strconv.Itoa(ecode.OK.Code())
		if msg == "" {
			msg = "success"
		}
	} else {
		resCode = strconv.Itoa(ecode.FAIL.Code())
		if msg == "" {
			msg = "error"
		}
	}
	c.JSON(http.StatusOK, gin.H{"rescode": resCode, "data": data, "msg": msg})
}

//自定义错误码
func APIResponseError(c *gin.Context, data interface{}, err error) {
	if apiErr, ok := err.(ecode.Code); ok {
		c.JSON(http.StatusOK, gin.H{"rescode": strconv.Itoa(apiErr.Code()), "data": data, "msg": apiErr.Message()})
	}
	c.JSON(http.StatusOK, gin.H{"rescode": strconv.Itoa(ecode.FAIL.Code()), "data": data, "msg": err.Error()})
}

//自由模式 所有参数自定义
func APIResponseFree(c *gin.Context, httpCode int, resCode string, data interface{}, msg string) {
	c.JSON(httpCode, gin.H{"rescode": resCode, "data": data, "msg": msg})
}

//400
func APIResponseBadRequest(c *gin.Context, rescode string, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{"rescode": rescode, "data": nil, "msg": msg})
}

//401
func APIResponseUnauthorized(c *gin.Context, rescode string, msg string) {
	c.JSON(http.StatusUnauthorized, gin.H{"rescode": rescode, "data": nil, "msg": msg})
}

//403
func APIResponseForbidden(c *gin.Context, rescode string, msg string) {
	c.JSON(http.StatusForbidden, gin.H{"rescode": rescode, "data": nil, "msg": msg})
}

//404
func APIResponseNotFound(c *gin.Context, rescode string, msg string) {
	c.JSON(http.StatusNotFound, gin.H{"rescode": rescode, "data": nil, "msg": msg})
}

//405
func APIResponseNotAllowed(c *gin.Context, rescode string, msg string) {
	c.JSON(http.StatusMethodNotAllowed, gin.H{"rescode": rescode, "data": nil, "msg": msg})
}

//406
func APIResponseNotAcceptable(c *gin.Context, rescode string, msg string) {
	c.JSON(http.StatusNotAcceptable, gin.H{"rescode": rescode, "data": nil, "msg": msg})
}

//500
func APIResponseException(c *gin.Context, msg string) {
	c.JSON(http.StatusInternalServerError, gin.H{"rescode": ecode.FAIL.Code(), "data": nil, "msg": msg})
}
