package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnectionDatabase(t *testing.T) {
	asserts := assert.New(t)
	asserts.NoError(Gorm.DB().Ping(), "Db should be able to ping")
}
