package main

import "fmt"

/**
 * @Author safoti
 * @Date Created in 2022/7/28
 * @Description  receiving from a closed channel
 **/
func main() {
	sc:=make(chan string,2)
	sc<-"hello"
	close(sc)
	s := <-sc
	fmt.Printf("sc 1st value: %v, len: %v, cap: %v\n", s, len(sc), cap(sc))
	s = <-sc //不存在的数据也可以进行接收？？？
	fmt.Printf("sc 2nd value: %v, len: %v, cap: %v\n", s, len(sc), cap(sc))
}

