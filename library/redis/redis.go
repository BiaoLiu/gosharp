package redis

import (
	"fmt"
	"gopkg.in/redis.v5"
	"gosharp/library/config"
)

var Client *redis.Client

//获取连接
func Init() {
	address := config.Viper.GetString("REDIS_ADDRESS")
	password := config.Viper.GetString("REDIS_PASSWORD")
	db := config.Viper.GetInt("REDIS_DB")

	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
		PoolSize: 10,
	})

	err := client.Ping().Err()
	if err != nil {
		panic(fmt.Sprintf("redis connection error. %s", err.Error()))
	}

	Client = client
}

func Close() {
	Client.Close()
}
