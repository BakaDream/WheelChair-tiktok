package controller

import (
	m "WheelChair-tiktok/model"
	resp "WheelChair-tiktok/model/response"
	"WheelChair-tiktok/service"
	"WheelChair-tiktok/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

// 上传视频
func Publish(c *gin.Context) {
	title := c.PostForm("title")
	file, err := c.FormFile("data")
	token := c.PostForm("token")
	iClaims, _ := utils.ParseToken(token)
	uid := iClaims.ID
	if err != nil {
		return
	}

	//检测文件是否为视频文件
	if utils.IsVideoFile(file.Filename) {
		c.JSON(http.StatusOK, resp.Publish{
			StatusCode: 1,
			StatusMsg:  "is’t a valid video file",
		})
		return
	}

	//给视频一个hashName
	file.Filename = utils.HashFileName(file.Filename)
	//保存文件
	playUrl, coverUrl, err := service.UploadVideo(file)
	if err != nil {
		c.JSON(http.StatusOK, m.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	result := m.DB.Create(&m.Video{UserID: uid, Title: title, FavoriteCount: 0, CommentCount: 0, PlayUrl: playUrl, CoverUrl: coverUrl})
	if result != nil {
		log.Fatal("An error occurred in the creation")
	}
	result = m.DB.Model(&m.User{}).Update("WorkCount", gorm.Expr("WorkCount + ?", 1))
	if result != nil {
		c.JSON(http.StatusOK, m.Response{
			StatusCode: 1,
		})
		return
	}
	c.JSON(http.StatusOK, m.Response{
		StatusCode: 0,
		StatusMsg:  "success",
	})
}

// PublishList shows user's published videos
func PublishList(c *gin.Context) {
	if !utils.CheckToken(c.Query("token")) {
		c.JSON(http.StatusOK, m.Response{StatusCode: 1, StatusMsg: "Unauthorized request"})
		return
	}
	UserID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		log.Println("The type of VideoID is incorrect")
		return
	}
	var videoList []m.Video
	result := m.DB.Where("UserID = ?", UserID).Find(&videoList)
	if result.Error != nil {
		log.Fatal("Unknow error")
	}
	c.JSON(http.StatusOK, m.UserPublishResponse{
		StatusCode: 0,
		VideoList:  videoList,
	})
}
