package main

import "fmt"

/**
 * @Author safoti
 * @Date Created in 2022/7/26
 * @Description   types --反射初步
 **/
func main() {
	  var x interface{}
	  x=3.14
	  fmt.Printf("x:type = %T,value = %v\n",x,x)

	  goo:=x
	  fmt.Printf("goo: type = %T, value = %v\n", goo, goo)
	  x=&struct {
		  name string
	  }{}
	  fmt.Printf("goo: type = %T, value = %v\n", goo, goo)
	  hoo := x
	  fmt.Printf("hoo: type = %T, value = %v\n", hoo, hoo)
}
