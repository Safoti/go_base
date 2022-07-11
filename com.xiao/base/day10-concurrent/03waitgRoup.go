package main

import (
	"fmt"
	"sync"
	"time"
)

/**
 * Created by safoti on 2022/7/11 9:50
WaitGroup :使用方式

当你不关心并发操作的结果，或者你有其他方怯来收集它们的结果时，
WaitGrou p 是等待一组并发操作完成的好方峰。如果这两个条件都不满足，
我建议你使用c h an n el 和sel e ct 语句。

**/
func main() {
  var wg sync.WaitGroup
  wg.Add(1) //一个gorunting 开始
  go func() {
  		defer wg.Done()  //表示结束
	  fmt.Println ("1st goroutine sleeping .. . ")
  	  time.Sleep(1)

  }()
  wg.Add(1)
  go func() {
	  defer wg.Done()
	  fmt.Println ("2st goroutine sleeping .. . ")
	  time.Sleep(1)
  }()

  wg.Wait() //阻塞主routing 等待所有的gorouting 执行完毕在退出
  fmt.Println("全部执行完毕gogo")
}
