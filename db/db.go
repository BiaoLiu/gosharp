package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"gosharp/config"
	//"github.com/garyburd/redigo/redis"
	"fmt"
)

var Xorm *xorm.Engine

func Init() {
	var err error
	Xorm, err = xorm.NewEngine("mysql", config.Viper.GetString("MYSQL_DSN"))
	Xorm.ShowSQL(true)

	if err != nil {
		panic(fmt.Sprintf("mysql init error. %s", err.Error()))
	}
}

func Close() {
	Xorm.Close()
}
