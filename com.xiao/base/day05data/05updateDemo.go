package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

)
//type UserInfo struct {
//	ID uint
//	Name string
//	Gender string
//	Hobby string
//}
func main() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/test_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&UserInfo{})


    var us *UserInfo
	//根据id 修改字段
	//db.Model(&UserInfo{}).Where("id=?",2).Update("name","新疆")
	//
	//db.Where("id",2).Find(&us)
	//fmt.Println(us)

	//多列更新
	//db.Model(&UserInfo{}).Where(
	//	"id=?",1).Updates(
	//	UserInfo{Name: "北京",Hobby: "政都"})
	//
	//db.Where("id",1).Find(&us)
	//fmt.Println(us)



	//使用map 方式进行多列更新

	db.Model(&UserInfo{}).Where("id",9).Updates(
		map[string]interface{}{"name":"雄安","hobby":"副都"})
	db.Where("id",9).Find(&us)
	fmt.Println(us)
	type Company struct {

	}
	//子查询方式更新
	//更多查看  https://gorm.io/docs/update.html
	db.Model(&UserInfo{}).Update("company_name", db.Model(&Company{}).Select("name").Where("companies.id = users.company_id"))
	// UPDATE "users" SET "company_name" = (SELECT name FROM companies WHERE companies.id = users.company_id);


}
