package models

import "localhost/database"



func LiveStatistics() (int, int, int) {
	row1 := database.QueryRowDB("select count(*) from gyms")
	gym := 0
	row1.Scan(&gym)
	row2 := database.QueryRowDB("select count(*) from canteens")
	canteen := 0
	row2.Scan(&canteen)
	row3 := database.QueryRowDB("select count(*) from pools")
	pool := 0
	row3.Scan(&pool)
	return gym, pool, canteen
}
