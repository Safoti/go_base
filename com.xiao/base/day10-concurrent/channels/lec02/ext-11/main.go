package main

import "fmt"

/**
 * @Author safoti
 * @Date Created in 2022/7/28
 * @Description  Working with closed Channels
 **/
var (
	ch chan int
)
func main() {
	// iterating over a channel using the range operator
	fillCh(5, 3)
	close(ch) // close channel to tell 'range' no more data is expected
	for v := range ch { // fails on deadlock
		fmt.Println(v)
	}

	fillCh(5, 1)
	//通过关闭通道来告知不需要更多的数据
	close(ch) // close channel to tell 'range' no more data is expected
	for v := range ch {
		fmt.Println(v)
	}
}
func fillCh(c,l int){
	ch=make(chan int ,c)
	for i := 0; i <l;  i++ {
		ch<-i
	}
}

