package main

import "fmt"

func  ptf(a *int)  {
	*a=748
}
func main() {
var x=100
fmt.Println("The value of x before function call is: %d\n", x)

var pa *int =&x
ptf(pa)
fmt.Println("The value of x after function call is: %d\n", x)


//测试传递地址
ptf(&x)
fmt.Println("The value of x after function call is: %d\n", x)
}
