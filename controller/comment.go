package controller

import (
	l "WheelChair-tiktok/logger"
	resp "WheelChair-tiktok/model/response"
	"WheelChair-tiktok/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CommentAction(c *gin.Context) {
	// 获取查询字符串参数
	videoIDStr := c.Query("video_id")
	actionType := c.Query("action_type")
	commentText := c.Query("comment_text")
	commentIDStr := c.Query("comment_id")
	userName, _ := c.Get("username")
	userID, _ := c.Get("uid")

	// 检查必选参数是否存在
	if videoIDStr == "" || actionType == "" {
		l.Logger.Infof("user '%s' operate comment failed, because params is nil.Client IP:%s", userName.(string), c.ClientIP())
		commentActionErr(c, "operate")
		return
	}

	// 转换字符串参数为uint类型
	videoID, err := strconv.ParseUint(videoIDStr, 10, 64)
	if err != nil {
		l.Logger.Infof("user '%s' publisg comment failed, because videoID is invalid.Client IP:%s", userName.(string), c.ClientIP())
		commentActionErr(c, "publish")
		return
	}

	// 处理点赞/取消点赞逻辑

	//点赞
	if actionType == "1" {
		//判断评论内容是否为空
		if commentText == "" {
			l.Logger.Infof("user '%s' publish comment failed, because commentText is nil.Client IP:%s", userName.(string), c.ClientIP())
			commentActionErr(c, "publish")
			return
		}

		// 开始存储commentText
		comment, err := service.PublishComment(uint(videoID), userID.(uint), commentText)
		if err != nil {
			l.Logger.Errorf("user '%s' publish comment failed, because %s.Client IP:%s", userName.(string), err.Error(), c.ClientIP())
			commentActionErr(c, "publish")
			return
		}
		//获取评论者的信息
		user, _ := service.GetUserInfo(userID.(uint))
		commentResp := comment.ToResponse(user.ToResponse(true))

		//响应
		l.Logger.Infof("user '%s' publish comment success.Client IP:%s", userName.(string), c.ClientIP())
		c.JSON(http.StatusOK, resp.CommentAction{
			StatusCode: 0,
			StatusMsg:  "success",
			Comment:    commentResp,
		})

	} else if actionType == "2" {
		// 检查comment_id不为空
		if commentIDStr == "" {
			l.Logger.Infof("user '%s' delete comment failed, because commentIDStr is null.Client IP:%s", userName.(string), c.ClientIP())
			commentActionErr(c, "delete")
			return
		}

		// 转换comment_id为uint类型
		commentID, err := strconv.ParseUint(commentIDStr, 10, 64)
		if err != nil {
			l.Logger.Infof("user '%s' delete comment failed, because commentIDStr is invalid.Client IP:%s", userName.(string), c.ClientIP())
			commentActionErr(c, "delete")
			return
		}

		// 执行删除评论操作
		err = service.DeleteComment(uint(commentID))
		if err != nil {
			l.Logger.Errorf("user '%s' delete comment failed, because %s.Client IP:%s", userName.(string), err.Error(), c.ClientIP())
			return
		}
		//成功
		l.Logger.Infof("user '%s' delete comment success.Client IP:%s", userName.(string), c.ClientIP())
		c.JSON(http.StatusOK, resp.CommentAction{
			StatusCode: 0,
			StatusMsg:  "success",
			Comment:    resp.Comment{},
		})
		return
	} else {
		commentActionErr(c, "delete")
		return
	}
}

// func CommentList(c *gin.Context) {
//
// }
func commentActionErr(c *gin.Context, action string) {
	c.JSON(http.StatusOK, resp.CommentAction{
		StatusCode: 1,
		StatusMsg:  action + "comment err,please retry it",
		Comment:    resp.Comment{},
	})
	return
}
