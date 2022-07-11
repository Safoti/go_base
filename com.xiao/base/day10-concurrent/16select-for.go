package main

import (
	"fmt"
	"time"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/11
 * @Description  select 结合for 循环
 **/
func main() {
  done:=make(chan interface{})
  go func() {
  	time.Sleep(5*time.Second)
  	close(done)
  }()
  workCount:=0
  loop:
	  for  {
		  select {
		  case <-done:
			  break loop
		  default:

		  }
		  workCount++
		  time.Sleep(1*time.Second)
	  }
	fmt.Printf ("Achieved %v cycles of work before signalled to stop. \n", workCount)
}
