package controller

import (
	g "WheelChair-tiktok/global"
	m "WheelChair-tiktok/model"
	"WheelChair-tiktok/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	ip := c.ClientIP()
	var user m.User
	result := g.DB.Where("UserName = ?", username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 没有找到记录
			hash, result := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if result != nil {
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
	//生成token
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
	result := g.DB.Where("UserName = ?", username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 没有找到记录
			log.Println("The username is not exist")
			c.JSON(http.StatusOK, m.Response{StatusCode: 1, StatusMsg: "The username is not exist\n"})
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
	//返回 UserLoginResponse
	//生成token
	//c.JSON(http.StatusOK, UserLoginResponse{
	//	Response: model.Response{StatusCode: 0},
	//	UserId:   user.ID,
	//	Token:    token,
	//})
}

// 获取用户信息
func UserInfo(c *gin.Context) {
	Uid, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		log.Println("The type of VideoID is incorrect")
	}
	if !utils.CheckToken(c.Query("token")) {
		c.JSON(http.StatusOK, m.Response{
			StatusCode: 1,
			StatusMsg:  "Invalid token",
		})
		return
	}
	var user m.User
	result := g.DB.Where("ID = ?", Uid).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Println("The username is not exist")
		} else {
			log.Fatal("Unknown error\n")
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "Success",
		"user":        user,
	})
}
