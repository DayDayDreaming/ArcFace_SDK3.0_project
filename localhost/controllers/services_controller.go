package controllers

import (
	"github.com/gin-gonic/gin"
	"localhost/database"
	"net/http"
)

func ServiceGet(c *gin.Context)  {

	c.HTML(http.StatusOK, "services.html", gin.H{})

}

func ServicePost(c *gin.Context)  {
	roomid := c.PostForm("roomid")
	umbrellanum := c.PostForm("umbrella")
	chargenum := c.PostForm("charge")
	mobilepowernum := c.PostForm("mobilepower")
	println(roomid, umbrellanum,chargenum,mobilepowernum)

	database.ModifyDB("update lend set umbrella = umbrella + '"+umbrellanum+"', charge = charge + '"+chargenum+"', " +
		"mobilepower = mobilepower + '"+mobilepowernum+"' where roomid = '"+roomid+"' ")

	c.Redirect(http.StatusMovedPermanently,"/")
}
