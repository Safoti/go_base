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
	for v := range ch { // fails on deadlock
		fmt.Println(v)
	}
}
func fillCh(c,l int){
	ch=make(chan int ,c)
	for i := 0; i <l;  i++ {
		ch<-i
	}
}

