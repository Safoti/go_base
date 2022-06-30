package main

import (
	"fmt"
	//"os"
	//"runtime"
)

func main() {
	//变量定义
	//var a_v int
	//var b_v bool
	//var str_v string
	//整合定义变量
	//var (
	//	a_v int
	//	b_v bool
	//	str_v string
	//	)


	//变量类型的推断
	//a_v := 15
	//b_v:= "string"
	//fmt.Println( a_v,b_v)

	var (
		a = 15
		b = false
		str = "Go says hello to the world!"
		numShips = 50
		city string
	)
  fmt.Println(a,b,str ,numShips,city)


	var n int =2
	fmt.Println(n)
	//
	//var goos string=runtime.GOOS
	//fmt.Printf("The operating system is: %s\n", goos)
	//path := os.Getenv("PATH")
	//fmt.Printf("Path is %s\n", path)

}
