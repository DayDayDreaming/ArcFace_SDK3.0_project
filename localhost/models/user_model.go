package models

import "localhost/database"

type User struct {
	Id int
	Username string
	Roomid string
	Idcard string
	Facelocation string
	Createtime int64
}

func InsertUser(user User) (int64, error) {
	return database.ModifyDB("insert into residents(username, roomid, idcard, facelocation, createtime) values (?,?,?,?,?)",
		user.Username, user.Roomid, user.Idcard, user.Facelocation, user.Createtime)
}




