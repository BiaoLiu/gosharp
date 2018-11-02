package controllers

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gosharp/forms"
	"gosharp/utils/rescode"
	"net/http"
	"strings"
)

//表单绑定与表单验证
func bindAndValidateForm(c *gin.Context, form forms.Form) bool {
	if !bindForm(c, form) {
		return false
	}
	if ok, errors := form.IsValid(form); !ok {
		errorMsg := formatErrorMsg(errors)
		APIResponse(c, false, nil, errorMsg)
		return false
	}
	return true
}

//表单绑定
func bindForm(c *gin.Context, form forms.Form) bool {
	b := binding.Default(c.Request.Method, c.ContentType())
	if err := c.ShouldBindWith(form, b); err != nil {
		err := err.Error()
		APIResponseException(c, "参数绑定失败:"+err)
		return false
	}
	return true
}

//格式化错误信息
func formatErrorMsg(errors []*validation.Error) string {
	err := errors[0]
	return fmt.Sprintf("%s:%s", strings.ToLower(err.Field), err.Message)
}

//设置统一返回错误信息
func CheckError(c *gin.Context, errors []*validation.Error) {
	errorMsg := formatErrorMsg(errors)
	APIResponse(c, false, nil, errorMsg)
	return
}

/**
全局成功\失败
**/
func APIResponse(c *gin.Context, success bool, data interface{}, msg string) {
	var resCode string
	if success {
		resCode = rescode.Success
		if msg == "" {
			msg = "操作成功"
		}
	} else {
		resCode = rescode.Error
		if msg == "" {
			msg = "操作失败"
		}
	}
	c.JSON(http.StatusOK, gin.H{"rescode": resCode, "data": data, "msg": msg})
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
	c.JSON(http.StatusInternalServerError, gin.H{"rescode": rescode.Error, "data": nil, "msg": msg})
}
