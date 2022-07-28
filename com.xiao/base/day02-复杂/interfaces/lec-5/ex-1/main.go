package main

import "fmt"

/**
 * @Author safoti
 * @Date Created in 2022/7/28
 * @Description   Type Assertion
 **/

type (
	Person struct {
		name string
	}
)
func main() {
   var foo interface{}
   var f =1920.1
   foo=f
	fmt.Printf("foo's value: %v, type: %T\n", foo, foo)

	foo = &Person{name: "Jane Doe"}
	fmt.Printf("foo's value: %v, type: %T\n", foo, foo)
	var goo interface{}
	goo = foo
	fmt.Printf("goo's value: %v, type: %T\n", goo, goo)
	foo = f
}