package server

import (
	"gosharp/library/config"
	"gosharp/library/database/mysql"
	"gosharp/library/log"
	"gosharp/library/pkg/validation"
)

func Init(configPath string, logPath string) {
	//配置文件初始化
	config.Init(configPath)
	//日志初始化
	log.Init(logPath)
	//数据库初始化
	db.Init()
	//设置beego validation 错误信息
	validation.SetValidationMessage()
}
