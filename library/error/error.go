package e

import "errors"

var (
	ERR_FOREIGN_KEY_VIOLATION = errors.New("已经被使用了不能删除")
	ERR_DUPLIATED_DATA        = errors.New("存在重复的数据")
	ERR_DATA_NOT_FOUND        = errors.New("数据未找到")
	ERR_DATA_INVALID          = errors.New("数据无效")
	ERR_STATE_CANNOT_OPERATED = errors.New("当前状态无法进行操作")
)

type APIError struct {
	S    string
	Code string
}

func (err APIError) Error() string {
	return err.S
}
