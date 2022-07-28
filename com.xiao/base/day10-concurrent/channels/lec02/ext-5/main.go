package main

import "fmt"

/**
 * @Author safoti
 * @Date Created in 2022/7/28
 * @Description  incorrect ways of iterating over a channel's values
 **/
var (
	ch chan int
)
func main() {
	fillCh(5, 5)
	for i := 0; i < cap(ch); i++ {
		fmt.Println(<-ch)
	}
}
func fillCh(c,l int){
	ch=make(chan int ,c)
	for i := 0; i <l;  i++ {
		ch<-i
	}
}

