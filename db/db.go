package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gosharp/config"
	//"github.com/garyburd/redigo/redis"
	"fmt"
)

var Gorm *gorm.DB

func Init() {
	var err error
	Gorm, err = gorm.Open("mysql", config.Viper.GetString("MYSQL_DSN"))
	if err != nil {
		panic(fmt.Sprintf("mysql init error. %s", err.Error()))
	}
	Gorm.LogMode(true)
	Gorm.SingularTable(true)
}

func Close() {
	Gorm.Close()
}
