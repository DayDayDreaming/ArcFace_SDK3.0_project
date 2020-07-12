package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ExitGet(c *gin.Context)  {

	session := sessions.Default(c)
	session.Delete("loginuser")
	session.Save()

	//session.Clear()

	fmt.Println("delete session...",session.Get("loginuser"))
	c.Redirect(http.StatusMovedPermanently,"/")

}
