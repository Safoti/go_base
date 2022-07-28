package main

import "fmt"

/**
 * @Author safoti
 * @Date Created in 2022/7/28
 * @Description
 **/
func main() {
	// declaring an interface
	// ----
	type Empty interface {
	}

	// interface variable
	// ----
	var e Empty
	fmt.Printf("e's value: %v, type: %T\n", e, e)
}
