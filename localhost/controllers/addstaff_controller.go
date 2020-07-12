package controllers

import (
	"github.com/gin-gonic/gin"
	"localhost/models"
	"net/http"
	"os/exec"
	"strconv"
	"time"
)

func AddemployeeGet(c *gin.Context)  {

	c.HTML(http.StatusOK, "add.html", gin.H{})

}


func Addemployee(c *gin.Context)  {

	islogin := GetSession(c)
	if islogin == false {
		c.JSON(http.StatusOK, gin.H{"Islogin":islogin, "code": 0, "message": "注册失败"})
	} else {

		name := c.PostForm("name")
		Authority := c.PostForm("authority")
		authority, _ := strconv.Atoi(Authority)
		Status := c.PostForm("status")
		status, _ := strconv.Atoi(Status)
		//facelocation := c.PostForm("facelocation")
		//file, err := c.FormFile("upload")
		//
		//if err != nil {
		//	c.JSON(http.StatusOK, gin.H{"code":0, "message":"图片输入有误"})
		//}
		rename := time.Now().Unix()
		path := "C:/Users/Assassin/Desktop/localhost/static/image/employee/" + strconv.FormatInt(rename, 10) + ".jpg"
		println(path)
		//c.SaveUploadedFile(file, path)

		com := exec.Command("python", "C:/Users/Assassin/Desktop/ArcFace-python-master/test.py", path)
		com.Run()

		employee := models.Employee{0, name, authority, status, path, time.Now().Unix()}

		_, er := models.InsertEmployee(employee)
		if er != nil {
			c.JSON(http.StatusOK, gin.H{"Islogin":islogin,"code": 0, "message": "注册失败"})
		} else {
			c.JSON(http.StatusOK, gin.H{"Islogin":islogin,"code": 1, "message": "注册成功"})
		}
	}
}