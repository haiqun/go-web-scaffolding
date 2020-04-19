package model

import "time"

type T_user struct {
	Id              int       `xorm:"not null pk autoincr INT(11)" json:"id"`
	Nikename        string    `xorm:"not null VARCHAR(128)" json:"nikename"`
	Username        string    `xorm:"not null VARCHAR(32)" json:"username"`
	Password        string    `xorm:"not null CHAR(32)" json:"password"`
	Salt            string    `xorm:"not null CHAR(6)" json:"salt"`
	CreateTime   	time.Time `xorm:"created" json:"create_time" description:"添加时间"`
	UpdateTime 		time.Time `xorm:"updated" json:"update_time" description:"更新时间"`
	IsDelete        int      `xorm:"not null default 0 comment('1 为删除，0 为不删除') TINYINT(3)" json:"is_delete"`
}

var Role_A = 1 // "web"
var Role_B = 2 // 预留字段 "wechat"

// 登录
func DoLogin(username string) (* T_user, bool) {
	u := &T_user{}
	has , _ := Db.Where("username = ?",username).Cols("id","password","is_delete","Salt").Get(u)
	return u, has
}

// 注册
func registered()  {

}