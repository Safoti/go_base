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
func main() {
	// Currency implements interface main.Stringer
	var c Currency = 11.04
	fmt.Println(c.String()) // $11.0	4
}
func (c Currency) String() string {
	return fmt.Sprintf("$%.2f", float64(c))
}