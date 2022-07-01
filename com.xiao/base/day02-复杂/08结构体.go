package main

import "fmt"

/**
  结构体  就是java 中的对象，object
 */
func main() {
	//实例化并且打印
	address := Address{"Akshay", "PremNagar", "Dehradun", "Uttarakhand"}
	fmt.Println(address)


	//空对象
	var a Address
	fmt.Println(a)

	//根据属性进行特殊赋值操作
	ad:=Address{name: "lktbz",city: "Beijing"}
    fmt.Println(ad)
}

//type Address struct {
//	name string
//	street string
//	city string
//	state string
//
//}
//简化一下
type Address struct {
	name, street, city, state string
}