package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gosharp/library/config"
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
	Gorm.DB().SetMaxIdleConns(10)
	Gorm.DB().SetMaxOpenConns(100)
}

func Close() {
	Gorm.Close()
}

// WhereIgnoreBlank ?????,?????
func WhereIgnoreBlank(query *gorm.DB, cond string, value string) *gorm.DB {
	if value != "" {
		query = query.Where(cond, value)
	}
	return query
}

// WhereLikeIgnoreBlank: ?????like ??
func WhereLikeIgnoreBlank(query *gorm.DB, cond string, value string) *gorm.DB {
	if value != "" {
		query = query.Where(cond, "%"+value+"%")
	}
	return query
}

// PaginateWithCount
func PaginateWithCount(query *gorm.DB, order string, offset, limit int, objs interface{}) (int, error) {
	var total int
	if err := query.Count(&total).Error; err != nil {
		return 0, err
	}

	if order != "" {
		query = query.Order(order)
	}
	if err := query.Offset(offset).Limit(limit).Find(objs).Error; err != nil {
		return 0, err
	}
	return total, nil
}
