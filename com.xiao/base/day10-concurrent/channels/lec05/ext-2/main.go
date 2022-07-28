package main

import (
	"fmt"
	"time"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/28
 * @Description  closing a channel
 **/

func main() {
	// inserting delays with time.Sleep()
	// ----
	fmt.Printf("Message 1 at: %v\n", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Printf("Message 2 at: %v\n", time.Now())
}
