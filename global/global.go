package global

import (
	"WheelChair-tiktok/cache"
	"WheelChair-tiktok/utils/storage"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

var Logger *zap.Logger
var DSN = os.Getenv("MYSQL_DSN")
var DB *gorm.DB
var RedisClient *cache.RedisClient
var Storage storage.Store
