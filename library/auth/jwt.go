package auth

import (
	"errors"
	jwt_lib "github.com/dgrijalva/jwt-go"
	"time"
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
