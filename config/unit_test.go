package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig(t *testing.T) {
	asserts := assert.New(t)

	Init("../config")

	asserts.NotNil(Viper, "config init error")
}
