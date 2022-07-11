package main

import (
	"fmt"
	"sync"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/11
 * @Description
 **/
func main() {
  begin:=make(chan interface{})
  var wg sync.WaitGroup
	for i := 0; i <5 ; i++ {
	  wg.Add(1)
	  go func(i int) {
	  	defer wg.Done()
	  	<-begin
	  	fmt.Println(i," has begun \n")
	  }(i)
	}
	fmt.Println("非阻塞....")
	close(begin)
	wg.Wait()

	//创建带缓冲的channel
	//buch:=make(chan interface{},10) 10个容量的通道的channel
}
