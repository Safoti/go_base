package main

import "fmt"
//指针在变量中的使用
//取消地址引用
func main() {
	var y = 458
	var p = &y
	fmt.Println("Value stored in y = ", y)
	fmt.Println("Address of y = ", &y)
	fmt.Println("Value stored in pointer variable p = ", p)

	fmt.Println("取消地址引用p的值之前为:",p)  //指向的是地址
	fmt.Println("取消地址引用p的值之后为:",*p) //直接获取值、

	*p=500
   fmt.Println("改变*p 的值为：",p)
   fmt.Println("改变*p 的值为：",*p)
}
