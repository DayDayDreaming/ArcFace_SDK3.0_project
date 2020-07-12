package models

import "localhost/database"

func DeleteGuest(roomid string)  {
	sql := "delete from residents where roomid = " + roomid
	database.ModifyDB(sql)
}

func DeleteLend(roomid string)  {
	sql := "delete from lend where roomid = " + roomid
	database.ModifyDB(sql)
}

