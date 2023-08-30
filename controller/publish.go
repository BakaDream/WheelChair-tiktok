package controller

import (
	l "WheelChair-tiktok/logger"
	m "WheelChair-tiktok/model"
	resp "WheelChair-tiktok/model/response"
	"WheelChair-tiktok/service"
	"WheelChair-tiktok/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 上传视频
func Publish(c *gin.Context) {
	title := c.PostForm("title")
	file, err := c.FormFile("data")
	// 从上下文中获取用户信息
	anyUid, _ := c.Get("uid")
	username, _ := c.Get("username")
	uid, _ := anyUid.(uint)
	//检测文件是否为视频文件
	if utils.IsVideoFile(file.Filename) {
		l.Logger.Infof("User '%s' publish an unvaild video file. Client IP %s", username, c.ClientIP())
		publishRespErr(c, "It's a valid video file")
		return
	}

	//给视频一个hashName
	file.Filename = utils.HashFileName(file.Filename)
	//保存文件
	playUrl, coverUrl, err := service.UploadVideo(file)
	if err != nil {
		l.Logger.Errorf("User '%s' publish failed,because %s, client ip: %s", username, err.Error(), c.ClientIP())
		publishRespErr(c, "publish failed ,please retry it")
		return
	}
	err = service.AddPublish(playUrl, coverUrl, title, uid)
	if err != nil {
		l.Logger.Errorf("User '%s' publish failed,because %s, client ip: %s", username, err.Error(), c.ClientIP())
		//todo:添加失败后cos和数据库对应操作
		publishRespErr(c, "punlish failed,please retry it")
		return
	}
	//响应
	l.Logger.Infof("User %s has successfully publish the %s.Client IP %s", username, title, c.ClientIP())
	c.JSON(http.StatusOK, resp.Publish{
		StatusCode: 0,
		StatusMsg:  "successful",
	})
	return
}

// PublishList 展示该用户上传的所有视频
func PublishList(c *gin.Context) {
	username, _ := c.Get("username")
	userID := c.Query("user_id")
	uid, _ := c.Get("uid")
	authorID, err := strconv.Atoi(userID)
	if err != nil {
		l.Logger.Errorf("User '%s' get publish failed,because %s.Client IP %s", username, err.Error(), c.ClientIP())
		publishListRespErr(c)
		return
	}
	//获取author的视频列表
	var videos []m.Video
	videos, err = service.GetPublishList(uint(authorID))
	if err != nil {
		l.Logger.Errorf("User '%s' get publish failed,because %s.Cinent IP %s", username, err.Error(), c.ClientIP())
		publishListRespErr(c)
		return
	}
	respVideos := make([]resp.Video, len(videos))
	for i, video := range videos {
		isfavorite := service.IsFavorite(uid.(uint), video.ID)
		authorInfo, _ := service.GetUserInfo(uint(authorID))
		author := authorInfo.ToResponse(true)
		respVideos[i] = video.ToResponse(isfavorite, author)
	}
	l.Logger.Infof("User '%s' get publish list successful. Client IP %s", username, c.ClientIP())
	c.JSON(http.StatusOK, resp.PublishList{
		StatusCode: 0,
		StatusMsg:  "successful",
		VideoList:  respVideos,
	})
	return
}

func publishRespErr(c *gin.Context, err string) {
	c.JSON(http.StatusOK, resp.Publish{
		StatusCode: 1,
		StatusMsg:  err,
	})
	return
}

func publishListRespErr(c *gin.Context) {
	c.JSON(http.StatusOK, resp.PublishList{
		StatusCode: 0,
		StatusMsg:  "fail to get publish list,please retry it",
		VideoList:  nil,
	})
	return
}
