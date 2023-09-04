package middleware

import (
	"WheelChair-tiktok/cache"
	"WheelChair-tiktok/utils"
	"github.com/gin-gonic/gin"
)

func FeedAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Query("token")
		//token 不存在
		if token == "" {
			// 尝试从postform中获取token
			token = c.PostForm("token")
			//token获取不到
			if token == "" {
				invalidToken(c)
				return
			}
		}

		// 解析token
		iClaims, err := utils.ParseToken(token)
		// token无法解析
		if err != nil {
			invalidToken(c)
			return
		}

		// 根据ID 获取redis的token
		rdToken, err := cache.GetToken(iClaims.ID)
		if err != nil {
			invalidToken(c)
			return
		}

		if rdToken != token {
			invalidToken(c)
			return
		}

		// todo: token过期 续期
		c.Set("uid", iClaims.ID)
		c.Set("username", iClaims.UserName)
		c.Next()
	}
}
func invalidToken(c *gin.Context) {
	c.Set("uid", uint(1))
	c.Set("username", "default")
	c.Next()
}
