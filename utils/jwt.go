package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var JwtSignedKey = []byte(os.Getenv("JWT_SIGNED_KEY"))

type JwtCustomClaims struct {
	ID   uint
	Name string
	jwt.RegisteredClaims
}

// GenerateToken 生成token
func GenerateToken(id uint, name string) (string, error) {
	iJwtCustomClaims := JwtCustomClaims{
		ID:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(365 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "Token",
		}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, iJwtCustomClaims)
	return token.SignedString(JwtSignedKey)
}

// ParseToken 解析token
func ParseToken(tokenStr string) (JwtCustomClaims, error) {
	iJwtCustomClaims := JwtCustomClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, &iJwtCustomClaims, func(token *jwt.Token) (interface{}, error) {
		return JwtSignedKey, nil
	})
	if err == nil && !token.Valid {
		err = errors.New("invalid token")
	}
	return iJwtCustomClaims, err
}

// IsTokenValid 判断token是否合法
func IsTokenValid(tokenStr string) bool {
	_, err := ParseToken(tokenStr)
	if err != nil {
		return false
	}
	return true
}
