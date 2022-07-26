package main

import "fmt"

/**
 * @Author safoti
 * @Date Created in 2022/7/26
 * @Description  types-1 为断言的学习    -- 类型测试
 **/
type (
	ID string
	Peson struct {
		name string
	}
)
func main() {
	Println(true)
	Println(2010)
	Println(9.15)
	Println(7 + 7i)
	Println("Hello world!")
	Println(ID("19520925"))
	Println([5]byte{})
	Println([]byte{})
	Println(map[string]int{})
	Println(Peson{name: "Jane Doe"})
	Println(&Peson{name: "Jane Doe"})
	Println(make(chan int))
}
// Println is my simple println function
func Println(x interface{}) {
	fmt.Printf("type is '%T', value: %v\n", x, x)
}
func Println2(x int) {
	fmt.Printf("type is '%T', value: %v\n", x, x)
}