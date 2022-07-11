package main

import (
	"fmt"
	"sync"
)

/**
 * @Description //TODO
 * @Date Created in 2022/7/11 12:13
 * @Author safoti
 **/
func main() {
  var count int
  incre:= func() {count++}
  var once sync.Once
  var incres sync.WaitGroup
  incres.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer incres.Done()
			once.Do(incre)
		}()
	}
	incres.Wait()
	fmt.Println("count is :",count) //count is : 1  只调用一次


}
