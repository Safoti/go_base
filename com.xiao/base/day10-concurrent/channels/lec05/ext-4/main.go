package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/28
 * @Description
 **/


func main() {
	// ----
	fmt.Printf("Message 1 at: %v\n", time.Now())
	<-time.After(1 * time.Second)
	fmt.Printf("Message 2 at: %v\n", time.Now())
}
