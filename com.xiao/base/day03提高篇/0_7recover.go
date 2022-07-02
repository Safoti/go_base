package main

import (
	"fmt"
	"runtime/debug"
)
/**
   获取堆栈中的数据
 */
func recoverFullNamere() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
		debug.PrintStack() //打印堆栈
	}
}

func fullNamere(firstName *string, lastName *string) {
	defer recoverFullNamere()
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}
func main() {
	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	fullNamere(&firstName, nil)
	fmt.Println("returned normally from main")
}
