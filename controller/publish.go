package controller

import (
	g "WheelChair-tiktok/global"
	m "WheelChair-tiktok/model"
	"WheelChair-tiktok/utils"
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

func Publish(c *gin.Context) {
	tokenString := c.PostForm("token")
	uid, err := utils.GetUserID(tokenString)
	if err != nil {
		c.JSON(http.StatusOK, m.Response{
			StatusCode: 1,
			StatusMsg:  "invalid token",
		})
		return
	}

	title := c.PostForm("title")
	file, err := c.FormFile("data")
	if err != nil {
		return
	}
	file.Filename = hashFileName(file.Filename)
	err = c.SaveUploadedFile(file, "./public/"+file.Filename)
	if err != nil {
		c.JSON(http.StatusOK, m.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	_, err = utils.GetSnapshot("./public/"+file.Filename, "./public/snapshot/"+file.Filename, 1) //get the first frame of the video
	if err != nil {
		c.JSON(http.StatusOK, m.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	cover_url := g.ServerInfo.Server.StaticFileUrl + "snapshot/" + file.Filename + ".jpg"
	play_url := g.ServerInfo.Server.StaticFileUrl + file.Filename

	result := g.DB.Create(&m.Video{UserID: uid, Title: title, FavoriteCount: 0, CommentCount: 0, PlayUrl: play_url, CoverUrl: cover_url})
	if result != nil {

	}
	result = g.DB.Model(&m.User{}).Update("WorkCount", gorm.Expr("WorkCount + ?", 1))
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
	result := g.DB.Where("UserID = ?", UserID).Find(&videoList)
	if result.Error != nil {
		log.Fatal("Unknow error")
	}
	c.JSON(http.StatusOK, m.UserPublishResponse{
		StatusCode: 0,
		VideoList:  videoList,
	})
}

func hashFileName(fileName string) string {
	// 创建SHA256哈希对象
	hash := sha256.New()

	// 将文件名转换为字节数组并进行哈希计算
	hash.Write([]byte(fileName))

	// 获取哈希值并转换为十六进制字符串
	hashedBytes := hash.Sum(nil)
	hashedString := hex.EncodeToString(hashedBytes)

	return hashedString
}
