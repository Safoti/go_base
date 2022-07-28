package main

import "fmt"

/**
 * @Author safoti
 * @Date Created in 2022/7/28
 * @Description
 **/
type (
	Currency float64
	Stringer interface {
		String() string
	}
)
func (c Currency) String() string {
	return fmt.Sprintf("$%.2f", float64(c))
}

func main() {
	// Currency implements interface main.Stringer
	// ----
	var c Currency = 11.04
	fmt.Println(c.String()) // explicit call to Currency's String() method

	// Currency ALSO implements interface fmt.Stringer
	// ----
	fmt.Println(c) // implicitly call to Currency's String(), call by fmt.PrintX()
}
