package router

import (
	"github.com/gin-gonic/gin"
	"reminder/control"
)

// apiRouter
func apiRouter(api *gin.Engine) {
	// 首页入口
	api.GET("/", control.IndexView)
	// 登录
	api.POST("/login",control.Login)
	// 登出
	api.POST("/signout",control.SignOut)
	// 查看帖子
	api.GET("/info/:id",ApiAuth(),control.GetInfo)
	// 录入帖子
	api.POST("/info",ApiAuth(),control.AddInfo)
	// 修改帖子

	// 删除帖子

	// 帖子列表
	api.GET("/list",ApiAuth(),control.AddInfo)
}

