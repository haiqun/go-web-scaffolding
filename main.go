package main

import (
	"log"
	"os"
	"reminder/conf"
	"reminder/model"
	"reminder/router"
)


func main() {
	log.Println("app initializing")
	// 加载配置环境
	conf.Init()
	// db 初始化
	model.Init()
	// 载入扩展
	quit := make(chan os.Signal)
	// 启动http服务
	log.Println("app running")
	go router.RunApp()
	<-quit
	// 服务退出
	log.Println("app quitted")
}
