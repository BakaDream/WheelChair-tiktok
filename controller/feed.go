package controller

import (
	l "WheelChair-tiktok/logger"
	resp "WheelChair-tiktok/model/response"
	"WheelChair-tiktok/service"
	"WheelChair-tiktok/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Feed(c *gin.Context) {
	uid, _ := c.Get("uid")
	lastTime := c.Query("latest_time")
	//解析时间戳
	parseLateTime, _ := utils.StringToTime(lastTime)

	//如果为空，则设置为最新一个视频的时间戳
	if parseLateTime == (time.Time{}) {
		var err error
		parseLateTime, err = service.GetLastVideoTime()
		//错误处理
		if err != nil {
			l.Logger.Errorf("Get feed failed,Because %s,Client IP %s", err.Error(), c.ClientIP())
			feedRespErr(c, "latest_time is error")
			return
		}
	}
	// 获取视频列表
	videos, err := service.GetVideoList(parseLateTime)
	if err != nil {
		l.Logger.Errorf("Get feed failed,Because %s,Client IP %s", err.Error(), c.ClientIP())
		feedRespErr(c, "get feed failed ,please retry it")
		return
	}
	//把视频列表转换为resp的视频列表
	var respVideos []resp.Video
	for _, video := range videos {
		authorInfo, _ := service.GetUserInfo(video.AuthorID)
		//构建视频响应列表
		//todo 优化
		respVideos = append(respVideos, video.ToResponse(service.IsFavorite(uid.(uint), video.ID), authorInfo.ToResponse(service.IsFollowing(uid.(uint), video.AuthorID))))
	}
	lastIndex := len(videos) - 1
	lastElement := videos[lastIndex]
	nextTime := lastElement.CreatedAt
	c.JSON(http.StatusOK, resp.Feed{
		StatusCode: 0,
		StatusMsg:  "successful",
		VideoList:  respVideos,
		NextTime:   nextTime.Unix(),
	})
}

// 响应错误
func feedRespErr(c *gin.Context, err string) {
	c.JSON(http.StatusOK, resp.Feed{
		StatusCode: 1,
		StatusMsg:  err,
		VideoList:  nil,
		NextTime:   0,
	})
}
