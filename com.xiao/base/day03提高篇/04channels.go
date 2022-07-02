package main

import "fmt"

func  hello(done chan bool)  {
	fmt.Println("Hello world goroutine")
	done <-true
}
func main() {
   done:=make(chan bool)  //创建channel
     go hello(done)
   <-done  //获取数据
   fmt.Println("main function")
}
