package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserInfo struct {
	ID uint
	Name string
	Gender string
	Hobby string
}


func main() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/test_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&UserInfo{})
    var usin UserInfo
	db.Where("name = ?", "北京").Delete(&usin)
	// DELETE from UserInfo where name = "北京";

	// DELETE FROM users WHERE id = 10;


	db.Delete(&UserInfo{}, 10)
	// DELETE FROM users WHERE id = 10;
	//等价类似
	db.Delete(&UserInfo{}, "10")
	// DELETE FROM users WHERE id = 10;

	db.Delete(&UserInfo{}, []int{1,2,3})
	// DELETE FROM users WHERE id IN (1,2,3);
	//Batch Delete
	db.Where("name LIKE ?", "%北京%").Delete(&UserInfo{})
	// DELETE from emails where email LIKE "%jinzhu%";

}
