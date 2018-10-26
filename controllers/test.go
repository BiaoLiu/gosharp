package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	//_ "github.com/spf13/viper/remote"
	"gosharp/db"
	"gosharp/models"
)

func TestConsul(c *gin.Context) {
	var user models.AuthUser
	db.Gorm.Where("id=1").First(&user)
	fmt.Println(user.Username)

	v := viper.New()
	v.AddRemoteProvider("consul", "consul.robo2025.com", "MY_CONSUL_KEY")
	v.SetConfigType("json")
	err := v.ReadRemoteConfig()
	fmt.Println(err)

	host := v.Get("test/MYSQL_HOST")
	fmt.Println(host)
	APIResponse(c, true, nil, "")
}
