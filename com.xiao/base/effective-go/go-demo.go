package main

import (
	"fmt"
	"time"
)

/**
 * Created by safoti on 2022/7/11 9:20
**/
func main() {

	var data int
	go func() {
		data++
	}()
	//
	time.Sleep(1*time.Second)
	if data==0{
		fmt.Println(data)
	}
}
