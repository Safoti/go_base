package main

import "go_base_demo/com.xiao/base/day08reflects/types-3/ext-2/foo"

/**
 * @Author safoti
 * @Date Created in 2022/7/26
 * @Description 
 **/
type (
	ID     string
	Person struct {
		name string
	}
)

func main() {
	foo.Println(Person{name: "Jane Doe"})
	foo.Println(&Person{name: "John Smith"})
}
