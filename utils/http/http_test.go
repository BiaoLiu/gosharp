package http

import (
	"github.com/stretchr/testify/assert"
	"gosharp/config"
	"gosharp/utils/log"
	"testing"
)

func TestHttp(t *testing.T) {
	asserts := assert.New(t)

	//配置文件初始化
	config.Init("../../config")
	//日志初始化
	log.Init("logs")

	b, data := SendRequest("GET", "https://testapi.xxx.com", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MzY4MjA1MDIsInNpZCI6Ik1UVXpOamd5TURVd01YeEJkMUZCV0VFOVBYeEFtV1ZJaVJrdWtTXzFudlRKNmxEYUhlN0RrcXUycF9XZlc1akJsUjVQenc9PSIsInVpZCI6NDZ9.IJ5_ZZbL_0v_FTBqEHjQgRe9zXBhgH5CNk6ALQb0ujc", nil)
	asserts.Equal(true, b, "send request error:", data)
}
