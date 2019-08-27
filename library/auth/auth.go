package auth

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	jwt_lib "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/pbkdf2"
	"gosharp/library/encrypt"
	"strconv"
	"strings"
	"time"
)

const (
	JwtSecret    = "mysupersecretpassword"
	SessionStore = "session_store"
	AuthUser     = "user"
)

type JwtToken struct {
	Uid int    `json:"uid"`
	Sid string `json:"sid"`
	Iat int64  `json:"iat"`
}

//生成access_token
func GenAccessToken(userId int, sid string) string {
	token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
	// Set some claims
	token.Claims = jwt_lib.MapClaims{
		"uid": userId,
		"sid": sid,
		"iat": time.Now().Unix(),
		//"exp": time.Now().Add(time.Hour * 1).Unix(),
	}
	// Sign and get the complete encoded token as a string
	tokenString, _ := token.SignedString([]byte(JwtSecret))
	return tokenString
}

func ParseAccessToken(rawAccessToken string) (*JwtToken, error) {
	token, err := jwt_lib.Parse(rawAccessToken, func(token *jwt_lib.Token) (interface{}, error) {
		b := []byte(JwtSecret)
		return b, nil
	})
	if err != nil {
		return nil, errors.New("非法的token")
	}
	claims, ok := token.Claims.(jwt_lib.MapClaims)
	if !(ok && token.Valid) {
		return nil, errors.New("非法的token")
	}

	jwtToken := &JwtToken{
		Uid: int(claims["uid"].(float64)),
		Sid: claims["sid"].(string),
		Iat: int64(claims["iat"].(float64)),
	}
	return jwtToken, nil
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
