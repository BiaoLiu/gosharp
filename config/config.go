package config

import (
	"github.com/spf13/viper"
	"path/filepath"
)

var Viper *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init() {
	var err error
	v := viper.New()
	//v.SetConfigType("yaml")
	v.AutomaticEnv()
	v.SetConfigName("conf")
	//v.AddConfigPath("../config/")
	v.AddConfigPath("config/")
	err = v.ReadInConfig()
	if err != nil {
		panic("error on parsing configuration file")
	}
	Viper = v
}

func relativePath(basedir string, path *string) {
	p := *path
	if p != "" && p[0] != '/' {
		*path = filepath.Join(basedir, p)
	}
}

func GetConfig() *viper.Viper {
	return Viper
}
