package main

import (
	"fmt"
	"time"
)
/**
 select 默认走的路径
 */
func process_select(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}
func main() {
	ch:=	make(chan string)
	go process_select(ch)
	for  {
		time.Sleep(1000*time.Millisecond)
		select {
		   case v:=<-ch:
			   fmt.Println("received value: ", v)
			   return
		default: //防止select 阻塞  运行走默认的
			fmt.Println("no value received")
		}
	}
}
