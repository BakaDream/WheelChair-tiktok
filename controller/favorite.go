package controller

import (
	//g "WheelChair-tiktok/global"
	m "WheelChair-tiktok/model"
	"WheelChair-tiktok/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

func FavoriteAction(c *gin.Context) {
	token := c.Query("toekn")
	VideoID, err := strconv.Atoi(c.Query("video_id"))
	if err != nil {
		log.Println("The type of VideoID is incorrect")
		return
	}
	actionType, err := strconv.Atoi(c.Query("action_type"))
	if err != nil {
		log.Println("The type parameter is incorrect")
		return
	}
	UserID, err := utils.GetUserID(token)
	if err != nil {
		log.Println("User don't exit")
		c.JSON(http.StatusOK, m.Response{StatusCode: 1})
		return
	}
	var storeUserVideoLike m.UserVideoLike
	result := m.DB.Where("UserID = ? and VideoID = ?", UserID, VideoID).First(&storeUserVideoLike)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) && actionType == 1 {
			// 没有找到记录
			m.DB.Create(&m.UserVideoLike{VideoID: uint(VideoID), UserID: UserID})
			m.DB.Model(&m.Video{}).Where("ID = ?", VideoID).Update("FavoriteComment", gorm.Expr("FavoriteComment + ?", 1))
			var video m.Video
			m.DB.Where("ID = ?", VideoID).First(&video)
			m.DB.Model(&m.User{}).Where("ID = ?", video.UserID).Update("TotalFavorited", gorm.Expr("TotalFavorited + ?", 1))
			c.JSON(http.StatusOK, m.FavoriteActionResponse{StatusCode: 0})
		} else {
			// 查询过程中发生了其他错误
			log.Fatal("Unknown error")
		}
	} else if actionType == 2 {
		m.DB.Delete(&storeUserVideoLike)
		m.DB.Model(&m.Video{}).Where("ID = ?", VideoID).Update("FavoriteComment", gorm.Expr("FavoriteComment + ?", -1))
		var video m.Video
		m.DB.Where("ID = ?", VideoID).First(&video)
		m.DB.Model(&m.User{}).Where("ID = ?", video.UserID).Update("TotalFavorited", gorm.Expr("TotalFavorited + ?", -1))
		c.JSON(http.StatusOK, m.FavoriteActionResponse{StatusCode: 0})
	}
}
func FavoriteList(c *gin.Context) {
	UserID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		log.Println("The type of UserID is incorrect")
		return
	}
	var storeUserVideoLike []m.UserVideoLike
	result := m.DB.Where("UserID = ?", UserID).Order("CreateAt DESC").Find(&storeUserVideoLike)
	if result != nil {
		fmt.Println("Make UserVideoLikeList Error:", err)
		return
	}
	var storeVideoID []uint
	for _, Video := range storeUserVideoLike {
		storeVideoID = append(storeVideoID, Video.VideoID)
	}
	var storeVideoList []m.Video
	result = m.DB.Where("ID IN (?)", storeVideoID).Find(&storeVideoList)
	if result != nil {
		fmt.Println("Make VideoList Error:", err)
		return
	}
	c.JSON(http.StatusOK, m.FavoriteListResponse{
		StatusCode: 0,
		VideoList:  storeVideoList,
	})
}
