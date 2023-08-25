package cache

import (
	"WheelChair-tiktok/global"
	"context"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"os"
	"time"
)

type RedisClient struct {
	RC *redis.Client
}

var nDuration = 30 * 24 * 60 * 60 * time.Second

func RedisInit() {

	rc := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	_, err := rc.Ping(context.Background()).Result()
	if err != nil {
		global.Logger.Fatal("Redis init failed", zap.Error(err))
	}
	global.RedisClient.RC = rc

}
func (rc *RedisClient) Set(key string, value any) error {
	return global.RedisClient.RC.Set(context.Background(), key, value, nDuration).Err()
}

func (rc *RedisClient) Get(key string) (any, error) {
	return global.RedisClient.RC.Get(context.Background(), key).Result()
}

func (rc *RedisClient) Delete(key ...string) error {
	return global.RedisClient.RC.Del(context.Background(), key...).Err()
}
