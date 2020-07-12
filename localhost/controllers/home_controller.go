package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomeGet(c *gin.Context) {
	islogin := GetSession(c)
	c.HTML(http.StatusOK, "home.html", gin.H{"Islogin":islogin})
}
