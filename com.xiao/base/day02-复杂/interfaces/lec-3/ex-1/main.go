package main

import "fmt"

/**
 * @Author safoti
 * @Date Created in 2022/7/28
 * @Description   Method Sets
 **/

type (
	Person struct {
		fname string
		lname string
		age   uint8
	}
)
func main() {
	// method sets for T
	// ----
	mark := Person{"Mark", "Smith", 35}
	fmt.Printf("Mark's name: %v\n", mark.Name())
	fmt.Printf("Mark's age: %v\n", mark.Age())
}

func (p Person) Name() string {
	return fmt.Sprintf("%v, %v", p.lname, p.fname)
}
func (p Person) Age() uint8 {
	return p.age
}