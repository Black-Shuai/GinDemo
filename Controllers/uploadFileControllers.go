package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)
//上传单个文件
func Upload(ctx *gin.Context) {
	//获取文件头
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.String(http.StatusBadRequest, "请求失败")
		return
	}
	//获取文件名
	fileName := file.Filename
	var fName=time.Now().Format("2006-01")//获取当前时间以便于文件创建
	_,ferr := os.Stat("file/"+fName)//检查文件夹是否存在
	if ferr == nil {
	}
	if ferr!=nil {//文件夹不存在则创建新的文件夹
		os.MkdirAll("file/"+fName, os.ModePerm)
	}
	//保存文件到服务器本地
	//存储与项目文件夹下的file文件夹下
	filePath :="file/"+fName+"/"+fileName
	//SaveUploadedFile(文件头，保存路径)
	xerr := ctx.SaveUploadedFile(file, filePath);
	if  xerr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code":0,
			"Message":"上传文件失败",
			"error":xerr.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Code":1,
		"Message":"上传文件成功",
		"urlPath":"/api/file/getimage?imageName="+fName+"/"+fileName,
	})
}
//上传多个文件
func MultipartUpload(c *gin.Context)  {
	// 多文件
	form, _ := c.MultipartForm()
	var fileList[] string
	files := form.File["MultipartFile"]
	var fName=time.Now().Format("2006-01")//获取当前时间以便于文件创建
	_,ferr := os.Stat("file/"+fName)//检查文件夹是否存在
	if ferr == nil {
	}
	if ferr!=nil {//文件夹不存在则创建新的文件夹
		os.MkdirAll("file/"+fName, os.ModePerm)
	}
	for _, file := range files {
		log.Println(file.Filename)

		dst :="file/"+fName+"/"+file.Filename
		// 上传文件到指定的路径
		c.SaveUploadedFile(file, dst)
		fileList=append(fileList,"/getimage?imageName="+fName+"/"+file.Filename )
	}
	c.JSON(http.StatusOK, gin.H{
		"Code":1,
		"Message":"上传文件成功",
		"urlPath":fileList,
	})
}
func GetImage(c *gin.Context){
	imageName := c.Query("imageName")
	fmt.Println(imageName)
	path:="file/"+imageName
	file, _ := ioutil.ReadFile(path)
	c.Writer.WriteString(string(file))
}