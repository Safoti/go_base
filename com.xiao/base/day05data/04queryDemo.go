package main

import (
	"fmt"
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
	//u1 := UserInfo{1, "七米", "男", "篮球"}
	//u2 := UserInfo{2, "沙河娜扎", "女", "足球"}
    //添加数据
    //db.Create(&u1)
	//db.Create(&u2)
      //查询
    var us=new(UserInfo)
    //db.First(us)
	//fmt.Printf("%v \n",us)

   //条件查询
   db.Where("id=2",).Find(&us)
   //fmt.Println(us)
    //var uss []UserInfo
    //db.First(&uss)
    //fmt.Println(uss)
	result := map[string]interface{}{}
	// db.Table("user_infos").First(&result)
	//fmt.Println(result)
	//map 接收多个数据
	db.Table("user_infos").Take(&result)
	//fmt.Println(result)
	//不会进行匹配
	//first := db.First(&us, "七米")
	//first := db.First(&us, "2")
    //fmt.Println(first)
    //另一种条件语句
    //var use=UserInfo{ID: 2}
	//db.First(&use)
    //fmt.Println(use)


  var users []UserInfo
    //条件语句
    //in
    db.Where("name IN ?",[]string{"北京","北京2"}).Find(&users)
	fmt.Println(users)

	//like
	db.Where("name LIKE ?","%州%").Find(&users)
	fmt.Println(users)



}
