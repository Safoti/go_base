package main

import "fmt"

/**
 * @Author safoti
 * @Date Created in 2022/7/28
 * @Description   test reading from an empty channel without closing
 **/
func main() {
	sc:=make(chan string,4)
	sc <- "hi"
	s, ok := <-sc  //添加判断条件
	fmt.Printf("1st value from sc: %v, ok: %v\n", s, ok)
	s, ok = <-sc
	fmt.Printf("2nd value from sc: %v, ok: %v\n", s, ok)

}

