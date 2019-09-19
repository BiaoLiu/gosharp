package auth

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/pbkdf2"
	"gosharp/library/cache/redis"
	cookie_util "gosharp/library/cookie"
	"gosharp/library/encrypt"
	e "gosharp/library/error"
	"strconv"
	"strings"
	"time"
)

const (
	JwtSecret        = "mysupersecretpassword"
	SessionKeyApp    = "session:app"
	SessionKeyManage = "session:manage"
	AuthUser         = "user"
	ManageUser       = "manage_user"
)

func SaveLoginSession(sessionKey string, uid int, v interface{}) (string, error) {
	userDict, _ := json.Marshal(v)
	//生成会话id
	sid, _ := cookie_util.Encode(uid)
	//用户信息存入redis
	d, _ := time.ParseDuration(cookie_util.CookieExpiration)
	strDict := string(userDict)
	if err := redis.Client.Set(sessionKey+":"+strconv.Itoa(uid)+":"+sid, strDict, d).Err(); err != nil {
		return "", err
	}
	if err := redis.Client.Save().Err(); err != nil {
		return "", err
	}
	return sid, nil
}

// 从接口请求token中获取登录用户
func GetLoginSession(c *gin.Context, sessionKey string, v interface{}) error {
	rawAccessToken := c.Request.Header["Authorization"]
	if len(rawAccessToken) == 0 {
		return e.APIError{}
	}
	var jwtToken *JwtToken
	var err error
	if jwtToken, err = ParseAccessToken(rawAccessToken[0]); err != nil {
		return e.APIError{}
	}
	userDict := redis.Client.Get(sessionKey + ":" + strconv.Itoa(jwtToken.Uid) + ":" + jwtToken.Sid).Val()
	if len(userDict) == 0 {
		return e.APIError{}
	}
	err = json.Unmarshal([]byte(userDict), v)
	if err != nil {
		return e.APIError{}
	}
	return nil
}

func MakePassword(password string) string {
	pwd := []byte(password)                     // 用户设置的原始密码
	salt := []byte(encrypt.GetRandomString(12)) // 盐
	iterations := 100000                        // 加密算法的迭代次数
	digest := sha256.New                        // digest 算法，使用 sha256

	dk := pbkdf2.Key(pwd, salt, iterations, 32, digest)
	str := base64.StdEncoding.EncodeToString(dk)
	// 组合加密算法、迭代次数、盐、密码和分割符号 "$"
	encoded := "pbkdf2_sha256" + "$" + strconv.FormatInt(int64(iterations), 10) + "$" + string(salt) + "$" + str
	return encoded
}

func CheckPassword(password string, encoded string) bool {
	encodeds := strings.Split(encoded, "$")

	pwd := []byte(password)                    // 用户设置的原始密码
	salt := []byte(encodeds[2])                // 盐
	iterations, _ := strconv.Atoi(encodeds[1]) // 加密算法的迭代次数
	digest := sha256.New                       // digest 算法，使用 sha256

	// 第一步：使用 pbkdf2 算法加密
	dk := pbkdf2.Key(pwd, salt, iterations, 32, digest)
	// 第二步：Base64 编码
	str := base64.StdEncoding.EncodeToString(dk)
	// 第三步：组合加密算法、迭代次数、盐、密码和分割符号 "$"
	encoded2 := "pbkdf2_sha256" + "$" + strconv.FormatInt(int64(iterations), 10) + "$" + string(salt) + "$" + str
	return encoded == encoded2
}
