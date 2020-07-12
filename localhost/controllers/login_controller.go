package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
)

func LoginGet(c *gin.Context)  {
	c.HTML(http.StatusOK, "login.html", gin.H{"title": "登录页"})
}


func LoginPost(c *gin.Context)  {
	username := c.PostForm("username")
	password := c.PostForm("password")
	//repassword := c.PostForm("repassword")
	fmt.Println(username, password)

	cmd := exec.Command("python", "C:/Users/Assassin/Desktop/ArcFace-python-master/staff_check.py")
	//out, _ := cmd.CombinedOutput()

	//if err := cmd.Run() ; err != nil {
	//	if exitError, ok := err.(*exec.ExitError); ok {
	//		if exitError.ExitCode() != 1 {
	//			println(exitError.ExitCode())
	//			c.Redirect(http.StatusMovedPermanently,"/login")
	//		}
	//	}
	//}


	err := cmd.Run()
	exitError, _ := err.(*exec.ExitError)







	if username == "admin" && password == "123456" && exitError.ExitCode() == 1 {

		session := sessions.Default(c)
		session.Set("loginuser", username)
		session.Save()

		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "登录成功"})
	} else {
		println(exitError.ExitCode())
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "登录失败"})
	}
}