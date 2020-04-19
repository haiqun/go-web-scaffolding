package router

import (
	"github.com/gin-gonic/gin"
	"reminder/conf"
	"reminder/internal/jwt"
	"reminder/utils"
)

// 登录校验
func ApiAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 权限校验 token
		auth_token := c.Request.Header.Get("auth_token")
		if auth_token == "" {
			c.JSON(utils.ErrSvr("缺少参数 auth_token"))
			c.Abort()
			return
		}
		// token
		verify, err := jwt.Verify(auth_token, conf.App.Jwt.Jwtkey)
		if err != nil {
			c.JSON(utils.ErrSvr("无效的Token"))
			c.Abort() // 不让往下走
			return
		}
		// 判断是否手动登出

		//log.Println(verify)
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("userId", verify.Id)
		c.Set("userName", verify.Name)
		c.Set("userRole", verify.Role)
		// 执行函数
		c.Next()
		// 中间件执行完后续的一些事情
		//status := c.Writer.Status()
	}
}
