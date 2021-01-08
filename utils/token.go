package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	UserID string `json:"userid"`
	Level  byte   `json:"level"`
	jwt.StandardClaims
}

var expireTime = 60 * 24 * 15 // 过期时间15天
var key = []byte("adfadf!@#2")

// CreateToken 创建token
func (this *Utils) CreateToken(userId string, level byte) string {
	claims := CustomClaims{
		userId, level, jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(expireTime)).Unix(),
			Issuer:    "xwj",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstr, err := token.SignedString(key) //对应的字符串请自行生成，最后足够使用加密后的字符串

	if err == nil {
		return tokenstr
	}
	return ""
}

// GetUserID 获取用户id
func (this *Utils) GetUserID(tokenstr string) (string, byte) {
	token := this.Parse(tokenstr)
	if token == nil {
		return "", 0
	}
	if claims, ok := token.Claims.(*CustomClaims); ok {
		return claims.UserID, claims.Level
	}
	return "", 0
}

// Parse 解析token
func (this *Utils) Parse(tokenstr string) *jwt.Token {
	token, err := jwt.ParseWithClaims(tokenstr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return nil
	}
	result := token.Valid
	if result {
		return token
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			println("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			println("The token is expired or not valid.")
		} else {
			println("Couldn't handle this token:", err)
		}
	} else {
		println("Couldn't handle this token:", err)
	}
	return nil
}
