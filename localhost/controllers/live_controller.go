package controllers

import (
	"github.com/gin-gonic/gin"
	"localhost/models"
	"net/http"
)

func LiveGet(c *gin.Context)  {

	gym, pool, canteen := models.LiveStatistics()
	println(gym, pool, canteen)
	c.HTML(http.StatusOK, "live.html", gin.H{"gym":gym, "pool":pool, "canteen":canteen})

}
