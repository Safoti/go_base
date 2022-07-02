package main

import (
	"fmt"
	"time"
)

func hell()  {
	fmt.Println("Hello world goroutine")
}
func main() {
	go hell()
	time.Sleep(2*time.Second) //不然执行不了goroutine
	fmt.Println("main function")
}
