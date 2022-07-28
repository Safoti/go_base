package main

import "fmt"

/**
 * @Author safoti
 * @Date Created in 2022/7/28
 * @Description  指针
 **/
type (
	Person struct {
		name string
	}
)

func main() {
	// type assertion check
	// ----
	var foo = getValue()

	intVal, ok := foo.(int)
	fmt.Printf("is 'int' type? ok: %v, intVal: %v\n", ok, intVal)

	floatVal, ok := foo.(float64)
	fmt.Printf("is 'float64' type? ok: %v, floatVal: %v\n", ok, floatVal)

	strVal, ok := foo.(string)
	fmt.Printf("is 'string' type? ok: %v, strVal: %v\n", ok, strVal)

	pPerson, ok := foo.(*Person)
	fmt.Printf("is '*Person' type? ok: %v, pPerson: %v\n", ok, pPerson)
}
func getValue() interface{} {
	ch := make(chan interface{}, 1)
	select {
	case ch <- 2010:
	case ch <- 9.15:
	case ch <- "Hello world!":
	case ch <- &Person{name: "Jane Doe"}:
	}

	return <-ch
}