package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"localhost/database"
	"localhost/models"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func RegisterGet(c *gin.Context)  {

	c.HTML(http.StatusOK, "register.html", gin.H{"title":"注册房客"})

}



//需要upload保存图片
func RegisterPost(c *gin.Context)  {
	username := c.PostForm("username")
	roomid := c.PostForm("roomid")
	idcard := c.PostForm("idcard")

	islogin := GetSession(c)
	if islogin == false {
		c.JSON(http.StatusOK, gin.H{"Islogin":islogin, "code": 0, "message": "注册失败"})
	} else  {
		name := time.Now().Unix()
		dirpath := ".." + roomid
		path := "C:/Users/Assassin/Desktop/localhost/static/image/guest/" + roomid + "/" + strconv.FormatInt(name, 10) + ".jpg"
		//println(path)

		exist, err := models.PathExists(dirpath)
		if err != nil {
			fmt.Printf("get dir error![%v]\n", err)
			return
		}
		_dir := "C:/Users/Assassin/Desktop/localhost/static/image/guest/" + roomid
		if exist {
			fmt.Printf("has dir![%v]\n", _dir)
		} else {
			fmt.Printf("no dir![%v]\n", _dir)
			// 创建文件夹
			err := os.Mkdir(_dir, os.ModePerm)
			if err != nil {
				fmt.Printf("mkdir failed![%v]\n", err)
			} else {
				fmt.Printf("mkdir success!\n")
			}
		}





		com := exec.Command("python", "C:/Users/Assassin/Desktop/ArcFace-python-master/test.py", path)
		com.Run()



		database.ModifyDB("insert into lend(roomid, umbrella, charge, mobilepower, createtime) values (?,?,?,?,?)",
			roomid, 0, 0, 0, time.Now().Unix())

		user := models.User{0, username, roomid, idcard, path, time.Now().Unix()}

		_, er := models.InsertUser(user)
		if er != nil {
			c.JSON(http.StatusOK, gin.H{"Islogin":islogin,"code": 0, "message": "注册失败"})
		} else {
			c.JSON(http.StatusOK, gin.H{"Islogin":islogin, "code": 1, "message": "注册成功"})
		}
	}


}
