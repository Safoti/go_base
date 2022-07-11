package main

import (
	"fmt"
	"sync"
)

/**
 * Created by safoti on 2022/7/11 10:11
**/
func main() {
  var cout int
  var lock sync.Mutex
  fmt.Println("init count is number: ",cout)
  increment:= func() {
  	  lock.Lock()
	  defer lock.Unlock()
  	  cout++
  	  fmt.Println(" 自增：",cout)
  }
  decrement:= func() {
  	lock.Lock()
  	defer lock.Unlock()
	cout--
	fmt.Println("自减：",cout)
  }
  var ar sync.WaitGroup
	for i := 0; i <5 ; i++ {
		ar.Add(1)
		go func() {
			defer ar.Done()
			increment()
		}()
	}

	for i := 0; i <5 ; i++ {
		ar.Add(1)
		go func() {
			defer ar.Done()
			decrement()
		}()

	}
	ar.Wait()
	fmt.Println("end number of count: ",cout)
	fmt.Println("全部执行完毕")
}
