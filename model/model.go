package model

import (
	"log"
	"reminder/conf"
	"xorm.io/xorm"
	"xorm.io/xorm/caches"
	// 数据库驱动
	_ "github.com/go-sql-driver/mysql"
)

// Db 数据库操作句柄
var Db *xorm.Engine

// 扩容

func Init() {
	// 初始化数据库操作的 Xorm
	db, err := xorm.NewEngine("mysql", conf.App.Dsn())
	if err != nil {
		log.Fatalln("数据库 dsn:", err.Error())
	}
	if err = db.Ping(); err != nil {
		log.Fatalln("数据库 ping:", err.Error())
	}
	db.SetMaxIdleConns(conf.App.Xorm.Idle)
	db.SetMaxOpenConns(conf.App.Xorm.Open)
	// 是否显示sql执行的语句
	db.ShowSQL(conf.App.Xorm.Show)
	if conf.App.Xorm.CacheEnable {
		// 设置xorm缓存
		cacher := caches.NewLRUCacher(caches.NewMemoryStore(), conf.App.Xorm.CacheCount)
		db.SetDefaultCacher(cacher)
	}
	//if conf.App.Xorm.Sync {
	//	err := db.Sync2(new(User), new(Cate), new(Tag), new(Post), new(PostTag), new(Opts))
	//	if err != nil {
	//		log.Fatalln("数据库 sync:", err.Error())
	//	}
	//}
	Db = db
	//缓存
	initMap()
}