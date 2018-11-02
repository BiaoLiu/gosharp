package forms

import (
	"github.com/astaxie/beego/validation"
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
