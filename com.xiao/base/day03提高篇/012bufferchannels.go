package main

import "fmt"

/**
可以从已经关闭的缓冲通道中读取数据。
通道将返回已经写入通道的数据，
一旦读取所有数据，它将返回通道的零值
 */
func main() {
	ch := make(chan int, 5)
	ch <- 5
	ch <- 6
	close(ch)  //写入两个数据关闭通道
    //读取三个数据测试下
	//n, open := <-ch
	//fmt.Printf("Received: %d, open: %t\n", n, open)
	//n, open = <-ch
	//fmt.Printf("Received: %d, open: %t\n", n, open)
	//n, open = <-ch  //没有的数据返回0  并且通道状态为false
	//fmt.Printf("Received: %d, open: %t\n", n, open)
    /**
        使用循环测试
     */
	for n := range ch {
		fmt.Println("Received:", n)
	}
}
