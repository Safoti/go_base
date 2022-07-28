package main

import (
	"fmt"

	"math/rand"
	"time"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/28
 * @Description  Channels and Goroutines
 **/
var (
	s = rand.NewSource(time.Now().Unix())
	r = rand.New(s)
)
func main() {
  d:=make(chan string,5)
  producer(1,d)
  consumer(d)

}
func consumer(in chan string){
	count:=0
	for v := range in {
		count++
		fmt.Printf("Consumer got: %v\n", v)
	}
	if count ==0{
		fmt.Println("No data received")
		return
	}
	fmt.Printf("Processed %v items\n", count)
}

/**
  往管道里存放数据
 */
func producer(id int, out chan string) {
	n := r.Int() % cap(out)
	for i := 0; i < n; i++ {
		out <- fmt.Sprintf("Producer: %v, item: %v", id, i+1)
	}
	close(out)
}