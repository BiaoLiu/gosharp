package http

import (
	"crypto/tls"
	"errors"
	"fmt"
	"gosharp/library/log"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
)

func SendRequest(method string, url string, token string, body io.Reader) (interface{}, error) {
	req, err := http.NewRequest(method, url, body)
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	//忽略证书校验
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	js, err := simplejson.NewJson(result)
	if err != nil {
		return nil, err
	}
	ecode := js.Get("rescode").MustString()
	data, _ := js.Get("data").MarshalJSON()
	emsg := js.Get("msg").MustString()
	log.Logger.Debug(string(data))
	if ecode == "10000" {
		return data, nil
	}
	return data, errors.New(emsg)
}

func SetPagingHeader(c *gin.Context, offset int, len int, total int) {
	content_range := fmt.Sprintf("%d-%d", offset, offset+len)
	c.Header("Access-Control-Expose-Headers", "X-Content-Range,X-Content-Total")
	c.Header("X-Content-Range", content_range)
	c.Header("X-Content-Total", strconv.Itoa(total))
}
