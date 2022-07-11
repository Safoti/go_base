package main

import (
	"fmt"
	"sync"
)

/**
 * Created by safoti on 2022/7/11 9:58
**/
func main() {
  hs:= func(wg *sync.WaitGroup,id int) {
     defer wg.Done()
     fmt.Printf("hs form %v\n",id)
  }
	const numbs=5
	var wg sync.WaitGroup
	wg.Add(numbs)
	for i := 0; i <numbs ; i++ {
		go hs(&wg,i)
	}
	wg.Wait()
}
