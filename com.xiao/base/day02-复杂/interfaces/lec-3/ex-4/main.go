package main

import "fmt"

/**
 * @Author safoti
 * @Date Created in 2022/7/28
 * @Description
 **/

type (
	Person struct {
		fname string
		lname string
		age   uint8
	}
)
func (p Person) Name() string {
	return fmt.Sprintf("%v, %v", p.lname, p.fname)
}
func (p Person) Age() uint8 {
	return p.age
}
func (p *Person) SetAge(a uint8) {
	if a <= 150 && a > p.age {
		fmt.Printf("Chaning age from %v to %v\n", p.age, a)
		p.age = a
	}
}
func main() {
	mark:=Person{"Mark","Smith",35}
	mark.SetAge(mark.Age()+1)

	family:=map[string]Person{
		"mom": {"Jane", "Smith", 42},
		"dad": {"Mark", "Smith", 35},
	}
	fmt.Println("Family members:")
	for r, p := range family {
		fmt.Printf("Role: %v, Name: %v, Age: %v\n", r, p.Name(), p.Age())
	}
	fmt.Println("Increment age of each family member")

	for _, p := range family {
		fmt.Println(&p)
		p.SetAge(p.Age() + 1)
	}
	for r, p := range family {
		fmt.Printf("Role: %v, Name: %v, Age: %v\n", r, p.Name(), p.Age())
	}
}
