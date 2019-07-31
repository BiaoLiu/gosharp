package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig(t *testing.T) {
	asserts := assert.New(t)

	Init("../../config")

	asserts.NotNil(Viper, "config init error")
}

//func TestRemote(t *testing.T) {
//	v := viper.New()
//	v.AddRemoteProvider("consul", "consul.api.com", "")
//	v.SetConfigType("json")
//	err := viper.ReadRemoteConfig()
//
//	fmt.Println(err)
//
//	fmt.Println(v.Get("port"))
//	fmt.Println(v.Get("hostname"))
//	value := v.Get("test/MYSQL_HOST")
//	fmt.Println(v.AllKeys())
//	fmt.Println(value)
//}