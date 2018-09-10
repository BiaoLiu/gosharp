package unit_test

import (
	"bytes"
	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gosharp/config"
	"gosharp/db"
	"gosharp/middlewares"
	"gosharp/routers"
	"gosharp/startup"
	"gosharp/utils/log"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var requestTests = []struct {
	init     func(*http.Request)
	url      string
	method   string
	bodyData string
	rescode  string
}{
	{
		func(request *http.Request) {},
		"/login",
		"POST",
		`{"username":"lbi","password":"lbi"}`,
		"10000",
	},
}

func TestLogin(t *testing.T) {
	asserts := assert.New(t)

	engine := gin.New()
	//配置文件初始化
	config.Init("../config")
	//日志初始化
	log.Init("../logs")
	//数据库初始化
	db.Init()
	defer db.Close()
	//gin配置初始化
	server.Init(engine)

	//注册中间件
	middlewares.Register(engine)
	//注册路由
	routers.Register(engine)

	for _, testData := range requestTests {
		//创建一个请求
		bodyData := testData.bodyData
		req, err := http.NewRequest(testData.method, testData.url, bytes.NewBufferString(bodyData))
		req.Header.Set("Content-Type", "application/json")
		asserts.NoError(err)

		w := httptest.NewRecorder()
		engine.ServeHTTP(w, req)

		result, err := ioutil.ReadAll(w.Body)
		j, err := simplejson.NewJson(result)
		asserts.NoError(err, "response error")

		asserts.Equal(testData.rescode, j.Get("rescode").MustString(), "Response Content - "+j.Get("msg").MustString())
		asserts.Equal(200, w.Code, "Response status -"+w.Body.String())
	}
}
