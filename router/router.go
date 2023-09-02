package router

import (
	"WheelChair-tiktok/controller"
	"WheelChair-tiktok/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/public", "./public")
	// 主路由组
	douyinGroup := r.Group("/douyin")
	{
		// user路由组
		userGroup := douyinGroup.Group("/user")
		{
			userGroup.POST("/register/", controller.Register)
			userGroup.GET("/", middleware.Auth(), controller.UserInfo)
			userGroup.POST("/login/", controller.Login)
		}
		//
		// publish路由组
		publishGroup := douyinGroup.Group("/publish")
		{
			publishGroup.POST("/action/", middleware.Auth(), controller.Publish)
			publishGroup.GET("/list/", middleware.Auth(), controller.PublishList)

		}

		//feed
		douyinGroup.GET("/feed/", middleware.Auth(), controller.Feed)

		favoriteGroup := douyinGroup.Group("favorite")
		{
			favoriteGroup.POST("/action/", middleware.Auth(), controller.FavoriteAction)
			favoriteGroup.GET("/list/", middleware.Auth(), controller.FavoriteList)
		}

		// comment路由组
		commentGroup := douyinGroup.Group("/comment")
		{
			commentGroup.POST("/action/", middleware.Auth(), controller.CommentAction)
			commentGroup.GET("/list/", middleware.Auth(), controller.CommentList)
		}
		//
	}

	return r
}
