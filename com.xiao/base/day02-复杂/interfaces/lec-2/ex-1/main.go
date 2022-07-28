package main

import "fmt"

/**
 * @Author safoti
 * @Date Created in 2022/7/28
 * @Description  Implementing an interface
 **/

type Currency float64
func main() {
	// ----
	var c Currency = 11.04
	fmt.Println(c.String()) // $11.04
}
//method
func (c Currency) String() string {
	return fmt.Sprintf("$%.2f", float64(c))
}