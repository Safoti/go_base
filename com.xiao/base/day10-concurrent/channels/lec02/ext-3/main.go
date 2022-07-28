package main

import "fmt"

/**
 * @Author safoti
 * @Date Created in 2022/7/28
 * @Description  receiving from a closed channel
 **/
func main() {
	sc:=make(chan string,10)
	sc <- "hi"
	sc <- "howdy"
	close(sc)
	var ok bool
	s, ok := <-sc  //添加判断条件
	fmt.Printf("1st value from sc: %v, ok: %v\n", s, ok)
	s, ok = <-sc
	fmt.Printf("2nd value from sc: %v, ok: %v\n", s, ok)
	s, ok = <-sc
	fmt.Printf("3rd value from sc: %v, ok: %v\n", s, ok)
}

