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
	fillCh(5, 5)
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

