package database

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitMysql()  {
	fmt.Println("Initial......")
	if db == nil {
		db, _ = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/hotelcontrol?charset=utf8")
		CreateTableWithCanteen()
		CreateTableWithEmployee()
		CreateTableWithGym()
		CreateTableWithPool()
		CreateTableWithResident()
		CreatTableWithLend()
		CreateTableWithDelEmployee()
		ModifyDB("create trigger new_residents after insert on residents for each row " +
			" insert ignore into " +
			"lend(roomid,umbrella,charge,mobilepower,createtime) values(new.roomid,0,0,0,0) " )
		ModifyDB("create trigger del_emp after delete on employees for each row insert into del_employees(id,name,authority,status,facelocation,createtime) values(old.id,old.name,old.authority,old.status,old.facelocation,old.createtime);")
	}
}

func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}



//创建住户信息:顺序， 用户姓名，房间号，身份证号， 面部特征位置， 创建时间
func CreateTableWithResident()  {
	sql := `CREATE TABLE IF NOT EXISTS residents(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		roomid VARCHAR(6),
		idcard VARCHAR(18),
		facelocation VARCHAR(100),
		createtime INT(10)
		);`

	ModifyDB(sql)
}


//游泳池：用户姓名， 房间号， 创建时间
func CreateTableWithPool()  {
	sql := `CREATE TABLE IF NOT EXISTS pools(
		username VARCHAR(64),
		roomid VARCHAR(6),
		facelocation VARCHAR(100)
		);`

	ModifyDB(sql)
}



//健身房：用户姓名， 房间号， 创建时间
func CreateTableWithGym()  {
	sql := `CREATE TABLE IF NOT EXISTS gyms(
		username VARCHAR(64),
		roomid VARCHAR(6),
		facelocation VARCHAR(100)
		);`

	ModifyDB(sql)
}



//餐厅：用户姓名， 房间号， 创建时间
func CreateTableWithCanteen()  {
	sql := `CREATE TABLE IF NOT EXISTS canteens(
		username VARCHAR(64),
		roomid VARCHAR(6),
		facelocation VARCHAR(100)
		);`

	ModifyDB(sql)
}


//创建员工信息:id， 员工姓名，权限，状态（0休假， 1上班， 2请假），面部特征位置， 创建时间
func CreateTableWithEmployee()  {
	sql := `CREATE TABLE IF NOT EXISTS employees(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		name VARCHAR(64),
		authority INT(2),
		status INT(2),
		facelocation VARCHAR(100),
		createtime INT(10)
		);`

	ModifyDB(sql)
}

//创建借物品
func CreatTableWithLend()  {
	sql := `CREATE TABLE IF NOT EXISTS lend(
		roomid INT(6) PRIMARY KEY AUTO_INCREMENT NOT NULL ,
		umbrella INT(1),
		charge INT(1),
		mobilepower INT(1),
		createtime INT(10)
		);`

	ModifyDB(sql)
}

func CreateTableWithDelEmployee()  {
	sql := `CREATE TABLE IF NOT EXISTS del_employees(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		name VARCHAR(64),
		authority INT(2),
		status INT(2),
		facelocation VARCHAR(100),
		createtime INT(10)
		);`

	ModifyDB(sql)
}




func QueryDB(sql string) (*sql.Rows, error) {
	return db.Query(sql)
}



func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}