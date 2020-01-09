package unit_test

import (
	"net/http"
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
	//asserts := assert.New(t)
	//engine := NewEngine()
	//
	//for _, testData := range requestTests {
	//	//创建一个请求
	//	bodyData := testData.bodyData
	//	req, err := http.NewRequest(testData.method, testData.url, bytes.NewBufferString(bodyData))
	//	req.Header.Set("Content-Type", "application/json")
	//	asserts.NoError(err)
	//
	//	w := httptest.NewRecorder()
	//	engine.ServeHTTP(w, req)
	//
	//	result, err := ioutil.ReadAll(w.Body)
	//	j, err := simplejson.NewJson(result)
	//	asserts.NoError(err, "response error")
	//
	//	asserts.Equal(testData.rescode, j.Get("rescode").MustString(), "Response Content - "+j.Get("msg").MustString())
	//	asserts.Equal(200, w.Code, "Response status -"+w.Body.String())
	//}
}
