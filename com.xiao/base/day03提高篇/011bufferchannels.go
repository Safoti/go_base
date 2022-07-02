package main

import "fmt"

/**
   两个缓冲区，往里面存放三个数据，读取两个会发生什么样的结果
 */
func main() {
	ch := make(chan string, 2)
	ch <- "17"
	ch <- "4AM"
	ch <- "SQ"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	//fatal error: all goroutines are asleep - deadlock!
}
