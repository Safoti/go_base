package main

import (
	"fmt"

	"math/rand"
	"time"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/28
 * @Description  closing a channel
 **/
const (
	chCap = 10
)
var (
	s = rand.NewSource(time.Now().Unix())
	r = rand.New(s)
)

func main() {
	// passing channels to functions
	// ----
	d := make(chan int, chCap)
	if d == nil {

	}
	producer(d)
	consumer(d)
}
/**
   消费者拿到数据
 */
func consumer( nums chan int){
	for v := range nums {
		fmt.Printf("Consumer got: %v\n", v)
	}
}

// producer sends between 1 to cap(nums) random integers into the channel 'nums'
func producer(nums chan int) {
	n := r.Int()%cap(nums) + 1
	for i := 0; i < n; i++ {
		nums <- r.Int() % 200
	}
	close(nums)
}