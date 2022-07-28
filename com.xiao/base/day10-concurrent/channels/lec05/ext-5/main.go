package main

import (
	"log"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/28
 * @Description
 **/



func main() {
	// the 'select' statement
	// ----
	var ch1 chan int
	select {
	case <-ch1:
		log.Printf("Got data from ch1")
	default:
		log.Panic("No data from ch1")
	}
}
