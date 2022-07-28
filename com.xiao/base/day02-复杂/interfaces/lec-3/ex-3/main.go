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
	p := &Person{"Mark", "Smith", 35}
	fmt.Printf("Mark's name: %v\n", p.Name()) // -> *(p).Name()
	fmt.Printf("Mark's age: %v\n", p.Age())   // -> *(p).Age()
	p.SetAge(p.Age() + 1)
	fmt.Printf("Mark's age: %v\n", p.Age()) // -> *(p).Age()
}
