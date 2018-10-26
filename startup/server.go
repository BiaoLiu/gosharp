package server

import (
	"gosharp/config"
	"gosharp/utils/log"
	"gosharp/utils/validation"
)

func Init(configPath string, logPath string) {
	//配置文件初始化
	config.Init(configPath)
	//日志初始化
	log.Init(logPath)

	//设置beego validation 错误信息
	validation.SetValidationMessage()
}
