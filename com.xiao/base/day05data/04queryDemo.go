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
	// SELECT * FROM users WHERE name IN ('jinzhu','jinzhu 2');
	db.Where("name IN ?",[]string{"北京","北京2"}).Find(&users)
	fmt.Println(users)
	//like
	// SELECT * FROM users WHERE name LIKE '%jin%';
	db.Where("name LIKE ?","%州%").Find(&users)
	fmt.Println(users)
	// AND
	// SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;
	db.Where("name =? AND hobby = ?","深圳","地产").Find(&users)
	fmt.Println(users)
	//db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
	//结构体查询
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 ORDER BY id LIMIT 1;
	db.Where(&UserInfo{Name: "北京"}).First(&users)
	//db.Where(&UserInfo{Name: "北京",Hobby: "帝都"}).First(&users)
	fmt.Println(users)
	//Map 方式
	//// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;
	db.Where(map[string]interface{}{"name":"上海","Hobby":"魔都"}).First(&users)
    fmt.Println(users)
	//切片查询
   //// SELECT * FROM users WHERE id IN (20, 21, 22);
	db.Where([]int64{1,2,3,4}).Find(&users)
	fmt.Println(users)
	//指定字段查询
	db.Where(&UserInfo{Name: "广州"}, "name", "hobby").Find(&users)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 0;
	fmt.Println(users)
	db.Where(&UserInfo{Name: "上海"}, "hobby").Find(&users)
	// SELECT * FROM users WHERE age = 0;
	fmt.Println(users)
     //内联条件

    //sql
	db.Find(&users,"name=?","北京")
	fmt.Println(users)
	//查询具体的字段
	db.Select("name", "hobby").Find(&users)
	fmt.Println(users)


	//排序
	db.Order("id desc,name ").Find(&users)
	fmt.Println(users)
	//以下不经过测试，只是学习拼接sql 的方式
	type resultss struct {
		Name  string
		Email string
	}

	//join
	db.Model(&UserInfo{}).Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&resultss{})

   //详情参考  https://gorm.io/docs/query.html

var orders UserInfo
var results UserInfo
	type User struct {

	}
	type  Pet struct {

	}
   //子查询
	db.Where("amount > (?)", db.Table("orders").Select("AVG(amount)")).Find(&orders)
	// SELECT * FROM "orders" WHERE amount > (SELECT AVG(amount) FROM "orders");


	subQuery := db.Select("AVG(age)").Where("name LIKE ?", "name%").Table("users")
	db.Select("AVG(age) as avgage").Group("name").Having("AVG(age) > (?)", subQuery).Find(&results)
	// SELECT AVG(age) as avgage FROM `users` GROUP BY `name` HAVING AVG(age) > (SELECT AVG(age) FROM `users` WHERE name LIKE "name%")
	subQuery1 := db.Model(&User{}).Select("name")
	subQuery2 := db.Model(&Pet{}).Select("name")
	db.Table("(?) as u, (?) as p", subQuery1, subQuery2).Find(&User{})
	// SELECT * FROM (SELECT `name` FROM `users`) as u, (SELECT `name` FROM `pets`) as p

   //遍历
	rows, err := db.Model(&User{}).Where("name = ?", "jinzhu").Rows()
	defer rows.Close()

	for rows.Next() {
		var user User
		// ScanRows is a method of `gorm.DB`, it can be used to scan a row into a struct
		db.ScanRows(rows, &user)

		// do something
	}

}


