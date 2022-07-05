package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}
func main() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/test_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// Migrate the schema
	db.AutoMigrate(&Product{})
	// Create
	db.Create(&Product{Code: "D42", Price: 100})
	// Read
	var product Product
	 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42
	fmt.Println(product)
}
