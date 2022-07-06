package main

//多对一

//type Cat struct {
//	ID    int
//	Name  string
//	Toy   Toy `gorm:"polymorphic:Owner;"`
//}
//
//type Dog struct {
//	ID   int
//	Name string
//	Toy  Toy `gorm:"polymorphic:Owner;"`
//}
//type Toy struct {
//	ID        int
//	Name      string
//	OwnerID   int
//	OwnerType string
//}

func main() {
	//dsn := "root:1234@tcp(127.0.0.1:3306)/test_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	panic(err)
	//}
	//db.AutoMigrate(&Cat{})
	//db.AutoMigrate(&Dog{})
	//db.AutoMigrate(&Toy{})
	//
	//db.Create(&Dog{Name: "dog1", Toy: Toy{Name: "toy1"}})

}
