package model

import (
	"time"
)

type T_posts struct {
	Id              int       `xorm:"not null pk autoincr INT(11)" json:"id"`
	UserId          int       `xorm:"not null INT(11)" json:"user_id"`
	TagName         string    `xorm:"not null VARCHAR(128)" json:"tag_name"`
	CoverImg        string    `xorm:"not null VARCHAR(255)" json:"cover_img"`
	ReminderTime    time.Time `xorm:"default 'NULL' index TIMESTAMP" json:"reminder_time"`
	CreateTime   	time.Time `xorm:"created" json:"create_time" description:"添加时间"`
	UpdateTime 		time.Time `xorm:"updated" json:"update_time" description:"更新时间"`
	IsDelete        int      `xorm:"not null default 0 comment('1 为删除，0 为不删除') TINYINT(3)" json:"is_delete"`
}

var is_Del_y = 1
var is_Del_n = 0

// 录入一条信息
func PostAdd(mod *T_posts) bool {
	sess := Db.NewSession()
	defer sess.Close()
	sess.Begin()
	affect, _ := sess.InsertOne(mod)
	if affect != 1 {
		sess.Rollback()
		return false
	}
	sess.Commit()
	return true
}

// 查询数据
//PostGet 一个
func PostGet(id int) (*T_posts, bool) {
	mod := &T_posts{
		Id: id,
		IsDelete:is_Del_n,
	}
	has, _ := Db.Get(mod)
	return mod, has
}


