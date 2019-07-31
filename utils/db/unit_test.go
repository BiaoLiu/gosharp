package db

import (
	"github.com/stretchr/testify/assert"
	"gosharp/utils/config"
	"testing"
)

func TestConnectionDatabase(t *testing.T) {
	asserts := assert.New(t)

	//配置初始化
	config.Init("../config")
	Init()

	asserts.NoError(Gorm.DB().Ping(), "Db should be able to ping")
}
