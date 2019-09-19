package cookie_util

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"net/http"
	"time"
)

const (
	CookieName       = "gosharp"
	CookieSecret     = "gosharp@2019"
	CookieExpiration = "72h"
)

func Encode(value interface{}) (string, error) {
	return securecookie.EncodeMulti(CookieName, value, securecookie.CodecsFromPairs([]byte(CookieSecret))...)
}

func Decode(value string, dst interface{}) error {
	return securecookie.DecodeMulti(CookieName, value, dst, securecookie.CodecsFromPairs([]byte(CookieSecret))...)
}

func SetCookie(c *gin.Context, name string, value string) {
	d, _ := time.ParseDuration("72h")
	cookie := &http.Cookie{
		Name:     name,
		Path:     "/",
		Value:    value,
		HttpOnly: true,
		Expires:  time.Now().Add(d),
	}
	http.SetCookie(c.Writer, cookie)
}
