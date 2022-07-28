package main

import "fmt"

/**
 * @Author safoti
 * @Date Created in 2022/7/28
 * @Description   声明channel
 **/
func main() {
	 var ch chan int  //int 类型的channel
	fmt.Printf("ch: %v, len: %v, cap: %v\n", ch, len(ch), cap(ch))
	// making a channel without capacity (unbuffer)
	// ----
	ch = make(chan int)
	//有地址了
	
	fmt.Printf("ch: %v, len: %v, cap: %v\n", ch, len(ch), cap(ch))


	// sending data to unbuffered channel (without a receiver)
	// ----
	ch <- 2 // fails while trying to send



}
