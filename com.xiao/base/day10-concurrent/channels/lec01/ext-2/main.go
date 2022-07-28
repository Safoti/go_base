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

	//a:= <-ch //管道里没有数据  不能进行读取数据
    //fmt.Println(a)
	ch <- 4  //写入管道数据




}
