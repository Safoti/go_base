package main

import "fmt"

/**
 * Created by safoti on 2022/7/11 9:28
**/
func main() {
  	// 1:go func() {
  	//	fmt.Println("hello")
	//}()

  	//2
	//go sayHello()

	//3 第三种调用
	sayHellos:=func(){
		fmt.Println("hello")
	}
	//三种使用gorouting 方式
	go sayHellos()
}

func  sayHello()  {
	fmt.Println("hello")
}
