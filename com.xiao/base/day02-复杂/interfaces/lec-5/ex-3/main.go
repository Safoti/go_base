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
	// type assertion pitfall
	// ----
	var foo interface{}
	foo = 1971.07

	var f float64
	f = foo.(float64)
	fmt.Printf("f's value: %v, type: %T\n", f, f)
	foo = 7
	fmt.Printf("foo's value: %v, type: %T\n", foo, foo)
	f = foo.(float64)
	fmt.Printf("f's value: %v, type: %T\n", f, f)
}
