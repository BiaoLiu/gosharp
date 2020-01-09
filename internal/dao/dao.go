package dao

import (
	"context"
	"github.com/jinzhu/gorm"
	"gosharp/library/conf/viper"
	"gosharp/library/database/orm"
)

// Dao struct info of Dao.
type Dao struct {
	db *gorm.DB
	//redis       *redis.Pool
	//redisExpire int32
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// New new a dao and return.
func New() (d *Dao) {
	dc := &orm.Config{
		DSN:    config.Viper.GetString("MYSQL_DSN"),
		Active: 10,
		Idle:   100,
		//IdleTimeout: "",
	}
	d = &Dao{
		// mysql
		db: orm.NewMySQL(dc),
	}
	return
}

// DB .
func (d *Dao) DB() *gorm.DB {
	return d.db
}

// Ping ping the resource.
func (d *Dao) Ping(ctx context.Context) (err error) {
	return d.db.DB().Ping()
}

// Close close the resource.
func (d *Dao) Close() {
	d.db.Close()
}
