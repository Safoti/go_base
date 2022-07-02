package main

import (
	"fmt"
	"time"
)
/**
   通过sleep 函数更好的创建，主routine 会等待
 */
func  hellosleep(done chan bool)  {
	fmt.Println("Hello world goroutine")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")

	done <-true
}
func main() {
   done:=make(chan bool)  //创建channel
	fmt.Println("Main going to call hello go goroutine")
	go hellosleep(done)
   <-done  //获取数据
	fmt.Println("Main received data")

}
