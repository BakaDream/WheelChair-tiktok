package middleware

import (
	"WheelChair-tiktok/cache"
	resp "WheelChair-tiktok/model/response"
	"WheelChair-tiktok/utils"
	"github.com/gin-gonic/gin"
)

func tokenErr(c *gin.Context, statusMsg string) {
	c.JSON(200, resp.Auth{
		StatusCode: 1,
		StatusMsg:  statusMsg,
	})
}

func Auth() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Query("token")

		//token不存在，直接返回
		if token == "" {
			tokenErr(c, "please login")
			c.Abort()
			return
		}

		// 解析token
		iClaims, err := utils.ParseToken(token)
		// token无法解析 返回
		if err != nil {
			tokenErr(c, "invalid token")
			c.Abort()
			return
		}

		// 根据ID 获取redis的token
		rdToken, err := cache.GetToken(iClaims.ID)
		if err != nil {
			tokenErr(c, "please retry login")
			c.Abort()
			return
		}

		if rdToken != token {
			tokenErr(c, "Token mismatch")
			c.Abort()
			return
		}

		// todo: token过期 续期
		c.Set("uid", iClaims.ID)
		c.Set("username", iClaims.UserName)
		c.Next()
	}
}
