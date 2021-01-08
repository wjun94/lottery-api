package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type CustomClaims struct {
	UserID string `json:"userid"`
	Type   byte   `json:"type"`
	jwt.StandardClaims
}

type WxCustomClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

var expireTime = 60 * 24 * 15 // 过期时间15天
var key = []byte("adfadf!@#2")

// CreateToken 创建token
func CreateToken(userId string, level byte) string {
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
func GetUserID(tokenstr string) (string, byte) {
	token := Parse(tokenstr)
	if token == nil {
		return "", 0
	}
	if claims, ok := token.Claims.(*CustomClaims); ok {
		return claims.UserID, claims.Type
	}
	return "", 0
}

// CreateWxToken 创建token
func CreateWxToken(id string) string {
	claims := WxCustomClaims{
		id, jwt.StandardClaims{
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

// GetWxID 获取用户id
func GetWxID(tokenstr string) string {
	token := ParseWx(tokenstr)
	if token == nil {
		return ""
	}
	if claims, ok := token.Claims.(*WxCustomClaims); ok {
		return claims.ID
	}
	return ""
}

// Parse 解析后台token
func Parse(tokenstr string) *jwt.Token {
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

// ParseWx 解析微信token
func ParseWx(tokenstr string) *jwt.Token {
	token, err := jwt.ParseWithClaims(tokenstr, &WxCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
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

// DecryptToken 解析token,返回token对应的id
// func DecryptToken(r *http.Request) string {
// 	tokenStr := r.Header.Get("authorization")
// 	if tokenStr == "" {
// 		return ""
// 	}
// 	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
// 		return []byte("secret"), nil
// 	})
// 	if !token.Valid {
// 		return ""
// 	}
// 	// 进入下一步
// 	return getIdFromClaims("username", token.Claims)
// }

/* import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"reflect"
)

// GenerateToken 生成token
func GenerateToken(value string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": value,
		//"exp":      time.Now().Add(time.Hour * 2).Unix(),// 可以添加过期时间
	})

	return token.SignedString([]byte("secret")) //对应的字符串请自行生成，最后足够使用加密后的字符串
}

// getIdFromClaims 获取token对应的id
func getIdFromClaims(key string, claims jwt.Claims) string {
	v := reflect.ValueOf(claims)
	if v.Kind() == reflect.Map {
		for _, k := range v.MapKeys() {
			value := v.MapIndex(k)
			if fmt.Sprintf("%s", k.Interface()) == key {
				return fmt.Sprintf("%v", value.Interface())
			}
		}
	}
	return ""
}


*/
