package main

import (
	"fmt"
	"time"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/11
 * @Description   多个channel 向通道写数据
 **/
func main() {

	c1:=make(chan  interface{});close(c1)
	c2:=make(chan interface{});close(c2)
	var c1c,c2c int
	for i := 1000; i >0 ; i-- {
		select {
		case <-c1:
			c1c++
		case <-c2:
			c2c++
		case <-time.After(1*time.Second):
			fmt.Println("超时机制")
		default:
			fmt.Println("默认机制")
		}
	}
	fmt.Printf("c1:= %v c2:= %v",c1c,c2c)
}
