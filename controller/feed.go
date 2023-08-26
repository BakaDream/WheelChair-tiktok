package controller

import (
	g "WheelChair-tiktok/global"
	m "WheelChair-tiktok/model"
	"WheelChair-tiktok/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

var currentPage = 1 //全局变量记录当前page

func Feed(c *gin.Context) { // 默认每页加载 15 个视频
	tokenString := c.Query("token")
	if tokenString == "" {
		videos := makeGuestVideoList(currentPage, g.MaxPerPage)
		c.JSON(http.StatusOK, m.FeedResponse{
			StatusCode: 0,
			VideoList:  videos,
			NextTime:   time.Now().Unix(),
		})
	} else {
		uid, _ := utils.GetUserID(tokenString) //调用方法返回视频列表
		videos := makeVideoList(currentPage, g.MaxPerPage, uid)
		c.JSON(http.StatusOK, m.FeedResponse{
			StatusCode: 0,
			VideoList:  videos,
			NextTime:   time.Now().Unix(),
		})
	}
	currentPage++
}

func makeGuestVideoList(page, perPage int) []m.Video {
	offSet := (page - 1) * perPage //offSet:视频开始位置
	var videos []m.Video
	// 构建查询
	err := g.DB.Order("CreateAt DESC").Limit(perPage).Offset(offSet).Find(&videos)
	if err != nil {
		fmt.Println("Failed to execute query:", err)
		return nil
	}
	if len(videos) < perPage {
		currentPage = 0
	}
	for i := range videos {
		videos[i].IsFavorite = false
	}
	return videos //返回视频列表
}

func makeVideoList(page, perPage int, uid uint) []m.Video {
	offSet := (page - 1) * perPage //offSet:视频开始位置
	var videos []m.Video
	err := g.DB.Order("publish_time DESC").Limit(perPage).Offset(offSet).Find(&videos)
	if err != nil {
		fmt.Println("Failed to execute query:", err)
		return nil
	}
	if len(videos) < perPage {
		currentPage = 0
	}
	for i, video := range videos {
		result := g.DB.Where("VideoID = ? AND UserID", video.ID, uid).First(&m.UserVideoLike{})
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				// 没有找到记录
				videos[i].IsFavorite = false
			} else {
				// 查询过程中发生了其他错误
				log.Fatal("Unknown error")
			}
		} else {
			videos[i].IsFavorite = true
		}
	}
	return videos //返回视频列表
}
