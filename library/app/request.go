package app

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gosharp/library/validation"
	"strings"
)

//表单绑定与表单验证
func BindAndValidate(c *gin.Context, obj interface{}) error {
	if err := bind(c, obj); err != nil {
		return err
	}

	var valid = validation.Validation{}
	ok, _ := valid.Valid(obj)
	if !ok {
		errorMsg := formatValidationError(valid.Errors)
		return errors.New(errorMsg)
	}
	return nil
}

//表单绑定
func bind(c *gin.Context, obj interface{}) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	if err := c.ShouldBindWith(obj, b); err != nil {
		return errors.New("参数绑定失败:" + err.Error())
	}
	return nil
}

//格式化验证器错误信息
func formatValidationError(errors []*validation.Error) string {
	err := errors[0]
	return fmt.Sprintf("%s:%s", strings.ToLower(err.Field), err.Message)
}

type Paginate struct {
	Offset int
	Limit  int
}
