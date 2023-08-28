package cache

import (
	"errors"
	"github.com/redis/go-redis/v9"
	"strconv"
)

func SetToken(userId uint, token string) {
	key := strconv.Itoa(int(userId))
	err := RDB.Set(key, token)
	if err != nil {
		SetToken(userId, token)
	}
	return
}
func GetToken(userId uint) (string, error) {
	key := strconv.Itoa(int(userId))
	token, err := RDB.Get(key)
	// 库中有token 返回token
	if err == nil {
		return token.(string), nil
	}
	// 库中没有token 返回错误
	if err == redis.Nil {
		return "", errors.New("token not found")
	}
	// 未知的一些错误
	return "", err
}
