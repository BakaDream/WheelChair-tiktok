package main

import (
	"WheelChair-tiktok/cache"
	"WheelChair-tiktok/config"
	"WheelChair-tiktok/logger"
	"WheelChair-tiktok/model"
	"WheelChair-tiktok/router"
	"WheelChair-tiktok/utils/storage"
)

func main() {
	config.LoadEnv()
	logger.Init()
	storage.Init()
	model.Init()
	cache.RedisInit()
	r := router.InitRouter()
	r.Run("0.0.0.0:80")

}
