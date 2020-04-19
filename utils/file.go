package utils

import (
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"os"
	"reminder/conf"
	"strings"
	"time"
)

func UpdateImgageFile(c *gin.Context,file *multipart.FileHeader)(retFileName string){
	// 按日期保存数据
	today := time.Now().Format("20060102") + "/"
	// 判断文件夹是否存在
	dir := "." + conf.App.Static.StaticImgBase + today;
	exi, _ := PathExists(dir)
	if !exi {
		// 创建文件夹
		CreateDir(dir)
	}
	// 保存图片
	retFileName = dir +file.Filename
	c.SaveUploadedFile(file, retFileName)
	return strings.Trim(retFileName,".")
}


// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// _dir 目录名称
func CreateDir(_dir string)  {
	log.Println("CreateDir",_dir)
	// 创建文件夹
	err := os.Mkdir(_dir, os.ModePerm)
	if err != nil {
		log.Fatal("mkdir failed![%v]\n", err)
	}
}