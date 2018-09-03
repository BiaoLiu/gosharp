package auth

import (
	"crypto/tls"
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	"gosharp/config"
	"gosharp/utils/log"
	"gosharp/utils/rescode"
	"io/ioutil"
	"net/http"
)

const (
	AuthUser = "user"
)

type User struct {
	Id          int
	Username    string
	Mobile      string
	Email       string
	UserType    int `json:"user_type"`
	IsSubuser   int `json:"is_subuser"`
	MainUserId  int
	Permissions []string
}

func GetUser(c *gin.Context) (user *User, status int, res string, msg string) {
	token := c.Request.Header["Authorization"]
	if len(token) == 0 {
		return nil, 400, rescode.Token_Missing, "token不存在"
	}

	SSO_VERIFY := config.Viper.GetString("SSO_HOST") + "/server/verify"
	req, err := http.NewRequest("GET", SSO_VERIFY, nil)
	req.Header.Set("Authorization", token[0])

	//忽略证书校验
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		log.Logger.Error("sso verify request error:", err.Error())
		return nil, 500, rescode.Error, "服务器异常，登录授权验证失败"
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	js, err := simplejson.NewJson(result)
	data, err := js.Get("data").MarshalJSON()
	json.Unmarshal(data, &user)

	return user, 200, "", ""
}
