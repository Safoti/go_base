package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/28
 * @Description
 **/
var (
	s  = rand.NewSource(time.Now().Unix())
	r  = rand.New(s)
	wg sync.WaitGroup
)
func main() {
	// processor - filter out even numbers
	// ----
	done := make(chan bool)
	d := intGen(done)
	d = passThroughCounter(d)
	d = filter(d, func(x int) bool {
		if x%2 == 1 {
			return true
		}
		return false
	})
	counter(d)
	time.Sleep(1 * time.Second)
	done <- true // signal to go routine to exit gracefully
	wg.Wait()
}
// passThroughCounter counts numbers receieved on a channel, but pass them on
func passThroughCounter(in chan int) (out chan int) {
	wg.Add(1)
	out = make(chan int)
	go func() {
		defer wg.Done()
		log.Println("Counter (PT) - starting work...")
		start := time.Now()
		var count int
		for v := range in {
			count++
			out <- v
		}
		fmt.Printf("Counter (PT) processed %v items in %v\n", count, time.Since(start))
		close(out)
	}()
	return
}

// filter keeps numbers matching condition, no filtering if 'cmp' == nil
func filter(in chan int, cmp func(int) bool) (out chan int) {
	wg.Add(1)
	out = make(chan int)

	if cmp == nil {
		cmp = func(int) bool { return true }
	}

	go func() {
		defer wg.Done()
		for v := range in {
			if cmp(v) {
				out <- v
			}
		}
		close(out)
	}()
	return
}

// counter counts numbers receieved on a channel
func counter(in chan int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
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
	wg.Add(1)
	out = make(chan int)
	go func() {
		defer wg.Done()
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