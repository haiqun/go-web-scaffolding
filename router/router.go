package router

import (
	"github.com/gin-gonic/gin"
)

// 路由入口 -  载入全局路由
func RunApp(){
	r := gin.Default()
	// 多级目录都可以访问
	//r.StaticFS("/static", http.Dir("static"))
	// 静态路由
	r.Static("/static", "./static")
	// 前台的路由配置
	apiRouter(r)
	// 需要登陆才能访问

	// 后台的路由配置
	r.Run(":8088")
}


