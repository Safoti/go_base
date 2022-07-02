package main

import "fmt"

//关闭通道
func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i //写入
	}
	close(chnl)  //关闭
}
func main() {
	ch := make(chan int)
	go producer(ch)
	/**
	for {
		v, ok := <-ch //读取
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
	}
   **/
	for v := range ch {
		fmt.Println("Received ",v)
	}
}