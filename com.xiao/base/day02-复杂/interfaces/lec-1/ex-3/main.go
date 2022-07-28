package main

import "fmt"

/**
 * @Author safoti
 * @Date Created in 2022/7/28
 * @Description
 **/

type(
	Reader interface {
		Read()
	}
	Writer interface {
		Write() string
	}
	Circle interface {
		Radius() float64
		Area() float64
	}

	Rect interface {
		Width() float64
		Length() float64
		Area() float64
	}
)
func main() {
	// more examples of interface declaration
	// ----
	var dataSource Reader
	var printer Writer

	fmt.Printf("dataSource's value: %v, type: %T\n", dataSource, dataSource)
	fmt.Printf("printer's value: %v, type: %T\n", printer, printer)

}
