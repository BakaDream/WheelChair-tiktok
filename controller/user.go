package controller

import (
	g "WheelChair-tiktok/global"
	m "WheelChair-tiktok/model"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	ip := c.ClientIP()
	var user m.User
	result := g.DB.Where("UserName = ?", username).Find(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 没有找到记录
			hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				log.Fatal("Hash error\n")
			}
			user = m.User{Password: string(hash), UserName: username, IP: ip}
			g.DB.Create(&user)
		} else {
			// 查询过程中发生了其他错误
			log.Fatal("Unknown error")
		}
	} else {
		log.Println("The user already exists")
		c.JSON(http.StatusOK, m.Response{StatusCode: 1, StatusMsg: "The user " + username + "already exists"})
		return
	}
	//返回 UserRegisterResponse
	//token
	//c.JSON(http.StatusOK, UserLoginResponse{
	//	Response: model.Response{StatusCode: 0},
	//	UserId:   user.ID,
	//	Token:    token,
	//})
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	var user m.User
	result := g.DB.Where("UserName = ? AND Password = ?", username, password).Find(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 没有找到记录
			log.Println("The username or password is incorrect")
			c.JSON(http.StatusOK, m.Response{StatusCode: 1, StatusMsg: "The username or password is incorrect"})
			return
		} else {
			// 查询过程中发生了其他错误
			log.Fatal("Unknown error\n")
		}
	} else {
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal("Hash error\n")
		}
		if string(hash) == user.Password {
			log.Println("Login successful")
		}
	}
	//返回 UserRegisterResponse
	//token
	//c.JSON(http.StatusOK, UserLoginResponse{
	//	Response: model.Response{StatusCode: 0},
	//	UserId:   user.ID,
	//	Token:    token,
	//})
}

// 检验身份信息token
func UserInfo(c *gin.Context) {

}
