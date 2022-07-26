package main

import (
	"fmt"
	"reflect"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/26
 * @Description
 **/
func main() {
	var x interface{}
	x = 3.14
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	fmt.Printf("x: type = %v, value = %v\n", t, v)
	goo := x
	fmt.Printf("goo: type = %T, value = %v\n", goo, goo)

	x = &struct{ name string }{}
	t = reflect.TypeOf(x)
	v = reflect.ValueOf(x) // x.(<type>)
	fmt.Printf("x: type = %v, value = %v\n", t, v)
	hoo := x
	fmt.Printf("hoo: type = %T, value = %v\n", hoo, hoo)

}
