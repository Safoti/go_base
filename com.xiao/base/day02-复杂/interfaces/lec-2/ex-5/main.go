package main

import "fmt"

/**
 * @Author safoti
 * @Date Created in 2022/7/28
 * @Description  指针
 **/

type (
	Currency float64
	Stringer interface {
		String() string
	}
)
func (c *Currency) String() string {
	return fmt.Sprintf("$%.2f", float64(*c))
}
func main() {
	// implementing an interface with *T
	// ----
	var c = new(Currency)
	*c = 11.04
	fmt.Println(c.String())
	fmt.Println(c)
	var mainStringer Stringer
	mainStringer = c
	fmt.Printf("mainStringer's value: %v, type: %T\n", mainStringer, mainStringer)
	var fmtStringer fmt.Stringer
	fmtStringer = c
	fmt.Printf("fmtStringer's value: %v, type: %T\n", fmtStringer, fmtStringer)
}
