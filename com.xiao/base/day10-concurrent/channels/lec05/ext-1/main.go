package main

import (
	"fmt"
	"log"
	"time"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/28
 * @Description  Channel Selection
 **/


func main() {
	// inserting delays the wrong way
	// ----
	fmt.Printf("Message 1 at: %v\n", time.Now())
	sleep(1 * time.Second)
	fmt.Printf("Message 2 at: %v\n", time.Now())
}
func sleep(delay time.Duration) {
	end := time.Now().Add(delay)
	for time.Now().Before(end) {
		log.Println("Just waiting here, nothing to do...")
	}
}