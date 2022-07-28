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
	// receiving data from a channel
	<-ch // blocked receiving from nil channel



}
