package utils

import "golang.org/x/crypto/bcrypt"

// 加密
func Encrypt(text string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// 判断密码和数据库中的是否相同
func BcryptCheck(textA string, hashtextB string) bool {
	hashTextA, _ := Encrypt(textA)
	if hashTextA == hashtextB {
		return true
	} else {
		return false
	}
}
