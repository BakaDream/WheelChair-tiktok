package global

import (
	"WheelChair-tiktok/cache"
	"WheelChair-tiktok/utils/storage"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var Logger *zap.Logger
var DB *gorm.DB
var RedisClient *cache.RedisClient
var Storage storage.Store
