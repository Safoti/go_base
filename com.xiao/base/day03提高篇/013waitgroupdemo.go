package main

import (
	"fmt"
	"sync"
	"time"
)
/**
   等待组，等待全部执行完毕后
 */
func  process(i int ,wg *sync.WaitGroup)  {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()  //计数器减一
}
func main() {
	no := 3
	var wg sync.WaitGroup

	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i,&wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}
