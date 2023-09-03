package cache

import (
	"WheelChair-tiktok/logger"
	"context"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"os"
	"time"
)

type RedisClient struct {
	RC *redis.Client
}

var RDB = RedisClient{}

var nDuration = 30 * 24 * 60 * 60 * time.Second

func RedisInit() {
	addr := os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")
	rc := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: os.Getenv("REDISCLI_AUTH"),
		DB:       0,
	})
	_, err := rc.Ping(context.Background()).Result()
	if err != nil {
		logger.Logger.Fatal("Redis init failed", zap.Error(err))
	}
	RDB.RC = rc

}
func (rc *RedisClient) Set(key string, value any) error {
	return RDB.RC.Set(context.Background(), key, value, nDuration).Err()
}

func (rc *RedisClient) Get(key string) (any, error) {
	return RDB.RC.Get(context.Background(), key).Result()
}

func (rc *RedisClient) Delete(key ...string) error {
	return RDB.RC.Del(context.Background(), key...).Err()
}
