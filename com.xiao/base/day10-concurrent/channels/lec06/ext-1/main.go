package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/28
 * @Description  Channel Selection
 **/

var (
	s = rand.NewSource(time.Now().Unix())
	r = rand.New(s)
)
func main() {
	done := make(chan bool)
	d := intGen(done)
	for i := 0; i < 10; i++ {
		fmt.Println(<-d)
	}
	done <- true // signal to go routine to exit gracefully

}
// intGen returns a channel on which random ints are produced
func intGen(done chan bool) (out chan int) {
	out = make(chan int)
	go func() {
		for {
			select {
			case <-done:  //关闭信号
				close(out)
				return
			case out <- r.Int() % 200:
			}
		}
	}()
	return out
}