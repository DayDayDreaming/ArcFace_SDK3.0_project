package models

import (
	"localhost/database"
)

type Employee struct {
	Id           int
	Name         string
	Authority    int
	Status       int
	Facelocation string
	Createtime   int64
}

func InsertEmployee(empolyee Employee) (int64, error) {
	return database.ModifyDB("insert into employees(name, authority, status, facelocation, createtime) values (?,?,?,?,?)",
		empolyee.Name, empolyee.Authority, empolyee.Status, empolyee.Facelocation, empolyee.Createtime)
}

func Getemployee() ([]Employee, error) {
	rows, err := database.QueryDB("SELECT id, name, authority, status FROM employees")
	if err != nil {
		return nil, err
	}

	var employees []Employee

	for rows.Next() {
		var employee Employee
		//遍历表中所有行的信息
			rows.Scan(&employee.Id, &employee.Name, &employee.Authority, &employee.Status)
		//将employee添加到employees中
		employees = append(employees, employee)
	}
	return employees, nil
}
