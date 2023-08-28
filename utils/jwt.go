package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var JwtSignedKey = []byte(os.Getenv("JWT_SIGNED_KEY"))

type Claims struct {
	ID       uint
	UserName string
	jwt.RegisteredClaims
}

/*
GenerateToken
params: id 用户id
return: token字符串 错误信息
*/
func GenerateToken(id uint, userName string) (string, error) {
	iClaims := Claims{
		ID:       id,
		UserName: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(365 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "Token",
		}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, iClaims)
	return token.SignedString(JwtSignedKey)
}

// ParseToken 解析token
func ParseToken(tokenStr string) (Claims, error) {
	iClaims := Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, &iClaims, func(token *jwt.Token) (interface{}, error) {
		return JwtSignedKey, nil
	})
	if err == nil && !token.Valid {
		err = errors.New("invalid token")
	}
	return iClaims, err
}

// IsTokenValid 判断token是否合法
func IsTokenValid(tokenStr string) bool {
	_, err := ParseToken(tokenStr)
	if err != nil {
		return false
	}
	return true
}
