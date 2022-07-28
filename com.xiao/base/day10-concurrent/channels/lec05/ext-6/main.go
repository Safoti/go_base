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
	// selecting from multiple channels
	// ----
	var ch1, ch2 chan string
	select {
	case <-ch1:
		log.Println("Got data *FROM* ch1")
	case ch2 <- "hello":
		log.Println("Sent data *TO* ch2")
	default:
		log.Panic("No communication for ch1 or ch2")
	}
}
