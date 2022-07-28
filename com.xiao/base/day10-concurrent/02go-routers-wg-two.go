package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/26
 * @Description   生产者模式演示   去掉sleep 协调整个服务
 **/

var (
	s  = rand.NewSource(time.Now().Unix())
	r  = rand.New(s)
	wg sync.WaitGroup
)
func main() {
	//Producer(1) //顺序执行
	//Producer(2)
	start := time.Now()
	wg.Add(2)//添加两个信号量
	go producer(1)
	go producer(2)
	// give groutines time to complete work
	wg.Wait()
	elapse := time.Since(start)
	fmt.Printf("Non-idea wait took %v\n", elapse)
	}




func producer(is int) {
	n := (r.Int() % 1000) + 1
	d := time.Duration(n) * time.Millisecond
	time.Sleep(d)
	fmt.Printf("Producer # %v ran for %v\n", is, d)
	wg.Done() //进行减一操作


}
