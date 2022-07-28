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
 * @Description  closing a channel
 **/
var (
	s = rand.NewSource(time.Now().Unix())
	r = rand.New(s)
)
func main() {
	// sink - count numbers
	// ----
	done := make(chan bool)
	d := intGen(done)
	counter(d)
	time.Sleep(1 * time.Second)
	done <- true // signal to go routine to exit gracefully
}
// counter counts numbers receieved on a channel
func counter(in chan int) {
	go func() {
		log.Println("Counter - starting work...")
		start := time.Now()
		var count int
		for range in {
			count++
		}
		fmt.Printf("Counter processed %v items in %v\n", count, time.Since(start))
	}()
}
// intGen returns a channel on which random ints are produced
func intGen(done chan bool) (out chan int) {
	out = make(chan int)
	go func() {
		for {
			select {
			case <-done:
				close(out)
				return
			case out <- r.Int() % 200:
			}
		}
	}()
	return out
}