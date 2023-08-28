package controller

import (
	"WheelChair-tiktok/cache"
	l "WheelChair-tiktok/logger"
	m "WheelChair-tiktok/model"
	resp "WheelChair-tiktok/model/response"
	"WheelChair-tiktok/service"
	"WheelChair-tiktok/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Register(c *gin.Context) {
	var u m.User
	u.UserName = c.Query("username")
	u.Password = c.Query("password")
	// todo: 参数校验
	u.IP = c.ClientIP()
	// 检测密码是否合法
	err := utils.IsPasswordValid(u.Password)
	// 密码不合法
	if err != nil {
		l.Logger.Infof("Registration attempt failed for user '%s' due to an invalid password. IP: %s", u.UserName, c.ClientIP())
		c.JSON(http.StatusOK, resp.Register{
			StatusCode: 1,
			StatusMsg:  "Password is invalid",
			UserId:     0,
			Token:      "",
		})
		return
	}
	if err == nil {
		//调用添加用户服务
		nUser, err := service.AddUser(u)
		if err == nil {
			token, err := utils.GenerateToken(nUser.ID, nUser.UserName)
			if err == nil {
				cache.SetToken(nUser.ID, token)
				// 全部流程走成功 构建响应 log
				c.JSON(http.StatusOK, resp.Register{
					StatusCode: 0,
					StatusMsg:  "successful",
					UserId:     int64(nUser.ID),
					Token:      token,
				})
				return
			}
		}
	}
	// 执行到这 err != nil 开始处理错误
	if err != nil {

		// 用户已注册
		if err.Error() == "user has already registered" {
			l.Logger.Infof("Registration attempt failed for user '%s' because the user is already registered. IP: %s", u.UserName, c.ClientIP())
			c.JSON(http.StatusOK, resp.Register{
				StatusCode: 0,
				StatusMsg:  "User has already registered",
				UserId:     0,
				Token:      "",
			})
			return
		}

		// 剩余一些 未知err 一般是service 执行err log error
		l.Logger.Errorf("Login attempt failed for user '%s' due to service err: %s. IP : %s", u.UserName, err.Error(), c.ClientIP())
		c.JSON(http.StatusOK, resp.Register{
			StatusCode: 0,
			StatusMsg:  "Register failed,please retry it",
			UserId:     0,
			Token:      "",
		})
	}
}

func Login(c *gin.Context) {
	var user m.User
	user.UserName = c.Query("username")
	user.Password = c.Query("password")

	//todo: 参数校验

	mUser, err := service.Login(user)
	// change: 如果err == nil 调用err.Error是危险操作
	if err == nil {
		// 生成token
		token, err := utils.GenerateToken(mUser.ID, mUser.UserName)
		if err == nil {
			// 设置token
			cache.SetToken(mUser.ID, token)
			// 全部流程走成功
			l.Logger.Infof("User '%s' has successfully logged in. IP: %s", user.UserName, c.ClientIP())
			c.JSON(http.StatusOK, resp.Login{
				StatusCode: 0,
				StatusMsg:  "successful",
				UserId:     int64(mUser.ID),
				Token:      token,
			})
			return
		}
	}

	// 走到这 err != nil 开始错误处理

	// 用户不存在
	if err.Error() == "the user is not exist" {
		l.Logger.Infof("Login attempt failed for user '%s' because the user does not exist. IP: %s", user.UserName, c.ClientIP())
		c.JSON(http.StatusOK, resp.Login{
			StatusCode: 1,
			StatusMsg:  "the user is not exist",
			UserId:     0,
			Token:      "",
		})
		return
	}
	//密码错误
	if err.Error() == "incorrect password" {
		l.Logger.Infof("Login attempt failed for user '%s' due to incorrect password. IP: %s", user.UserName, c.ClientIP())
		c.JSON(http.StatusOK, resp.Login{
			StatusCode: 0,
			StatusMsg:  "incorrect password",
			UserId:     0,
			Token:      "",
		})
		return
	}
	// 剩余一些 未知err 一般为某个函数未能正常执行 log等级为error
	l.Logger.Errorf("Login attempt failed for user '%s' due to service err: %s. IP : %s", user.UserName, err.Error(), c.ClientIP())
	c.JSON(http.StatusOK, resp.Login{
		StatusCode: 1,
		StatusMsg:  "login failed,please retry it",
		UserId:     0,
		Token:      "",
	})
	return
}

// / 获取用户信息
func UserInfo(c *gin.Context) {
	uId, err := strconv.Atoi(c.Query("user_id"))
	// uid 转换失败 返回错误
	if err != nil {
		l.Logger.Warnf("Failed to convert user_id to int: %s . IP %s", err.Error(), c.ClientIP())
		c.JSON(http.StatusOK, resp.UserInfo{
			StatusCode: 0,
			StatusMsg:  "invalid user id",
		})
		return
	}
	var user m.User
	user, err = service.GetUserInfo(uint(uId))
	if err == nil {
		l.Logger.Infof("%s info retrieved successfully for user ID %d. IP %s", user.UserName, uId, c.ClientIP())
		c.JSON(http.StatusOK, resp.UserInfo{
			StatusCode: 0,
			StatusMsg:  "successful",
			User:       user,
		})
		return
	}
	//用户不存在
	if err.Error() == "the user is not exist" {
		l.Logger.Infof("get userinfo attempt failed for id '%d' because the user does not exist. IP: %s", uId, c.ClientIP())
		c.JSON(http.StatusOK, resp.UserInfo{
			StatusCode: 0,
			StatusMsg:  "the user is not exist",
		})
		return
	}
	// 额外的错误
	l.Logger.Errorf("get userinfo attempt failed for id '%d' due to service err: %s. IP : %s", uId, err.Error(), c.ClientIP())
}
