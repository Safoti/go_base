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
	x = "Hello, world - 1104"
	t0 := reflect.TypeOf(x)
	v0 := reflect.ValueOf(x)
	fmt.Printf("x: type = %v, value = %v\n", t0, v0)



	type ID string
	x = ID("0707")
	t1 := reflect.TypeOf(x)
	v1 := reflect.ValueOf(x)
	fmt.Printf("x: type = %v, value = %v\n", t1, v1)



	fmt.Printf("t0: type = %v, kind = %v\n", t0, t0.Kind())
	fmt.Printf("t1: type = %v, kind = %v\n", t1, t1.Kind())
}
