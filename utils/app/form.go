package app

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gosharp/utils/validation"
	"strings"
)

type Form interface {
	IsValid(form Form) (bool, []*validation.Error)
	//Save() interface{}
}

type BaseForm struct {
}

func (form *BaseForm) Valid(v *validation.Validation) {

}

//表单验证
func (form *BaseForm) IsValid(validForm Form) (bool, []*validation.Error) {
	var valid = validation.Validation{}
	ok, _ := valid.Valid(validForm)
	if !ok {
		return false, valid.Errors
	}
	return true, nil
}

//表单保存操作
//func (form *BaseForm) Save() interface{} {
//	return nil
//}

//表单绑定与表单验证
func BindAndValidate(c *gin.Context, form Form) error {
	if err := bindForm(c, form); err != nil {
		return err
	}
	if ok, validationErrors := form.IsValid(form); !ok {
		errorMsg := formatValidationError(validationErrors)
		return errors.New(errorMsg)
	}
	return nil
}

//表单绑定
func bindForm(c *gin.Context, form Form) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	if err := c.ShouldBindWith(form, b); err != nil {
		return errors.New("参数绑定失败:" + err.Error())
	}
	return nil
}

//格式化验证器错误信息
func formatValidationError(errors []*validation.Error) string {
	err := errors[0]
	return fmt.Sprintf("%s:%s", strings.ToLower(err.Field), err.Message)
}
