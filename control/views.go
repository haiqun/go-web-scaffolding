package control

import (
	"github.com/gin-gonic/gin"
	"log"
	"reminder/model"
	"reminder/utils"
	"strconv"
	"time"
)

func IndexView(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "GET",
	})
}

type s struct {
	Name string `json:"name"`
	Age string	`json:"age"`
}

func GetInfo(c *gin.Context)  {
	// 获取uid
	uid, exists := c.Get("userId")
	if !exists {
		c.JSON(utils.Fail("请先登录"))
		return
	}
	// 获取参数
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(utils.ErrIpt(`数据输入错误,请重试`, err.Error()))
		return
	}
	// 查询数据
	mod, has := model.PostGet(id)
	if !has {
		c.JSON(utils.Fail(`数据不存在`))
		return
	}
	log.Println(mod.UserId)
	// 判断uid 是否一致
	if uid != mod.UserId {
		c.JSON(utils.Fail(`数据不存在`))
		return
	}
	c.JSON(utils.Succ("成功",mod))
}

func AddInfo(c *gin.Context) {
	// 获取参数
	file, err := c.FormFile("file")
	if err != nil {
		c.String(500, "上传图片出错")
	}
	// 用户的私有分类id
	tag_name := c.DefaultQuery("tag_name",time.Now().Format("20060102"))
	prompt_time := c.PostForm("prompt_time")
	// 获取uid
	uid, exists := c.Get("userId")

	if !exists {
		c.JSON(utils.Fail("请先登录"))
	}
	// 保存图片
	filename := utils.UpdateImgageFile(c,file)
	mod := &model.T_posts{
		UserId:       uid.(int),
		TagName:      tag_name,
		CoverImg:     filename,
		ReminderTime:     utils.Str2Time(prompt_time),
	}

	ret := model.PostAdd(mod)
	if ret {
		c.JSON(utils.Succ("录入成功"))
	}else{
		c.JSON(utils.Fail("录入失败，请重试"))
	}
}


