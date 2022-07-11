package main

import (
	"fmt"
	"time"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/11
 * @Description  描述
 **/
func main() {
   start:= time.Now()
	c:=make(chan interface{})
	go func() {
		fmt.Println("go 匿名函数执行前")
		time.Sleep(5*time.Second)
		fmt.Println("go 匿名函数执行后")
		close(c)
	}()
	fmt.Println("blocking on read....")
	select {
	case <-c:
		fmt.Printf("%v  后关闭",time.Since(start))

	}
}
