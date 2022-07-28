package main

import "fmt"

/**
 * @Author safoti
 * @Date Created in 2022/7/28
 * @Description  closing a channel
 **/
func main() {
	sc:=make(chan string,2)
	sc<-"hello"
	close(sc)
	fmt.Printf("sc: %v, len: %v, cap: %v\n", sc, len(sc), cap(sc))
	// sending to a closed channel
	sc <- "world" // failed sending on a closed channel
}

