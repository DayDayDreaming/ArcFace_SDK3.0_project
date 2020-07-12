package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"localhost/controllers"
)


func InitRouter() * gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	//设置session midddleware
	store := cookie.NewStore([]byte("loginuser"))
	router.Use(sessions.Sessions("mysession", store))

	{
		router.GET("/", controllers.HomeGet)

		router.GET("/login", controllers.LoginGet)
		router.POST("/login", controllers.LoginPost)

		router.GET("/register", controllers.RegisterGet)
		router.POST("/register", controllers.RegisterPost)

		router.GET("/live", controllers.LiveGet)


		router.GET("/add",controllers.AddemployeeGet)
		router.POST("/add",controllers.Addemployee)

		//router.GET("/info", controllers.ArrangeGet)

		router.GET("/remove", controllers.DeleteGet)
		router.POST("/remove", controllers.DeletePost)

		router.GET("/delete", controllers.StaffGet)
		router.POST("/delete", controllers.StaffPost)

		router.GET("/exit", controllers.ExitGet)

		router.GET("/services", controllers.ServiceGet)
		router.POST("/services", controllers.ServicePost)

		router.GET("/error", controllers.ErrorGet)
	}



	return router

}
