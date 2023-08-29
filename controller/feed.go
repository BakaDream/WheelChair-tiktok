package controller

import (
	m "WheelChair-tiktok/model"
	resp "WheelChair-tiktok/model/response"
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
var MaxPerPage = 15

func Feed(c *gin.Context) { // 默认每页加载 15 个视频
	tokenString := c.Query("token")
	if tokenString == "" {
		videos := makeGuestVideoList(currentPage, MaxPerPage)
		c.JSON(http.StatusOK, resp.Feed{
			StatusCode: 0,
			VideoList:  videos,
			NextTime:   time.Now().Unix(),
		})
	} else {
		uid, _ := utils.GetUserID(tokenString) //调用方法返回视频列表
		videos := makeVideoList(currentPage, MaxPerPage, uid)
		c.JSON(http.StatusOK, resp.Feed{
			StatusCode: 0,
			VideoList:  videos,
			NextTime:   time.Now().Unix(),
		})
	}
	currentPage++
}

func makeGuestVideoList(page, perPage int) []resp.Video {
	offSet := (page - 1) * perPage //offSet:视频开始位置
	var videos []m.Video
	var videoSponse []resp.Video
	// 构建查询
	err := m.DB.Order("CreateAt DESC").Limit(perPage).Offset(offSet).Find(&videos)
	if err != nil {
		fmt.Println("Failed to execute query:", err)
		return nil
	}
	if len(videos) < perPage {
		currentPage = 0
	}
	for i := range videos {
		var author resp.User
		m.DB.First(&author, videos[i].AuthorID)
		videoSponse[i] = videos[i].ToResponse(false, author)
	}
	return videoSponse //返回视频列表
}

func makeVideoList(page, perPage int, uid uint) []resp.Video {
	offSet := (page - 1) * perPage //offSet:视频开始位置
	var videos []m.Video
	var videoSponse []resp.Video
	err := m.DB.Order("publish_time DESC").Limit(perPage).Offset(offSet).Find(&videos)
	if err != nil {
		fmt.Println("Failed to execute query:", err)
		return nil
	}
	if len(videos) < perPage {
		currentPage = 0
	}
	for i, video := range videos {
		result := m.DB.Where("VideoID = ? AND UserID", video.ID, uid).First(&m.Favorite{})
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				// 没有找到记录
				var author resp.User
				m.DB.First(&author, videos[i].AuthorID)
				videoSponse[i] = videos[i].ToResponse(false, author)
			} else {
				// 查询过程中发生了其他错误
				log.Fatal("Unknown error")
			}
		} else {
			var author resp.User
			m.DB.First(&author, videos[i].AuthorID)
			videoSponse[i] = videos[i].ToResponse(true, author)
		}
	}
	return videoSponse //返回视频列表
}
