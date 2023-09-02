package controller

import (
	l "WheelChair-tiktok/logger"
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
	username, _ := c.Get("username")
	userIDs := c.Query("user_id")
	userID, err := strconv.Atoi(userIDs)
	if err != nil {
		l.Logger.Errorf("User '%s' get favorite list failed,because %s. Client IP:%s", username, err.Error(), c.ClientIP())
		favoriteListRespErr(c, err.Error())
		return
	}
	//获取user的点赞的视频的列表
	videos, err := service.GetFavoriteVideoList(uint(userID))
	if err != nil {
		l.Logger.Errorf("User '%s' get favorite list failed,because %s. Client IP:%s", username, err.Error(), c.ClientIP())
		favoriteListRespErr(c, err.Error())
		return
	}
	//把videos转换为resp.videos
	var respVideos []resp.Video
	for _, video := range videos {
		authorInfo, _ := service.GetUserInfo(video.AuthorID)
		//构建视频响应列表
		//todo 优化
		respVideos = append(respVideos, video.ToResponse(true, authorInfo.ToResponse(true)))
	}
	l.Logger.Infof("User '%s' get favorite list success. Client IP:%s", username, c.ClientIP())
	c.JSON(http.StatusOK, resp.FavoriteList{
		StatusCode: 0,
		StatusMsg:  "success",
		VideoList:  respVideos,
	})
	return

}

func favoriteActionRespErr(c *gin.Context, err string) {
	c.JSON(http.StatusOK, resp.FavoriteAction{
		StatusCode: 1,
		StatusMsg:  err,
	})
	return
}

func favoriteListRespErr(c *gin.Context, err string) {
	c.JSON(http.StatusOK, resp.FavoriteList{
		StatusCode: 1,
		StatusMsg:  err,
		VideoList:  nil,
	})
	return
}
