package orm

import (
	"gosharp/library/ecode"
	"gosharp/library/log"
	"strings"
	// database driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Config mysql config.
type Config struct {
	DSN    string // data source name.
	Active int    // pool
	Idle   int    // pool
	//IdleTimeout xtime.Duration // connect max life time.
}

type ormLog struct{}

func (l ormLog) Print(v ...interface{}) {
	log.Logger.Info(strings.Repeat("%v ", len(v)), v)
}

func init() {
	gorm.ErrRecordNotFound = ecode.NothingFound
}

// NewMySQL new db and retry connection when has error.
func NewMySQL(c *Config) (db *gorm.DB) {
	db, err := gorm.Open("mysql", c.DSN)
	if err != nil {
		log.Logger.Error("db dsn(%s) error: %v", c.DSN, err)
		panic(err)
	}
	db.DB().SetMaxIdleConns(c.Idle)
	db.DB().SetMaxOpenConns(c.Active)
	db.LogMode(true)
	db.SingularTable(true)
	//db.DB().SetConnMaxLifetime(time.Duration(c.IdleTimeout) / time.Second)
	//db.SetLogger(ormLog{})
	return
}

// Where
func Where(query *gorm.DB, cond string, value string) *gorm.DB {
	query = query.Where(cond, value)
	return query
}

// WhereIgnoreBlank ?????,?????
func WhereIgnoreBlank(query *gorm.DB, cond string, value string) *gorm.DB {
	if value != "" {
		query = query.Where(cond, value)
	}
	return query
}

// WhereIgnoreZero
func WhereIgnoreZero(query *gorm.DB, cond string, value interface{}) *gorm.DB {
	switch t := value.(type) {
	case int:
		if t <= 0 {
			return query
		}
	case int64:
		if t <= 0 {
			return query
		}
	case float64:
		if t <= 0 {
			return query
		}
	case *int:
		if t == nil || *t == -1 {
			return query
		}
	case *int64:
		if t == nil || *t == -1 {
			return query
		}
	case *float64:
		if t == nil || *t == -1 {
			return query
		}
	default:
		return query
	}
	query = query.Where(cond, value)
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
