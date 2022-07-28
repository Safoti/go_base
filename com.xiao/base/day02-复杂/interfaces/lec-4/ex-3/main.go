package main
import "fmt"
/**
 * @Author safoti
 * @Date Created in 2022/7/28
 * @Description
 **/

type (
	ID     uint64
	SSN    string
	Person struct {
		name string
		age  uint8
		ssn  SSN
	}
	Empty interface{}
)
// Println is package's main custom print func
func Println(e Empty) {
	fmt.Printf("[main.Println] e's value: %v, type: %T\n", e, e)
}
func main() {
	// interface variable as function parameter
	// ----
	var e Empty
	Println(e)
	Println(11.04)
	Println(ID(19721104))
	Println(SSN("019-72-1104"))
	Println(&Person{name: "Jane Doe", age: 35, ssn: SSN("019-72-1104")})
}
