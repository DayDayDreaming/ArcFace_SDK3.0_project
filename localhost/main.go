package main

import (
	"localhost/database"
	"localhost/routers"
)

func main()  {
	database.InitMysql()
	router := routers.InitRouter()
	router.Static("/static", "./static")
	router.Run(":8084")
}


