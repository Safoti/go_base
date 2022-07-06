package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
https://gorm.io/docs/has_many.html
 */

type Dog struct {
	ID   int
	Name string
	Toys []Toy `gorm:"polymorphic:Owner;"`
}

type Toy struct {
	ID        int
	Name      string
	OwnerID   int
	OwnerType string
}

func main() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/test_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Dog{})
	db.AutoMigrate(&Toy{})

	db.Create(&Dog{Name: "dog1", Toys: []Toy{{Name: "toy1"}, {Name: "toy2"}}})
	// INSERT INTO `dogs` (`name`) VALUES ("dog1")
	// INSERT INTO `toys` (`name`,`owner_id`,`owner_type`) VALUES ("toy1","1","dogs"), ("toy2","1","dogs")

}
