package main

import "fmt"

/**
 * @Author safoti
 * @Date Created in 2022/7/11
 * @Description   select 简单的使用
 **/
func main() {
 var c1,c2<-chan interface{}
 var c3 <-chan interface{}
	select {
	case <-c1:
		fmt.Println("执行c1")
	case <-c2:
		fmt.Println("执行c2")
	case <-c3:
		fmt.Println(c3)
 }
}
