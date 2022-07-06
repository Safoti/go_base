package main

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
 使用原始sql
 */

type Result struct {
	ID   int
	Name string
	Age  int
}
func main() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/test_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var result Result
	db.Raw("SELECT id, name, age FROM users WHERE name = ?", 3).Scan(&result)


	//row := db.Table("users").Where("name = ?", "jinzhu").Select("name", "age").Row()
	//row.Scan(&name, &age)


	//rows, err := db.Raw("select name, age, email from users where name = ?", "jinzhu").Rows()
	//defer rows.Close()
	//for rows.Next() {
	//	rows.Scan(&name, &age, &email)
	//
	//	// do something
	//}
}
