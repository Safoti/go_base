package main

import "fmt"

/**
 * @Author safoti
 * @Date Created in 2022/7/11
 * @Description
 **/
func main() {
     c:=make(chan int,4)
     go func() {
     	defer close(c)
     	defer fmt.Println("pro done")
		 for i := 0; i <5 ; i++ {
			 fmt.Println("sending :",i)
			 c <- i
		 }
	 }()

	for  ib := range c {
		fmt.Println("rec:",ib)
	}


	}

