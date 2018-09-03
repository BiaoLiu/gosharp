package validation

import "github.com/astaxie/beego/validation"

var messageTmpls = map[string]string{
	"Required":     "不能为空",
	"Min":          "最小值为 %d",
	"Max":          "最大值为 %d",
	"Range":        "Range is %d to %d",
	"MinSize":      "最小长度是%d",
	"MaxSize":      "最大长度是%d",
	"Length":       "Required length is %d",
	"Alpha":        "Must be valid alpha characters",
	"Numeric":      "Must be valid numeric characters",
	"AlphaNumeric": "Must be valid alpha or numeric characters",
	"Match":        "Must match %s",
	"NoMatch":      "Must not match %s",
	"AlphaDash":    "Must be valid alpha or numeric or dash(-_) characters",
	"Email":        "必须是一个合法的邮箱地址",
	"IP":           "Must be a valid ip address",
	"Base64":       "Must be valid base64 characters",
	"Mobile":       "Must be valid mobile number",
	"Tel":          "Must be valid telephone number",
	"Phone":        "Must be valid telephone or mobile phone number",
	"ZipCode":      "Must be valid zipcode",
}

func SetValidationMessage() {
	validation.SetDefaultMessage(messageTmpls)
}
