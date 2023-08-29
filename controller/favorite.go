package controller

import (
	m "WheelChair-tiktok/model"
	//g "WheelChair-tiktok/global"
	l "WheelChair-tiktok/logger"
	resp "WheelChair-tiktok/model/response"
	"WheelChair-tiktok/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func FavoriteAction(c *gin.Context) {
	token := c.Query("toekn")
	VideoID, err := strconv.Atoi(c.Query("video_id"))
	if err != nil {
		l.Logger.Infof("The type of VideoID is incorrect.IP: %s", c.ClientIP())
		return
	}
	actionType, err := strconv.Atoi(c.Query("action_type"))
	if err != nil {
		l.Logger.Infof("The type parameter is incorrect.IP: %s", c.ClientIP())
		return
	}
	UserID, err := utils.GetUserID(token)
	if err != nil {
		l.Logger.Infof("get userinfo attempt failed for id '%d' because the user does not exist. IP: %s", UserID, c.ClientIP())
		c.JSON(http.StatusOK, resp.FavoriteAction{StatusCode: 1, StatusMsg: "favorite failed"})
		return
	}
	var storeUserVideoLike m.Favorite
	result := m.DB.Where("UserID = ? and VideoID = ?", UserID, VideoID).First(&storeUserVideoLike)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) && actionType == 1 {
			// 没有找到记录
			m.DB.Create(&m.Favorite{VideoID: uint(VideoID), UserID: UserID})
			m.DB.Model(&m.Video{}).Where("ID = ?", VideoID).Update("FavoriteComment", gorm.Expr("FavoriteComment + ?", 1))
			var video m.Video
			m.DB.Where("ID = ?", VideoID).First(&video)
			m.DB.Model(&m.User{}).Where("ID = ?", video.AuthorID).Update("TotalFavorited", gorm.Expr("TotalFavorited + ?", 1))
			c.JSON(http.StatusOK, resp.FavoriteAction{StatusCode: 1, StatusMsg: "favorite successful"})
		} else {
			// 查询过程中发生了其他错误
			l.Logger.Error("Unknown error")
		}
	} else if actionType == 2 {
		m.DB.Delete(&storeUserVideoLike)
		m.DB.Model(&m.Video{}).Where("ID = ?", VideoID).Update("FavoriteComment", gorm.Expr("FavoriteComment + ?", -1))
		var video m.Video
		m.DB.Where("ID = ?", VideoID).First(&video)
		m.DB.Model(&m.User{}).Where("ID = ?", video.AuthorID).Update("TotalFavorited", gorm.Expr("TotalFavorited + ?", -1))
		c.JSON(http.StatusOK, resp.FavoriteAction{StatusCode: 1, StatusMsg: "favorite successful"})
	}
}
func FavoriteList(c *gin.Context) {
	UserID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		l.Logger.Infof("The type of UserID is incorrect")
		return
	}
	var storeFavorite []m.Favorite
	result := m.DB.Where("UserID = ?", UserID).Order("CreateAt DESC").Find(&storeFavorite)
	if result != nil {
		l.Logger.Error("Make UserVideoLikeList Error:%v", err)
		return
	}
	var storeVideoID []uint
	for _, Video := range storeFavorite {
		storeVideoID = append(storeVideoID, Video.VideoID)
	}
	var storeVideoList []m.Video
	result = m.DB.Where("ID IN (?)", storeVideoID).Find(&storeVideoList)
	if result != nil {
		l.Logger.Error("Make VideoList Error:%v", err)
		return
	}
	c.JSON(http.StatusOK, resp.FavoriteList{StatusCode: 0, StatusMsg: "Get FavoriteList successfully"})
}
