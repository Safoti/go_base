package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/28
 * @Description  receiving from a closed channel
 **/
const (
	chCap = 10
)

var (
	s = rand.NewSource(time.Now().Unix())
	r = rand.New(s)
)

func main() {
	d := generator()
	consumer(d)
}
func generator()(out chan int){
	fmt.Println("Generator of random ints")
	out =make(chan int, chCap)
	for i := 0; i <cap(out) ; i++ {
		out<-r.Int()%200
	}
	close(out)
	return
}
func consumer(nums chan int) {
	for v := range nums {
		fmt.Printf("Consumer got: %v\n", v)
	}
}


