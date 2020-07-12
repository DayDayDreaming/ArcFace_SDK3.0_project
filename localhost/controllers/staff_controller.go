package controllers

import (

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"localhost/database"
	"log"
	"net/http"
	"os"
)


var facelocation string

func StaffGet(c *gin.Context)  {

	c.HTML(http.StatusOK, "delete.html", gin.H{})

}

func StaffPost(c *gin.Context)  {


	islogin := GetSession(c)
	if islogin == false {
		c.JSON(http.StatusOK, gin.H{"Islogin":islogin, "code": 0, "message": "注册失败"})
	} else {

		name := c.PostForm("name")
		println(name)

		rows, _ := database.QueryDB("SELECT facelocation from employees where name = '" + name + "'")

		defer rows.Close()
		for rows.Next() {
			if err := rows.Scan(&facelocation); err != nil {
				log.Fatal(err)
			}

		}
		println(facelocation)
		os.RemoveAll(facelocation)

		sql := "delete from employees where name = '" + name + "'"
		database.ModifyDB(sql)

		c.JSON(http.StatusOK, gin.H{"Islogin":islogin,"code": 1, "message": "删除成功"})
	}

}
