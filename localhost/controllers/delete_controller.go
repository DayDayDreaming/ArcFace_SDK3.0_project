package controllers

import (
	"github.com/gin-gonic/gin"
	"localhost/models"
	"net/http"
	"os"
)

func DeleteGet(c *gin.Context)  {

	c.HTML(http.StatusOK, "remove.html", gin.H{})

}

func DeletePost(c *gin.Context)  {

	islogin := GetSession(c)
	if islogin == false {
		c.JSON(http.StatusOK, gin.H{"Islogin":islogin, "code": 0, "message": "失败"})
	} else {

		roomid := c.PostForm("roomid")
		path := "C:/Users/Assassin/Desktop/localhost/static/image/guest/" + roomid
		println(path)
		os.RemoveAll(path)

		models.DeleteGuest(roomid)
		models.DeleteLend(roomid)

		c.JSON(http.StatusOK, gin.H{"Islogin":islogin,"code": 1, "message": "删除成功"})
	}
}
