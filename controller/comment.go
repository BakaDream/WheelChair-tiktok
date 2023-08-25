package controller

import (
	g "WheelChair-tiktok/global"
	m "WheelChair-tiktok/model"
	u "WheelChair-tiktok/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType, err := strconv.Atoi(c.Query("action_type"))
	if err != nil {
		log.Println("The type parameter is incorrect")
		return
	}
	VideoID, err := strconv.Atoi(c.Query("video_id"))
	if err != nil {
		log.Println("The type of VideoID is incorrect")
		return
	}
	UserID, err := u.GetUserID(token)
	if err != nil {
		log.Println("User don't exit")
		return
	}
	//var user m.Comment
	if actionType == 1 {
		text := c.Query("comment_text")
		storeComment := m.Comment{UserID: UserID, VideoID: uint(VideoID), Content: text}
		err := g.DB.Create(&storeComment)
		if err.Error != nil {
			log.Fatal("Comment Upload failed")
		}
		err = g.DB.Model(&m.Video{}).Where("ID = ?", VideoID).Update("CommentCount", gorm.Expr("CommentCount + ?", 1))
		if err != nil {
			log.Fatal("Failed to update video comment count")
		}
		c.JSON(http.StatusOK, m.CommentActionResponse{
			StatusCode: 0,
			Comment:    storeComment,
		})
	} else if actionType == 2 {
		commentID := c.Query("comment_id")
		result := g.DB.Delete(&m.Comment{}, commentID)
		if result.Error != nil {
			log.Fatal("Comment delete failed")
		}
		err := g.DB.Model(&m.Video{}).Where("ID = ?", VideoID).Update("CommentCount", gorm.Expr("CommentCount + ?", -1))
		if err != nil {
			log.Fatal("Failed to update video comment count")
		}
		c.JSON(http.StatusOK, m.Response{StatusCode: 0})
	}
}
func CommentList(c *gin.Context) {
	videoID, err := strconv.Atoi(c.Query("video_id"))
	if err != nil {
		log.Println("The type of VideoID is incorrect")
		return
	}
	var Comments []m.Comment
	result := g.DB.Where("VideoID = ?", videoID).Order("CreateAt DESC").Find(&Comments)
	if result != nil {
		fmt.Println("Make CommentList Error:", err)
		return
	}
	c.JSON(http.StatusOK, m.CommentListResponse{
		StatusCode:  0,
		CommentList: Comments,
	})
}
