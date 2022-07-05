package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

//简单的crud demo

//type User struct {
//	ID           uint
//	Name         string
//	Email        *string
//	Age          uint8
//	Birthday     time.Time
//	MemberNumber sql.NullString
//	ActivatedAt  sql.NullTime
//	CreatedAt    time.Time
//	UpdatedAt    time.Time
//}

func main() {
	//dsn := "root:1234@tcp(127.0.0.1:3306)/test_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	panic(err)
	//}
	//// Migrate the schema
	//db.AutoMigrate(&User{})

	//use:=User{Name:"lktbz",Age: 18,Birthday: time.Now()}
	//result  := db.Create(&use)
	//db.Select("Name","Age","CreatedAt").Create(&user)
	//db.Omit("Name","Age","CreatedAt").Create(&user) 忽略某个字段
	// fmt.Println(result)

	//批量插入
	//var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	//db.Create(&users)
	//for _, user := range users {
	//	fmt.Println(user.ID)
	//}
	//批量插入方式二
	//var users = []User{{Name: "jinzhu_1"}, ...., {Name: "jinzhu_10000"}}
	//// batch size 100
	//db.CreateInBatches(users, 100)
	//Create From Map

	//db.Model(&User{}).Create(map[string]interface{}{
	//	"Name": "jinzhu", "Age": 18,
	//})
	//
	//// batch insert from `[]map[string]interface{}{}`
	//db.Model(&User{}).Create([]map[string]interface{}{
	//	{"Name": "jinzhu_1", "Age": 18},
	//	{"Name": "jinzhu_2", "Age": 20},
	//})



}
