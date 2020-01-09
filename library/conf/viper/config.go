package config

import (
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
)

var Viper *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(configPath string) {
	var err error
	v := viper.New()
	//v.SetConfigType("yaml")
	v.AutomaticEnv()
	v.SetConfigName("conf")
	//v.AddConfigPath("../config/")
	if configPath == "" {
		v.AddConfigPath("config/")
	} else {
		v.AddConfigPath(configPath)
	}

	err = v.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("error on parsing configuration file: %s", err.Error()))
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
