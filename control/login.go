package control

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"reminder/conf"
	"reminder/internal/jwt"
	"reminder/model"
	"reminder/utils"
	"time"
)

func Login(c *gin.Context)  {

	username := c.PostForm("username")
	password := c.PostForm("password")

	if len(username) <= 0 || len(password) <=0 {
		c.JSON(utils.Fail("参数有误"))
		return
	}
	// 获取账号情况与密码
	data, has := model.DoLogin(username)
	if !has || data.IsDelete == 1{
		c.JSON(utils.Fail("账号不存在"))
		return
	}
	// 校验密码
	pwd := getPwd(password,data.Salt)
	//log.Println(pwd)
	//return
	if pwd != data.Password {
		c.JSON(utils.Fail("密码错误"))
		return
	}
	// 生成token
	auth := jwt.JwtAuth{
		Id:    data.Id,
		Role:  model.Role_A,
		Name:  username,
		ExpAt: time.Now().Add(time.Hour * 2).Unix(), // 估期时间 2 小时
	}
	token := auth.Encode(conf.App.Jwt.Jwtkey)
	// 更新用户登录的信息
	c.JSON(utils.Succ("登录成功",token))
}

func SignOut(c *gin.Context)  {
	// todo 借助redis 做记录
}

func getPwd(pwd ,salt string) string  {
	return md5V3(pwd,salt)
}

func md5V3(str,salt string) string {
	w := md5.New()
	pwd := str + "<=" + salt
	//log.Println(pwd)
	io.WriteString(w, pwd)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}