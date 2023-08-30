package controller

import (
	l "WheelChair-tiktok/logger"
	m "WheelChair-tiktok/model"
	resp "WheelChair-tiktok/model/response"
	"WheelChair-tiktok/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FavoriteAction(c *gin.Context) {
	//从上下文中获取uid username
	uid, _ := c.Get("uid")
	username, _ := c.Get("username")
	//获取用户传入的videoID，Action type 并进行一定错误处理
	videoID, err := strconv.Atoi(c.Query("video_id"))
	if err != nil {
		l.Logger.Infof("User '%s' favoriteAction err,because %s. Client IP %s", username.(string), err.Error(), c.ClientIP())
		favoriteActionRespErr(c, "params video_id  invalid")
		return
	}
	actionType, err := strconv.Atoi(c.Query("action_type"))
	if err != nil {
		l.Logger.Infof("User '%s' favoriteAction err,because %s. Client IP %s", username.(string), err.Error(), c.ClientIP())
		favoriteActionRespErr(c, "params video_id  invalid")
		return
	}

	// action_type 1-点赞，2-取消点赞
	//点赞相关操作
	if actionType == 1 {
		err = service.Favorite(uint(videoID), uid.(uint))
		if err != nil {
			l.Logger.Errorf("User '%s' favoriteAction err,because %s. Client IP %s", username.(string), err.Error(), c.ClientIP())
			favoriteActionRespErr(c, err.Error())
			return
		}
		//点赞成功 响应
		l.Logger.Infof("User %s Favorite video %d success. Client IP %s", username, videoID, c.ClientIP())
		c.JSON(http.StatusOK, resp.FavoriteAction{
			StatusCode: 0,
			StatusMsg:  "successful",
		})
		return
	}

	//取消点赞
	if actionType == 2 {
		err = service.UnFavorite(uint(videoID), uid.(uint))
		if err != nil {
			l.Logger.Errorf("User '%s' UnfavoriteAction err,because %s. Client IP %s", username.(string), err.Error(), c.ClientIP())
			//todo 更加详细的错误
			favoriteActionRespErr(c, err.Error())
			return
		}
		//取消点赞成功 响应
		l.Logger.Infof("User %s UnFavorite video %d success. Client IP %s", username.(string), videoID, c.ClientIP())
		c.JSON(http.StatusOK, resp.FavoriteAction{
			StatusCode: 0,
			StatusMsg:  "successful",
		})
		return
	}
	// 未知的actionType
	l.Logger.Infof("User '%s' UnfavoriteAction %d err,because Illegal actionType. Client IP %s", username.(string), videoID, c.ClientIP())
	favoriteActionRespErr(c, "Illegal actionType")
	return
}

func FavoriteList(c *gin.Context) {
	UserID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		l.Logger.Infof("The type of UserID is incorrect")
		return
	}
	var storeFavorite []m.Favorite
	result := m.DB.Where("user_id = ?", UserID).Order("CreateAt DESC").Find(&storeFavorite)
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

func favoriteActionRespErr(c *gin.Context, err string) {
	c.JSON(http.StatusOK, resp.FavoriteAction{
		StatusCode: 1,
		StatusMsg:  err,
	})
	return
}
