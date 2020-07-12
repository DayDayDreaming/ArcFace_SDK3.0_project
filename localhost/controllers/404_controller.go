package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorGet(c *gin.Context)  {

	c.HTML(http.StatusOK, "404.html", gin.H{})
}
