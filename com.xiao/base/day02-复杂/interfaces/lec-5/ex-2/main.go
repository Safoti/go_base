package main

import "fmt"

/**
 * @Author safoti
 * @Date Created in 2022/7/28
 * @Description
 **/

type (
	Person struct {
		name string
	}
)

func main() {
	var foo interface{}
	foo = 1971.07
	var f float64
	f=foo.(float64)
	fmt.Printf("f's value: %v, type: %T\n", f, f)

}
