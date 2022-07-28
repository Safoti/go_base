package main

import (
	"fmt"
	"sync"

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
	//等待组解决
	wgProducers sync.WaitGroup
	wgConsumers sync.WaitGroup
)

func main() {
	d := make(chan string)
	producer(1, d)
	producer(2, d)
	producer(3, d)
	go consumer(d)
	wgProducers.Wait()
	close(d)           // close after all producers
	wgConsumers.Wait() // wait for consumers
}
/**
   消费者拿到数据
 */
func consumer(in chan string) {
	wgConsumers.Add(1)
	count := 0
	for v := range in {
		count++
		fmt.Printf("Consumer got: %v\n", v)
	}

	if count == 0 {
		fmt.Println("No data received")
		return
	}

	fmt.Printf("Processed %v items\n", count)
	wgConsumers.Done()
}
func producer(id int, out chan string) {
	wgProducers.Add(1)
	// launch goroutine to produce data
	go func() {
		i := 1
		end := time.Now().Add(1000 * time.Millisecond)

		for time.Now().Before(end) {
			out <- fmt.Sprintf("Producer: %v, item: %v", id, i)
			i++
		}
		// -- can't call close(out) anymore
		wgProducers.Done()
	}()
}