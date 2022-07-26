package main

import "go_base_demo/com.xiao/base/day08reflects/types-3/ext-1/foo"

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
	foo.Println(true)
	foo.Println(2010)
	foo.Println(9.15)
	foo.Println(7 + 7i)
	foo.Println("Hello world!")
	foo.Println(ID("19520925"))
	foo.Println([5]byte{})
	foo.Println([]byte{})
	foo.Println(map[string]int{})
	foo.Println(Person{name: "Jane Doe"})
	foo.Println(&Person{name: "Jane Doe"})
	foo.Println(make(chan int))
}
