package main

import "fmt"

func main() {
  var s * int
  var a = 45
  s=&a
 fmt.Println(s)
 fmt.Println(&s)
 fmt.Println(*s)
 fmt.Println(&a)



	var x int = 5748


	var p *int


	p = &x
	fmt.Println("Value stored in x = ", x)
	fmt.Println("Address of x = ", &x)
	fmt.Println("Value stored in variable p = ", p) //地址    p接收的是x 的地址
	fmt.Println("Value stored in variable p = ", *p)//值
	fmt.Println("Value stored in variable p = ", &p)//p  所在的地址
}
