package main

import (
	"fmt"
	"sync"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/26
 * @Description   生产者模式演示   去掉sleep 协调整个服务
 **/


var wgs sync.WaitGroup



func main() {

	launchWorkers(2)
	fmt.Println("main 执行")
   wgs.Wait()
	fmt.Println("全部执行完毕")
}


func  launchWorkers(c int)  {
	for i := 0; i <5 ; i++ {
		wgs.Add(1)
		go func() {
			fmt.Printf(" %d : i am worker %v\n",c,i)
			wgs.Done()
		}()
	}
}